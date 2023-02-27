package handler

import (
	"app/model"
	"net/http"

	"github.com/gin-gonic/gin"
)




func (H *DatabaseCollections)Blogs(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Add("Content-Type", "application/json")

	x:= model.ListOfBlog{
		{"Stack under attack: what we learned about handling DDoS attacks","https://heroicons.dev"},
		{"An unfiltered look back at April Fools’ 2022","https://heroicons.dev"},
		{"Solidity Tutorial - How to Build and Deploy an NFT Minting dApp with Solidity and React 🛠","https://heroicons.dev"},
		{"Understanding XA Transactions With Practical Examples in Go ","https://heroicons.dev"},
		{"Why you should use a Go backend in Flutter ","https://heroicons.dev"},
		{"Is anyone interested in contributing to a new OSS built with Go? Please join to develop a NO-CODE workflow engine! ","https://heroicons.dev"}	}
	c.JSON(http.StatusOK,x)
}