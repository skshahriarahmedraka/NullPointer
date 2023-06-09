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
	"go.mongodb.org/mongo-driver/mongo"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/x/mongo/driver/operation"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) SearchBlogListQueryWithMetadata(c *fiber.Ctx) error {


  
	if c.Params("search")!="" && c.Query("type") != "" && c.Query("start") != "" && c.Query("stop") != "" && c.Query("order") != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbBlogData := []model.BlogData{}
		opts := options.Find()
		order, _ := strconv.Atoi((c.Query("order")))
		var cursor *mongo.Cursor
		var err error
		var limit int
		var SearchQuery string =c.Params("search")
		var filter primitive.D
		//  type : time , views , unanswered  , vote
		switch c.Query("type") {
		case "time":

			opts.SetSort(bson.D{{Key: "WrittenTime", Value: order}})

			filter = bson.D{
				{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}},
				{Key:"Title",Value: bson.M{"$regex": SearchQuery, "$options": "i"}},
			}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoBlogCol.Find(ctx, filter, opts)
			logs.Error("search Error while finding Blog data  time", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching search blog data time",
				})
			}

		case "views":

			opts.SetSort(bson.D{{Key: "Views", Value: order}})

			// filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			filter = bson.D{
				{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}},
				{Key:"Title",Value: bson.M{"$regex": SearchQuery, "$options": "i"}},
			}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoBlogCol.Find(ctx, filter, opts)
			logs.Error("search Error while finding Blog data views", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching search Blog data views",
				})
			}

		case "uncommented":

			opts.SetSort(bson.D{{Key: "WrittenTime", Value: order}})

			// filter = bson.D{{Key: "QuesAnsAccepted", Value: bson.D{{Key: "$eq", Value: ""}}}}
			filter = bson.D{
				{Key: "Comments", Value: bson.D{{Key: "$not", Value: bson.D{{Key: "$size", Value: 0}}}}},
				{Key:"Title",Value: bson.M{"$regex": SearchQuery, "$options": "i"}},
			}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoBlogCol.Find(ctx, filter, opts)
			logs.Error("search Error while finding Blog data uncommented", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching search  blog data uncommented",
				})
			}
		case "votes":
			// not working
			limit, _ = strconv.Atoi(c.Query("stop"))
			matchStage := bson.D{{Key: "$match", Value: bson.D{}}} // Add any matching conditions if needed
			projectStage := bson.D{{Key: "$project", Value: bson.D{
				{Key: "_id", Value: 1},
				{Key: "Upvote", Value: 1},
				{Key: "Downvote", Value: 1},
				{Key: "Title", Value: 1},
				{Key: "MetaTitle", Value: 1},
				{Key: "Image", Value: 1},
				{Key: "Description", Value: 1},
				{Key: "WrittenBy", Value: 1},
				{Key: "WrittenTime", Value: 1},
				{Key: "EditedBy", Value: 1},
				{Key: "EditedTime", Value: 1},
				{Key: "Tags", Value: 1},
				{Key: "Comments", Value: 1},
				{Key: "Votes", Value: 1},
				{Key: "Bookmark", Value: 1},
				{Key: "Views", Value: 1},
				{Key: "ViewedBy", Value: 1},
				{Key: "Difference", Value: bson.D{{Key: "$subtract", Value: bson.A{"$Upvote", "$Downvote"}}}},
			}}}

			sortStage := bson.D{{Key: "$sort", Value: bson.D{
				{Key: "Difference", Value: order},
			}}}

			filterStage := bson.D{
				{Key: "$match", Value: bson.D{
					{Key: "_id", Value: bson.D{
						{Key: "$ne", Value: 0},
					}},
					{Key: "Title", Value: bson.M{
						"$regex": SearchQuery,
						"$options": "i",
					}},
				}},
			}
			// Aggregation pipeline
			pipeline := mongo.Pipeline{filterStage,matchStage, projectStage, sortStage}

			// Execute the aggregation query
			cursor, err = H.MongoBlogCol.Aggregate(context.Background(), pipeline)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching Blog  data votes",
				})
			}

			


		default:
			opts.SetSort(bson.D{{Key: "WrittenTime", Value: order}})

			filter = bson.D{
				{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}},
				{Key:"Title",Value: bson.M{"$regex": SearchQuery, "$options": "i"}},
			}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoBlogCol.Find(ctx, filter, opts)
			logs.Error("search Error while finding Blog data  time", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching default search blog data time",
				})
			}


		}

		defer cursor.Close(ctx)
		NumOfBlog := cursor.RemainingBatchLength()
		start, _ := strconv.Atoi(c.Query("start"))
        fmt.Println("🚀 ~ file: SearchBlogListQueryWithMetadata.go ~ line 198 ~ ifc.Params ~ start : ", start)
		for i := 0; i < start; i++ {
			if !cursor.Next(context.Background()) {
				fmt.Println("skipping", i)
			}
		}

		for i := start; i < limit; i++ {
			if !cursor.Next(context.Background()) {
				var BlogArrWithMetadataType BlogArrWithMetadataType

				BlogArrWithMetadataType.Metadata.Length = NumOfBlog
				BlogArrWithMetadataType.BlogList = dbBlogData
                fmt.Println("🚀 ~ file: SearchBlogListQueryWithMetadata.go ~ line 211 ~ if!cursor.Next ~ BlogArrWithMetadataType : ", BlogArrWithMetadataType)
				return c.Status(fiber.StatusOK).JSON(BlogArrWithMetadataType)
			}

			var blog model.BlogData
			if err := cursor.Decode(&blog); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while decoding blog",
				})
			}

			dbBlogData = append(dbBlogData, blog)
		}

		var BlogArrWithMetadataType BlogArrWithMetadataType

		BlogArrWithMetadataType.Metadata.Length = NumOfBlog
		BlogArrWithMetadataType.BlogList = dbBlogData
        fmt.Println("🚀 ~ file: SearchBlogListQueryWithMetadata.go ~ line 229 ~ ifc.Params ~ BlogArrWithMetadataType : ", BlogArrWithMetadataType)

		return c.Status(fiber.StatusOK).JSON(BlogArrWithMetadataType)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})

	}
}


type BlogArrWithMetadataType struct {
	BlogList []model.BlogData `json:"BlogList" bson:"BlogList"`
	Metadata MetadataType     `json:"Metadata" bson:"Metadata"`
}