package handler

import (
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)



func (H *DatabaseCollections)Home(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Content-Type", "application/json")

	x:= model.HotQuestions{

	}
		c.JSON(http.StatusOK,x)
}