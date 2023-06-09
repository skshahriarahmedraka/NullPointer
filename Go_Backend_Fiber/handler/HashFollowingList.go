package handler

import (
	"app/logs"
	"app/model"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections)HashFollowingList(c *fiber.Ctx) error  {
	c.Accepts("application/json")
	fmt.Println("🚀 ~ file: HashFollowingList.go ~ line 18 ~ ifc.Params ~ c.Params(\"ID\") : ", c.Params("ID"))
	if c.Params("ID") != "" {
		ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()
		user := model.UserActivity{}
		id,_:=primitive.ObjectIDFromHex(c.Params("ID"))
		err := H.MongoUserActivityCol.FindOne(ctx, bson.D{ {Key: "_id", Value: id}}).Decode(&user)
		logs.Error("Error while finding user activity", err)
        fmt.Println("🚀 ~ file: HashFollowingList.go ~ line 25 ~ ifc.Params ~ user : ", user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while fetching user activity",
			})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	}else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

}
