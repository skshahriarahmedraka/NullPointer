package handler

import (
	"app/logs"
	"app/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) QuesAsk(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var reqQuesData model.QuesData
	// var reqUserData model.UserData
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&reqQuesData)
    fmt.Println("🚀 /api/q/askquestion : ", reqQuesData)
	logs.Error("🚀 ~ file: quesAsk.go ~ line 20 ~ func ~ err : ", err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := H.MongoQuestionCol.CountDocuments(ctx, bson.M{"_id": reqQuesData.ID})
	logs.Error("🚀 ~ file: quesAsk.go ~ line 34 ~ func ~ err : ", err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while counting ques id",
		})
	}
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "same question already exists",
		})
	}

	count, err = H.MongoQuestionCol.CountDocuments(ctx, bson.M{"QuesTitle": reqQuesData.QuesTitle})
	logs.Error("🚀 ~ file: quesAsk.go ~ line 48 ~ func ~ err : ", err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while counting ques title",
		})
	}
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "same question already exists",
		})
	}

	if count == 0 {
		reqQuesData.ID = primitive.NewObjectID()
		opts := options.InsertOne().SetBypassDocumentValidation(false)
		// filter := bson.D{{Key: "Email", Value: reqUserData.Email}}

		res, err := H.MongoQuestionCol.InsertOne(ctx, reqQuesData, opts)
		if err != nil {
			logs.Error("mongodb InsertOne error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "mongodb InsertOne  error",
			})

		} else {
			// res.Password = ""
			fmt.Println("🚀 ~ file: quesAsk.go ~ line 64 ~ func ~ reqQuesData : ", reqQuesData)
			fmt.Println("🚀 ~ file: quesAsk.go ~ line 70 ~ func ~ res : ", res)

			return c.Status(fiber.StatusOK).JSON(res)

		}

	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while updating user data",
		})
	}

	// doc := bson.M{
	// 	"timestamp": primitive.NewDateTimeFromTime(t),
	// }

	// return c.Next()
}
