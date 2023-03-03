package handler

import (
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (H *DatabaseCollections)UserSpaces(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Content-Type", "application/json")

	x:= model.SpaceList{
		{"Go programming","https://stackoverflow.com/"},
		{"GTK","https://stackoverflow.com/"},
		{"QT framework","https://stackoverflow.com/"},
		{"Gui in Rust","https://stackoverflow.com/"},
		{"Gopher","https://stackoverflow.com/"},
		{"Rusty Boys","https://stackoverflow.com/"},
		{"Google Cloud","https://stackoverflow.com/"},
		{"Azure services","https://stackoverflow.com/"},
		{"Backend Engineering","https://stackoverflow.com/"},
		{"Web DEvelopment","https://stackoverflow.com/"}}
	c.JSON(http.StatusOK,x)
}