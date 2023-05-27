package handler

import (
	"app/logs"
	"app/model"
	"context"
	"strconv"

	// "strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func (H *DatabaseCollections)BlogsList(c *fiber.Ctx) error  {
	if c.Query("type") != "" && c.Query("start") != "" && c.Query("stop") != "" && c.Query("order") != "" {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbBlogData := []model.BlogData{}
		opts := options.Find()
		order, _ := strconv.Atoi((c.Query("order")))
		var cursor *mongo.Cursor
		var err error
		var limit int


		var filter primitive.D
		//  type : time , views , unanswered  , vote
		switch c.Query("type") {
		case "time":

			opts.SetSort(bson.D{{Key: "QuesAskedTime", Value: order}})

			filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoQuestionCol.Find(ctx, filter, opts)
			logs.Error("Error while finding ques data", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching ques data",
				})
			}

		case "views":

			opts.SetSort(bson.D{{Key: "QuesViewed", Value: order}})

			filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoQuestionCol.Find(ctx, filter, opts)
			logs.Error("Error while finding ques data", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching ques data",
				})
			}

		case "unanswered":

			opts.SetSort(bson.D{{Key: "QuesAskedTime", Value: order}})

			filter = bson.D{{Key: "QuesAnsAccepted", Value: bson.D{{Key: "$eq", Value: 0}}}}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoQuestionCol.Find(ctx, filter, opts)
			logs.Error("Error while finding ques data", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching ques data",
				})
			}
		case "votes":

			matchStage := bson.D{{Key: "$match", Value: bson.D{}}} // Add any matching conditions if needed
			projectStage := bson.D{{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "QuesUpvote", Value: 1},
				{Key: "QuesDownvote", Value: 1},
				{Key: "Difference", Value: bson.D{{Key: "$subtract", Value: bson.A{"$QuesUpvote", "$QuesDownvote"}}}},
			}}}

			sortStage := bson.D{{Key: "$sort", Value: bson.D{
				{Key: "Difference", Value: order},
			}}}

			// Aggregation pipeline
			pipeline := mongo.Pipeline{matchStage, projectStage, sortStage}

			// Execute the aggregation query
			cursor, err = H.MongoQuestionCol.Aggregate(context.Background(), pipeline)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching ques data",
				})
			}
			
		default:
			opts.SetSort(bson.D{{Key: "QuesAskedTime", Value: order}})

			filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoQuestionCol.Find(ctx, filter, opts)
			logs.Error("Error while finding ques data", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching ques data",
				})
			}

		}

		defer cursor.Close(ctx)
		start, _ := strconv.Atoi(c.Query("start"))
		for i := 0; i < start; i++ {
			if !cursor.Next(context.Background()) {
				fmt.Println("skipping", i)
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

			dbQuesData = append(dbQuesData, ques)
		}

		
		fmt.Println("🚀 ~ file: quesListQuery.go ~ line 46 ~ forcursor.Next ~ dbQuesData : ", dbQuesData)

		return c.Status(fiber.StatusOK).JSON(dbQuesData)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})

	}
}
