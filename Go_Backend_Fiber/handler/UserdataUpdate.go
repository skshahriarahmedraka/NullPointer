package handler

import (
	"app/logs"
	"app/model"
	"context"
	"encoding/json"
	"fmt"

	// "encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (H *DatabaseCollections) UpdateUserData(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var reqUserData model.UserData
	err := json.Unmarshal(c.Body(), &reqUserData)
    logs.Error("🚀 ~ file: UserdataUpdate.go ~ line 22 ~ func ~ err : ", err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := H.MongoUserCol.CountDocuments(ctx, bson.M{"Email": reqUserData.Email})
	logs.Error("Error while counting user data", err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while counting user data",
		})
	}

	if count > 0 {
		opts := options.Update().SetUpsert(true)
		filter := bson.D{{Key: "Email", Value: reqUserData.Email}}
		
		update := bson.D{
			{Key: "$set", Value: bson.D{{Key: "UserID", Value: reqUserData.UserID}}},
			{Key: "$set", Value: bson.D{{Key: "UserName", Value: reqUserData.UserName}}},
			{Key: "$set", Value: bson.D{{Key: "Email", Value: reqUserData.Email}}},
			{Key: "$set", Value: bson.D{{Key: "Password", Value: reqUserData.Password}}},
			{Key: "$set", Value: bson.D{{Key: "UserTitle", Value: reqUserData.UserTitle}}},
			{Key: "$set", Value: bson.D{{Key: "UserImage", Value: reqUserData.UserImage}}},
			{Key: "$set", Value: bson.D{{Key: "Location", Value: reqUserData.Location}}},
			{Key: "$set", Value: bson.D{{Key: "Aboutme", Value: reqUserData.Aboutme}}},
			{Key: "$set", Value: bson.D{{Key: "Mysite", Value: reqUserData.Mysite}}},
			{Key: "$set", Value: bson.D{{Key: "Github", Value: reqUserData.Github}}},
			{Key: "$set", Value: bson.D{{Key: "Twitter", Value: reqUserData.Twitter}}},
			{Key: "$set", Value: bson.D{{Key: "Linkedin", Value: reqUserData.Linkedin}}},
			
		}
		res, err := H.MongoUserCol.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			logs.Error("mongodb update one error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "mongodb update one error",
			})

		} else {
			fmt.Println("response ", res)
			return c.Status(fiber.StatusOK).JSON(res)
			
		}

	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while updating user data",
		})
	}

///////////////
	// dbUserData := model.UserData{}
	// fmt.Println("🚀 ~ file: Userdata.go ~ line 67 ~ ifc.Params ~  : ", c.Params("ID"))
	// id, _ := primitive.ObjectIDFromHex(c.Params("ID"))
	// fmt.Println("🚀 ~ file: Userdata.go ~ line 68 ~ ifc.Params ~ id : ", id)
	// err = H.MongoUserCol.FindOne(ctx, bson.M{"_id": id}).Decode(&dbUserData)
	// logs.Error("Error while finding user data", err)
	// if err != nil {
		
	// }
	
}
