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

func (H *DatabaseCollections) BlogCreate(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var reqBlogData model.BlogData
	// var reqUserData model.UserData
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&reqBlogData)
    fmt.Println("🚀 ~ file: blogCreate.go ~ line 26 ~ func ~ reqBlogData : ", reqBlogData)
    logs.Error("🚀 ~ file: blogCreate.go ~ line 26 ~ func ~ err : ", err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	count, err := H.MongoBlogCol.CountDocuments(ctx, bson.M{ "$or": []bson.M{ {"Title": reqBlogData.Title},{"Description": reqBlogData.Description}}})
    fmt.Println("🚀 ~ file: blogCreate.go ~ line 37 ~ func ~ count : ", count)
    logs.Error("🚀 ~ file: blogCreate.go ~ line 37 ~ func ~ err : ", err)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while counting Blog title and description",
		})
	}
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "same question already exists",
		})
	}

	
	if count == 0 {
		reqBlogData.ID = primitive.NewObjectID()
		opts := options.InsertOne().SetBypassDocumentValidation(false)
		// filter := bson.D{{Key: "Email", Value: reqUserData.Email}}

		res, err := H.MongoBlogCol.InsertOne(ctx, reqBlogData, opts)
		if err != nil {
			logs.Error("mongodb InsertOne error", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "mongodb InsertOne  error",
			})

		} else {
			// res.Password = ""
			fmt.Println("🚀 ~ file: quesAsk.go ~ line 64 ~ func ~ reqQuesData : ", reqBlogData)
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
