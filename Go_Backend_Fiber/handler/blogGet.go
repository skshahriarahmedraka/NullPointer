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

func (H *DatabaseCollections)BlogGet(c *fiber.Ctx) error  {
	c.Accepts("application/json")
	fmt.Println("🚀 ~ file: blogGet.go ~ line 16 ~ ifc.Params ~ c.Params(\"ID\") : ", c.Params("ID"))
	if c.Params("ID") != "" {
		ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()
		blog := model.BlogData{}
		id,_:=primitive.ObjectIDFromHex(c.Params("ID"))
		err := H.MongoBlogCol.FindOne(ctx, bson.D{ {Key: "_id", Value: id}}).Decode(&blog)
        fmt.Println("🚀 ~ file: blogGet.go ~ line 23 ~ ifc.Params ~ err : ", err)
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
