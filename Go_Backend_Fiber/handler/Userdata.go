package handler

import (
	"app/logs"
	"app/model"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (H *DatabaseCollections)UserData(c *fiber.Ctx) error  {
	c.Accepts("application/json")
	if c.Query("ID") != "" {
		ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbUserData := model.UserData{}
		err := H.MongoBlogCol.FindOne(ctx, bson.D{ {Key: "_id", Value: c.Query("ID")}}).Decode(&dbUserData)
		logs.Error("Error while finding blog", err)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while fetching blog",
			})
		}
		return c.Status(fiber.StatusOK).JSON(dbUserData)
	}else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

}
