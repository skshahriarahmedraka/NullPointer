package handler

import (
	// "app/model"
	// "context"
	// "net/http"
	// "time"

	// // "encoding/json"
	// "fmt"

	"app/LogError"
	"app/model"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) Login(c *gin.Context) {
	c.Request.Header.Set("Access-Control-Allow-Origin", "*")
	c.Request.Header.Set("Content-Type", "application/json")
	c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

	var ReqData model.LoginData
	var dbUser model.UserData

	err := c.BindJSON(&ReqData)
	if err != nil {
		LogError.LogError("❌🔥 error in c.bindjson() ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}
	fmt.Println("🚀", ReqData)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// SEARCH EMAIL
	err = H.Mongo.Collection(os.Getenv("MONGO_USER_COL")).FindOne(ctx, bson.M{"Email": ReqData.Email}).Decode(&dbUser)
	if err != nil {
		LogError.LogError("❌🔥 error in mongodb connection  ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "No user Found"})
		return
	}

	
	userPass := []byte(ReqData.Password)
	dbPass := []byte(dbUser.Password)
	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		LogError.LogError("❌🔥 error in bcrypt error ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//CREATE COOKIE
	expirationTime := time.Now().Add(time.Hour * 1000)
	myClaim := &model.TokenClaims{
		Name:        dbUser.UserName,
		Email:       dbUser.Email,
		UserID:      dbUser.UserID,
		Accounttype: dbUser.Accounttype,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// }
	// LOGIN SUCCESSFUL

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		LogError.LogError("❌🔥 error in token.SignedString ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	////////////////////////////////
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth1",tokenString,60*60*24,"/","127.0.0.1",false , true)
	
	c.JSON(http.StatusOK,"successfully Register")


}
