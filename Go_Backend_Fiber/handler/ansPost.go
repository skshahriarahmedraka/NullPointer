package handler

import (
	"app/logs"
	"app/model"
	"context"
	"encoding/json"

	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) AnsPost(c *fiber.Ctx) error {

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

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		count, err := H.MongoQuestionCol.CountDocuments(ctx, bson.M{"_id": c.Params("ID")})

		if err != nil {
			logs.Error("🚀 ~ file: ansPost.go ~ line 35 ~ ifc.Params ~ err : ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while counting user data",
			})
		}

		if count > 0 {

			count, err := H.MongoAnswerCol.CountDocuments(ctx, bson.M{"$or": []bson.M{
				{"$and": []bson.M{{"AnsweredBy": reqAnsData.AnsweredBy}, {"QuestionID": c.Params("ID")}}},
				{"Description": reqAnsData.Description}}})

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

	// if  c.Params("ID")!= "" {
	// 	ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	defer cancel()

	// 	dbQuesData := model.QuesData{}
	//     fmt.Println("🚀  c.Query(\"ID\") : ", c.Params("ID"))
	// 	id,_:=primitive.ObjectIDFromHex(c.Params("ID"))
	//     fmt.Println("🚀 ~ file: quesData.go ~ line 26 ~ ifc.Params ~ id : ", id)
	// 	err := H.MongoQuestionCol.FindOne(ctx, bson.M{"_id": id}).Decode(&dbQuesData)
	// 	logs.Error("Error while finding ques data", err)
	// 	if err != nil {
	// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 			"message": "Internal Server Error while fetching ques data",
	// 		})
	// 	}

	// 	fmt.Println("🚀 ~ file: quesData.go ~ line 28 ~ ifc.Params ~ dbQuesData : ", dbQuesData)
	// 	return c.Status(fiber.StatusOK).JSON(dbQuesData)
	// }else {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"message": "Bad Request",
	// 	})
	// }
}
