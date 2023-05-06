package handler

import (
	"app/logs"
	"app/model"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/operation"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) QuesListQuery(c *fiber.Ctx) error {

	
	fmt.Println("🚀 ~ file: quesListQuery.go ~ line 25 ~ ifc.Query ~ c.Query(\"order\") : ", c.Query("order"))
	fmt.Println("🚀 ~ file: quesListQuery.go ~ line 25 ~ ifc.Query ~ c.Query(\"stop\") : ", c.Query("stop"))
	fmt.Println("🚀 ~ file: quesListQuery.go ~ line 25 ~ ifc.Query ~ c.Query(\"start\") : ", c.Query("start"))
	fmt.Println("🚀 ~ file: quesListQuery.go ~ line 25 ~ ifc.Query ~ c.Query(\"type\") : ", c.Query("type"))
	if c.Query("type") != "" && c.Query("start") != "" && c.Query("stop") != "" && c.Query("order") != "" {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbQuesData := []model.QuesData{}
		opts := options.Find()
		var filter primitive.D
		//  type : time , views , unanswered  , bounty
		switch c.Query("type") {
		case "time":
			if c.Query("order") == "1" {
				opts.SetSort(bson.D{{Key: "QuesAskedTime", Value: 1}})
			} else {

				opts.SetSort(bson.D{{Key: "QuesAskedTime", Value: -1}})
			}
			filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			
		case "views":
			if c.Query("order") == "1" {

				opts.SetSort(bson.D{{Key: "QuesViewed", Value: 1}})
			} else {
				
				opts.SetSort(bson.D{{Key: "QuesViewed", Value: -1}})
			}
			filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			
		case "unanswered":
			if c.Query("order") == "1" {

				opts.SetSort(bson.D{{Key: "QuesAskedTime", Value: 1}})
			} else {
				opts.SetSort(bson.D{{Key: "QuesAskedTime", Value: -1}})

			}
			filter = bson.D{{Key: "QuesAnsAccepted", Value: bson.D{{Key: "$eq", Value: 0}}}}
			
			// opts.Set(bson.D{{Key: "QuestionID", Value: bson.D{{Key: "$ne", Value: 0}}}})

			// case "bounty":
			// 	opts.SetSort(bson.D{{Key :"QuesAskedTime", Value:  1}})
		}
		// opts.SetSort(bson.D{{Key: "QuesAskedTime", Value: 1}})
		limit , _ :=strconv.Atoi(c.Query("stop"))
		opts.SetLimit(int64(limit))
		cursor, err := H.MongoQuestionCol.Find(ctx, filter, opts)
		logs.Error("Error while finding ques data", err)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while fetching ques data",
			})
		}
		defer cursor.Close(ctx)
		start ,_ := strconv.Atoi(c.Query("start"))
		for i := 0; i < start; i++ {
			if !cursor.Next(context.Background()) {
				fmt.Println("skipping",i)
			}
		}
	
		for i := start; i < limit; i++ {
			if !cursor.Next(context.Background()) {
				fmt.Println("🚀 ~ file: quesListQuery.go ~ line 100 ~ if!cursor.Next ~ dbQuesData : ", dbQuesData)
				return c.Status(fiber.StatusOK).JSON(dbQuesData)
			}
	
			var ques model.QuesData
			if err := cursor.Decode(&ques); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while decoding question",
				})
			}
	
			// questions = append(questions, q)
			dbQuesData = append(dbQuesData, ques)
		}
	

		// for cursor.Next(ctx) {
		// 	var ques model.QuesData
		// 	cursor.Decode(&ques)
		// }
		fmt.Println("🚀 ~ file: quesListQuery.go ~ line 46 ~ forcursor.Next ~ dbQuesData : ", dbQuesData)

		return c.Status(fiber.StatusOK).JSON(dbQuesData)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})

	}
}
