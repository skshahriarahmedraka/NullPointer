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

func (H *DatabaseCollections) HashFollowCheck(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var reqData model.HashViewData
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&reqData)
    logs.Error("🚀 ~ file: HashFollow.go ~ line 26 ~ func ~ err : ", err)
    fmt.Println("🚀 ~ file: HashFollow.go ~ line 26 ~ func ~ reqData : ", reqData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	var HashFlair model.HashFlair 
	HashFlair.ID = reqData.ID
	HashFlair.Name = reqData.Name
	HashFlair.Image = reqData.Image


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

		// opts := options.CountOptions{}
		// filter := bson.D{{Key: "_id", Value: reqData.ID}}


		// update := bson.M{
		// 	"$push": bson.M{
		// 		"FollowingHash": HashFlair,
		// 	},
		// }
		filter := bson.M{
			"_id": reqData.UserID, 
			"FollowingHash": bson.M{
				"$elemMatch": HashFlair,
			},
		}
	
		count, err := H.MongoUserActivityCol.CountDocuments(ctx, filter)
        fmt.Println("🚀 ~ file: HashFollowCheck.go ~ line 58 ~ func ~ count : ", count)

		
	
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while pushing user activity",
			})
		}

		if count > 0 {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": true ,
			})
		
		}else {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": false ,
			})
		}


}
