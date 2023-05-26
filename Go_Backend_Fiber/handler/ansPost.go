package handler

import (
	"app/logs"
	"app/model"
	"context"
	"encoding/json"
	"fmt"

	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) AnsPost(c *fiber.Ctx) error {
fmt.Println("🚀🚀 ~ file: ansPost.go ~ line 20 ~ func  AnsPost : ")

	if c.Params("ID") != "" {

		c.Accepts("application/json")
		var reqAnsData model.AnswerData
		err := json.Unmarshal(c.Body(), &reqAnsData)
		logs.Error("🚀 ~ file: ansPost.go ~ line 25 ~ ifc.Params ~ err :  ", err)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad Request while unmarshalling data",
			})
		}
		fmt.Println("🚀 ~ file: ansPost.go ~ line 80 ~ ifc.Params ~ reqAnsData : ", reqAnsData)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		count, err := H.MongoQuestionCol.CountDocuments(ctx, bson.M{"_id": reqAnsData.QuesID})
        fmt.Println("🚀 ~ H.MongoQuestionCol.CountDocuments(ctx, bson.M{\"_id\": reqAnsData.QuesID}) count : ", count)

		if err != nil {
			logs.Error("🚀 ~ file: ansPost.go ~ line 35 ~ ifc.Params ~ err : ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while counting user data",
			})
		}

		if count > 0 {
			reqAnsData.ID = primitive.NewObjectID()
			filter:= bson.M{
				"$or": []bson.M{
					{
						"$and": []bson.M{
							{"AnsweredBy": reqAnsData.AnsweredBy},
							{"QuesID": reqAnsData.QuesID},
						},
					},
					{"Description": reqAnsData.Description},
				},
			}
			count, err := H.MongoAnswerCol.CountDocuments(ctx, filter)
            fmt.Println("🚀 ~ file: ansPost.go ~ line 52 ~ ifc.Params ~ count : ", count)

			logs.Error("🚀 ~ file: ansPost.go ~ line 47 ~ ifc.Params ~ err : ", err)
			if count > 0 {

				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Answer already exists",
				})
			} else {
				reqAnsData.ID = primitive.NewObjectID()
				reqAnsData.QuesID, _ = primitive.ObjectIDFromHex(c.Params("ID"))
				// reqAnsData.AnsweredBy, _ = reqAnsData.AnsweredBy


				_, err := H.MongoAnswerCol.InsertOne(ctx,reqAnsData)
                logs.Error("🚀 ~ file: ansPost.go ~ line 62 ~ ifc.Params ~ err : ", err)
				if err == nil {
					var quesAnsIndicator model.QuesAnsIndicatorData
					quesAnsIndicator.ID = reqAnsData.ID 
					quesAnsIndicator.UpVote = 0
					quesAnsIndicator.DownVote = 0
					quesAnsIndicator.AnsweredBy = reqAnsData.AnsweredBy
					quesAnsIndicator.Comment = []string{}
					res,err:=H.MongoQuestionCol.UpdateOne(ctx, bson.M{"_id": reqAnsData.QuesID}, bson.M{"$push": bson.M{"Answers": quesAnsIndicator}})
                    logs.Error("🚀 ~ file: ansPost.go ~ line 78 ~ ifc.Params ~ err : ", err)
                    fmt.Println("🚀 ~ Updated quesdata with answer id : ", res)
					
				}

				return c.Status(fiber.StatusOK).JSON(reqAnsData)
			}

		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Question does not exists",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	
}
