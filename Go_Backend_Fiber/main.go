package main

import (
	"os"
// "github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"app/config"
	"app/router"
)

func init(){
	config.LoadEnvVars()

}

func main() {
	// Custom config
app := fiber.New(fiber.Config{
    Prefork:       true,
    CaseSensitive: true,
    StrictRouting: true,
    ServerHeader:  "Fiber",
    AppName: "Null Pointer Go Fiber API",
})

// ...
// app.Use(middleware.Logger())
	// app.Use(middleware.Recover())

	router.RouteWithoutAuth(app)
	

	app.Listen(os.Getenv("HOST"))
}