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

}
