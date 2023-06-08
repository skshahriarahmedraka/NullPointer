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

func (H *DatabaseCollections) HashCreate(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var reqHashData model.HashData
	// var reqUserData model.UserData
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&reqHashData)
    logs.Error("🚀 ~ file: hashCreate.go ~ line 26 ~ func ~ err : ", err)
    fmt.Println("🚀 ~ file: hashCreate.go ~ line 26 ~ func ~ reqHashData : ", reqHashData)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := H.MongoHashCol.CountDocuments(ctx, bson.M{ "$or": []bson.M{ {"Name": reqHashData.Name},{"MetaTitle": reqHashData.MetaTitle}}})
    logs.Error("🚀 ~ file: hashCreate.go ~ line 37 ~ func ~ err : ", err)
	fmt.Println("🚀 ~ file: hashCreate.go ~ line 37 ~ func ~ count : ", count)
   
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while counting Hash  name and Metatitle",
		})
	}
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "same Name already exists",
		})
	}

	
	if count == 0 {
		reqHashData.ID = primitive.NewObjectID()
		opts := options.InsertOne().SetBypassDocumentValidation(false)
		// filter := bson.D{{Key: "Email", Value: reqUserData.Email}}

		res, err := H.MongoHashCol.InsertOne(ctx, reqHashData, opts)
		if err != nil {
			logs.Error("mongodb Hash InsertOne error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "mongodb Hash InsertOne  error",
			})
			
			} else {
				// res.Password = ""
				
				
				fmt.Println("🚀 ~ file: hashCreate.go ~ line 58 ~ func ~ reqHashData : ", reqHashData)
				fmt.Println("🚀 ~ file: hashCreate.go ~ line 58 ~ func ~ res : ", res)
			return c.Status(fiber.StatusOK).JSON(res)

		}

	} else {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while updating Hash Data",
		})
	}

	// doc := bson.M{
	// 	"timestamp": primitive.NewDateTimeFromTime(t),
	// }

	// return c.Next()
}
