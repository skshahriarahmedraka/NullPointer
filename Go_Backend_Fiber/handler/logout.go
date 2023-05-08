package handler

import (

	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	// "github.com/golang-jwt/jwt/v4"

)

func (H *DatabaseCollections) Logout(c *fiber.Ctx) error {
	

	c.Cookie(&fiber.Cookie{
		Name:     "Auth1",
		Value:    "",
		Path:     "/",
		Domain:   os.Getenv("IP"),
		Expires:  time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
		HTTPOnly: true,

		SameSite: "lax",
	})

	// tokenString2 := utils.GenerateHttpCookie(dbUser)
	
	c.Cookie(&fiber.Cookie{
		Name:     "Info1",
		Value:    "",
		Path:     "/",
		Domain:   os.Getenv("IP"),
		Expires:  time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC),
		HTTPOnly: false,
		SameSite: "lax",
	})
	return c.Redirect("/login")
}
