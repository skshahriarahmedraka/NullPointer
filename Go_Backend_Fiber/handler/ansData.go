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
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections)AnsData(c *fiber.Ctx) error  {

	if  c.Params("ID")!= "" {
		ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbAnsData := model.AnswerData{}
        fmt.Println("🚀Ans  c.Query(\"ID\") : ", c.Params("ID"))
		id,_:=primitive.ObjectIDFromHex(c.Params("ID"))
        fmt.Println("🚀 ~ file: ansData.go ~ line 26 ~ ifc.Params ~ id : ", id)
		err := H.MongoAnswerCol.FindOne(ctx, bson.M{"_id": id}).Decode(&dbAnsData)
        logs.Error("🚀 ~ file: ansData.go ~ line 28 ~ ifc.Params ~ err : ", err)
		
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while fetching Ans data",
			})
		}
		
		
        
		fmt.Println("🚀 ~ file: ansData.go ~ line 28 ~ ifc.Params ~ dbAnsData : ", dbAnsData)
		return c.Status(fiber.StatusOK).JSON(dbAnsData)
	}else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
}
