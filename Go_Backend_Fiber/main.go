package main

import (
	"app/config"
	"app/router"
	"fmt"
	"os"

	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
)

func init() {
	config.LoadEnvVars()

}



func main() {
	// Custom config
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Null Pointer Go Fiber API",
	})

	// ...
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	router.RouteWithoutAuth(app)

	fmt.Println("Server Started IN : ",os.Getenv("HOST"))
	app.Listen(os.Getenv("HOST"))
}
