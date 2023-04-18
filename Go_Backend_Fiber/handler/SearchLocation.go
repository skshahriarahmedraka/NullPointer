package handler

import (
	// "app/logs"
	// "app/model"
	// "context"
	// "time"

	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson"
)

func (H *DatabaseCollections)SearchLocation(c *fiber.Ctx) error  {
	c.Accepts("application/json")
	if c.Query("q") != "" {
		

		
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "ok",
		})
	}else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

}
