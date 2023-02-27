package handler

import (
	// "encoding/json"

	"github.com/gofiber/fiber/v2"
)


func (H *DatabaseCollections)TEST(c *fiber.Ctx)  error  {
	data := struct {
		message string `json:"message"`
	} {
		message: "Test Hello World",
	}
	c.Status(fiber.StatusOK)
	return c.JSON(data)
	
	// return c.Next()
}