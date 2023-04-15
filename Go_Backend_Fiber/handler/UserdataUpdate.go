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
	count, err := H.MongoUserCol.CountDocuments(ctx, bson.M{"UserID": reqUserData.UserID})
	logs.Error("Error while counting user data", err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while counting user data",
		})
	}

	if count > 0 {
		opts := options.Update().SetUpsert(true)
		filter := bson.D{{"UserID", reqUserData.UserID}}
		
		update := bson.D{
			{"$set", bson.D{{"UserID", reqUserData.UserID}}},
			{"$set", bson.D{{"UserName", reqUserData.UserName}}},
			{"$set", bson.D{{"Email", reqUserData.Email}}},
			{"$set", bson.D{{"Password", reqUserData.Password}}},
			{"$set", bson.D{{"UserTitle", reqUserData.UserTitle}}},
			{"$set", bson.D{{"UserImage", reqUserData.UserImage}}},
			{"$set", bson.D{{"Location", reqUserData.Location}}},
			{"$set", bson.D{{"Aboutme", reqUserData.Aboutme}}},
			{"$set", bson.D{{"Mysite", reqUserData.Mysite}}},
			{"$set", bson.D{{"Github", reqUserData.Github}}},
			{"$set", bson.D{{"Twitter", reqUserData.Twitter}}},
			{"$set", bson.D{{"Linkedin", reqUserData.Linkedin}}},
			
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
