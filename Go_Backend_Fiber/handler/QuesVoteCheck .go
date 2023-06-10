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
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) QuesVoteCheck(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var reqData model.QuesVoteCheck
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&reqData)
	logs.Error("🚀 ~ file: QuesVoteCheck .go ~ line 25 ~ func ~ err : ", err)
	fmt.Println("🚀 ~ file: QuesVoteCheck .go ~ line 25 ~ func ~ reqData : ", reqData)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	// var HashFlair model.HashFlair
	// HashFlair.ID = reqData.ID
	// HashFlair.Name = reqData.Name
	// HashFlair.Image = reqData.Image

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"_id": reqData.ID,
		
	}

	// count, err := H.MongoQuestionCol.CountDocuments(ctx, filter)
	// fmt.Println("🚀 ~ file: QuesVoteCheck .go ~ line 51 ~ func ~ count : ", count)

	var result model.QuesData

	err = H.MongoQuestionCol.FindOne(ctx, filter).Decode(&result)
    fmt.Println("🚀 ~ file: QuesVoteCheck .go ~ line 54 ~ func ~ result : ", result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while decoding result",
		})
	}

	for _, v := range result.QuesVotes {
		if v.UserID == reqData.UserID {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"vote": v.Vote,
			})
		}
	}
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"message": "Internal Server Error while counting user vote",
	// 	})
	// }
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"vote": 0,
	})

}
