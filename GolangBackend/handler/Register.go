package handler

import (
	// "app/database"
	"app/LogError"
	"app/model"
	"context"
	"log"
	"net/http"
	"os"

	// "strings"
	"time"

	// "encoding/json"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	// "go.mongodb.org/mongo-driver/bson"
	// "github.com/dgrijalva/jwt-go"
)

func (H *DatabaseCollections) Register(c *gin.Context) {


// c.Request.Header.Set("Access-Control-Allow-Origin", "*")
// c.Request.Header.Set("Content-Type", "application/json")
// c.Request.Header.Set("Access-Control-Allow-Credentials", "true")

	var ReqData model.UserData
	err := c.BindJSON(&ReqData)
	if err != nil {
		LogError.LogError("❌🔥 error in c.bindjson() ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("🚀", ReqData)
	if ReqData.Email == os.Getenv("SUPER_ADMIN_EMAIL") {

		ReqData.Accounttype = os.Getenv("SUPER_ADMIN_ACCOUNTTYPE")
	

	} else {
		ReqData.Accounttype = "normal"
		ReqData.UserID =uuid.New().String()

	}
	ReqData.ID = primitive.NewObjectID()
	ReqData.UserImage = ""
	ReqData.Aboutme = ""

    // ReqDataTitle string `json:"ReqDataTitle"`
    ReqData.UserImage = ""
    ReqData.Badges =   struct {
        Reputation int `json:Reputation bson:"Reputation"`
        Gold       int `json:"Gold" bson:"Gold"`
        Silver     int `json:"Silver" bson:"Silver"`
        Bronze     int `json:"Bronze" 		bson:"Bronze"`
    } {
		Reputation: 1,
		Gold:       0,
		Silver:     0,
		Bronze:     0,
	}

    ReqData.Follower=[]string{}
    // Badges map[string]int
    ReqData.Location=""
    ReqData.MembershipTime=primitive.Timestamp{T:uint32(time.Now().Unix())}
    ReqData.LastSeen= primitive.Timestamp{T:uint32(time.Now().Unix())}
    ReqData.Aboutme=""
    ReqData.Mysite =""
    ReqData.Github   =""
    ReqData.Twitter       =""
    ReqData.Linkedin    =""
    ReqData.TopTags       = []string{}
    ReqData.TopTagsPercent =[]int{}
    ReqData.SelectedPanel = ""  
	ReqData.Following = []string{}



	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	// SEARCH EMAIL
	count, err := H.Mongo.Collection(os.Getenv("MONGO_USER_COL")).CountDocuments(ctx, bson.M{"Email": ReqData.Email})
	if err != nil {
		LogError.LogError("❌🔥 error in mongodb connection  ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if count > 0 {
		LogError.LogError("❌🔥 error in User already registered ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already registered"})
		return
	}


	hash, err := bcrypt.GenerateFromPassword([]byte(ReqData.Password), bcrypt.DefaultCost)
	ReqData.Password = string(hash)
	if err != nil {
		log.Println(err)
	}
	// INSERT USER
	res, err := H.Mongo.Collection(os.Getenv("MONGO_USER_COL")).InsertOne(ctx, ReqData)
	LogError.LogError("🚀 ~ file: register.go ~ line 102 ~ func ~ err : ", err)
	if err == nil {

		fmt.Println("🚀 ~ file: register.go ~ line 116 ~ func ~ res : ", res)
	}
	// var uMoney model.UserMoney
	// uMoney.UserID = user.UserID
	// uMoney.Coin = 0.0
	// uMoney.ReqList = []model.UserMoneyReq{}
	// res, err = H.Mongo.Collection(os.Getenv("ADMIN_MONEY_MANAGE_COL")).InsertOne(ctx, uMoney)
	// fmt.Println("🚀 ~ file: SveltekitRegister.go ~ line 115 ~ returnfunc ~ res : ", res)
	// if err != nil {
	// 	LogError.LogError("❌🔥 error in InsertOne() ", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	//mongoRes, err := H.Mongo.Collection("userdb").InsertOne(ctx, user)
	//fmt.Println("🚀 ~ file: register.go ~ line 80 ~ func ~ mongoRes : ", mongoRes)

	//if err != nil {
	//	LogError.LogError("❌🔥 error in mongodb Connection error ", err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "mongodb Connection error"})
	//
	//	return
	//}

	expirationTime := time.Now().Add(time.Hour * 100)

	claims := &model.TokenClaims{
		Name:        ReqData.UserName,
		Email:       ReqData.Email,
		UserID:      ReqData.UserID,
		Accounttype: ReqData.Accounttype,
		// Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		LogError.LogError("❌🔥 error in StatusInternalServerError token generation  ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in StatusInternalServerError token generation"})

		return
	}
	////////////////////////////////
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth1",tokenString,60*60*24,"/","127.0.0.1",false , true)
		// c.SetCookie("Auth1Refresh",refreshToken,60*60*24,"/","localhost",false , true)
		c.JSON(http.StatusOK,"successfully Register")



	// var SendJwt struct {
	// 	JWT string
	// }
	// SendJwt.JWT = tokenString

	// c.JSON(http.StatusOK, SendJwt)

}

// var jwtKey = []byte("secret_key")
// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }
