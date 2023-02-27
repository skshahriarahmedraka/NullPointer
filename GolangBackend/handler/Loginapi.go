package handler

import (
	"app/model"
	// "context"
	"encoding/json"
	"net/http"
	// "time"

	// "encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
)





func (H *DatabaseCollections)LoginApi(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	// c.Writer.Header().Add("Content-Type", "application/json")

	var d1 model.LoginData
	err :=json.NewDecoder(c.Request.Body).Decode(&d1)
	if err !=nil {
		fmt.Println("🚀 ~ file: Login.go ~ line 40 ~ func ~ err : ", err)
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":"ok"})

	


}
