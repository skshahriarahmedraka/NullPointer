package handler

import (
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (H *DatabaseCollections)RelatedQuestionListHandler(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Content-Type", "application/json")

	x:= model.RelatedQuestionList {							
			{"How to differentiate empty string and nothing in a map",342,true,"https://www.youtube.com/"},

			{"go - checking if key of one map present in another",2,false,"https://www.youtube.com/"},
			{"How to tell if the value of a map is undefined in Go?",54,true,"https://www.youtube.com/"},

			{"Checking if key exist in map which return interface type in go",777,true,"https://www.youtube.com/"},


			{"What is \"_\" (underscore comma) in a Go declaration?",7,false,"https://www.youtube.com/"},
			{"What is the Big O performance of maps in golang?",99,true,"https://www.youtube.com/"},
			{"How to check if a slice has a given index in Go?",999,false,"https://www.youtube.com/"},

			{"How can I create an array that contains unique strings?",87,true,"https://www.youtube.com/"},


			{"Check if key exists in multiple maps in one condition",70,true,"https://www.youtube.com/"},

			{"What is benefit of not raising error (like Python) when key is not in map?",9,true,"https://www.youtube.com/"}}
	c.JSON(http.StatusOK,x)

}