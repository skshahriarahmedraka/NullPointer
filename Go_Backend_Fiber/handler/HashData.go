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
func (H *DatabaseCollections)HashData(c *fiber.Ctx) error  {
	c.Accepts("application/json")
    fmt.Println("🚀 ~ file: HashData.go ~ line 19 ~ func ~ c.Params(\"ID\") : ", c.Params("ID"))
	if c.Params("ID") != "" {
		ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()
		hash := model.HashData{}
		id,_:=primitive.ObjectIDFromHex(c.Params("ID"))
		err := H.MongoHashCol.FindOne(ctx, bson.D{ {Key: "_id", Value: id}}).Decode(&hash)
        fmt.Println("🚀 ~ file: HashData.go ~ line 25 ~ ifc.Params ~ hash : ", hash)
        logs.Error("🚀 ~ file: HashData.go ~ line 25 ~ ifc.Params ~ err : ", err)
       
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while fetching blog",
			})
		}
		return c.Status(fiber.StatusOK).JSON(hash)
	}else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	// return c.Next()
}
