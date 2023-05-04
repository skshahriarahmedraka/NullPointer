package router

import (
	"app/database"
	"github.com/gofiber/fiber/v2"
)

func RouteWithoutAuth(app *fiber.App) {

	db := database.MongoDBConnection()
	minioClient := database.MinioInit()
	H := database.DatabaseInit(db, minioClient)

	r := app.Group("/api")
	r.Get("/test", H.TEST)

	r.Post("/login",H.Login)
	r.Post("/register",H.Register)

	// r.Get("/blogs",H.BlogsList)
	// r.Get("/blog/:ID",H.Blog)

	
	r.Get("/user/:ID",H.UserData)
	r.Get("/user/flair/:ID",H.UserDataFlair)
	r.Post("/updateuser",H.UpdateUserData)
	
	
	r.Post("/uploadimage",H.UploadImage)
	r.Post("/q/askquestion",H.QuesAsk)

	r.Get("/q/:ID",H.QuesData)
	
	
	// /////////////////////////////////////////////////////
	r.Get("/q/answer/:ID",H.QuesAsk)
	r.Get("/api/search/locations",H.SearchLocation)
	// r.Get("/collection",H.Channel)
	// r.Get("/favcollection",H.FavChannelList)
	// r.Get("/channel/:ID",H.Channel)
	r.Get("/channel",H.Channel)
	r.Get("/favchannel",H.FavChannelList)
	r.Get("/channel/:ID",H.Channel)

	// r.Get("/q/:ID",H.QuesData)
	r.Get("/hotQues", H.HotQues)
	r.Post("/related",H.RelatedQues)

	r.Get("/hash",H.Hash)
	r.Get("/favhash",H.FavHash)

	



}
