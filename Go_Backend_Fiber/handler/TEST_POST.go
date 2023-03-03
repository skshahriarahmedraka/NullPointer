package handler

import "github.com/gofiber/fiber/v2"




func (H *DatabaseCollections)TEST_POST(c *fiber.Ctx)  error  {
	data := struct {
		message string `json:"message"`
	} {
		message: "Test Hello World",
	}
	c.Status(fiber.StatusOK)
	return c.JSON(data)
	
	// return c.Next()
}