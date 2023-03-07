package handler

import (
	"app/logs"
	"app/model"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (H *DatabaseCollections)Channel(c *fiber.Ctx) error  {
	c.Accepts("application/json")
	if c.Query("ID") != "" {
		ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()
		blog := model.BlogData{}
		err := H.MongoChannelCol.FindOne(ctx, bson.D{ {"ID", c.Query("ID")}}).Decode(&blog)
		logs.Error("Error while finding blog", err)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while fetching blog",
			})
		}
		return c.Status(fiber.StatusOK).JSON(blog)
	}else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

}
