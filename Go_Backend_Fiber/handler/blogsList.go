package handler

import (
	"app/logs"
	"app/model"
	"context"
	"fmt"
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
	fmt.Println("🚀 ~ file: blogsList.go ~ line 22 ~ ifc.Query ~ c.Query(\"order\") : ", c.Query("order"))
    fmt.Println("🚀 ~ file: blogsList.go ~ line 22 ~ ifc.Query ~ c.Query(\"stop\") : ", c.Query("stop"))
    fmt.Println("🚀 ~ file: blogsList.go ~ line 22 ~ ifc.Query ~ c.Query(\"start\") : ", c.Query("start"))
    fmt.Println("🚀 ~ file: blogsList.go ~ line 22 ~ ifc.Query ~ c.Query(\"type\") : ", c.Query("type"))
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

			opts.SetSort(bson.D{{Key: "WrittenTime", Value: order}})

			filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoBlogCol.Find(ctx, filter, opts)
			logs.Error("Error while finding blog data using time", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching blog time data",
				})
			}

		case "views":

			opts.SetSort(bson.D{{Key: "Views", Value: order}})

			filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoBlogCol.Find(ctx, filter, opts)
			logs.Error("Error while finding blog data using view", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching blog view data",
				})
			}

		case "uncommented":

			opts.SetSort(bson.D{{Key: "WrittenTime", Value: order}})

			filter = bson.D{{Key: "Comments", Value: bson.D{{Key: "$not", Value: bson.D{{Key: "$size", Value: 0}}}}}}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoBlogCol.Find(ctx, filter, opts)
			logs.Error("Error while finding blog data using uncommnent", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching blog uncomment data",
				})
			}
		case "votes":

			matchStage := bson.D{{Key: "$match", Value: bson.D{}}} // Add any matching conditions if needed
			projectStage := bson.D{{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 0},
				{Key: "Upvote", Value: 1},
				{Key: "Downvote", Value: 1},
				{Key: "Difference", Value: bson.D{{Key: "$subtract", Value: bson.A{"$Upvote", "$Downvote"}}}},
			}}}

			sortStage := bson.D{{Key: "$sort", Value: bson.D{
				{Key: "Difference", Value: order},
			}}}

			// Aggregation pipeline
			pipeline := mongo.Pipeline{matchStage, projectStage, sortStage}

			// Execute the aggregation query
			cursor, err = H.MongoBlogCol.Aggregate(context.Background(), pipeline)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching blog votes data",
				})
			}
			
		default:
			opts.SetSort(bson.D{{Key: "WrittenTime", Value: order}})

			filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoBlogCol.Find(ctx, filter, opts)
			logs.Error("Error while finding blog default data", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching blog default data",
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
				fmt.Println("🚀 ~ file: blogsList.go ~ line 147 ~ if!cursor.Next ~ dbBlogData : ", dbBlogData)
				return c.Status(fiber.StatusOK).JSON(dbBlogData)
			}
			
			var blog model.BlogData
			if err := cursor.Decode(&blog); err != nil {
				fmt.Println("🚀 ~ file: blogsList.go ~ line 147 ~ if!cursor.Next ~ dbBlogData : ", dbBlogData)
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while decoding blog",
				})
			}

			dbBlogData = append(dbBlogData, blog)
		}

		

		return c.Status(fiber.StatusOK).JSON(dbBlogData)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})

	}
}
