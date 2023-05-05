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

		dbQuesData := model.QuesData{}
        fmt.Println("🚀  c.Query(\"ID\") : ", c.Params("ID"))
		id,_:=primitive.ObjectIDFromHex(c.Params("ID"))
        fmt.Println("🚀 ~ file: quesData.go ~ line 26 ~ ifc.Params ~ id : ", id)
		err := H.MongoQuestionCol.FindOne(ctx, bson.M{"_id": id}).Decode(&dbQuesData)
		logs.Error("Error while finding ques data", err)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while fetching ques data",
			})
		}
		
		
        
		fmt.Println("🚀 ~ file: quesData.go ~ line 28 ~ ifc.Params ~ dbQuesData : ", dbQuesData)
		return c.Status(fiber.StatusOK).JSON(dbQuesData)
	}else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
}
