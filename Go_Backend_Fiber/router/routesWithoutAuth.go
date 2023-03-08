package router

import (
	"app/database"
	"github.com/gofiber/fiber/v2"
)

func RouteWithoutAuth(app *fiber.App) {

	db := database.MongoDBConnection()
	H := database.DatabaseInit(db)

	r := app.Group("/api")
	r.Get("/test", H.TEST)

	r.Post("/login",H.Login)
	r.Post("/register",H.Register)

	r.Get("/blogs",H.BlogsList)
	r.Get("/blog/:ID",H.Blog)

	r.Get("/channel",H.Channel)
	r.Get("/favchannel",H.FavChannelList)
	r.Get("/channel/:ID",H.Channel)


	// r.Get("/collection",H.Channel)
	// r.Get("/favcollection",H.FavChannelList)
	// r.Get("/channel/:ID",H.Channel)

	r.Get("/q/:ID",H.Ques)
	r.Get("/hotQues", H.HotQues)
	r.Post("/related",H.RelatedQues)

	r.Get("/hash",H.Hash)
	r.Get("/favhash",H.FavHash)

	



}
