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


func (H *DatabaseCollections)SearchHashListQueryWithMetadata(c *fiber.Ctx) error  {

	fmt.Println("🚀 ~ file: SearchHashListQueryWithMetadata.go ~ line 24 ~ ifc.Params ~ c.Query(\"order\") : ", c.Query("order"))
    fmt.Println("🚀 ~ file: SearchHashListQueryWithMetadata.go ~ line 24 ~ ifc.Params ~ c.Query(\"stop\") : ", c.Query("stop"))
    fmt.Println("🚀 ~ file: SearchHashListQueryWithMetadata.go ~ line 24 ~ ifc.Params ~ c.Query(\"start\") : ", c.Query("start"))
    fmt.Println("🚀 ~ file: SearchHashListQueryWithMetadata.go ~ line 24 ~ ifc.Params ~ c.Query(\"type\") : ", c.Query("type"))
    fmt.Println("🚀 ~ file: SearchHashListQueryWithMetadata.go ~ line 24 ~ ifc.Params ~ c.Params(\"search\") : ", c.Params("search"))
	
	if c.Params("search")!="" && c.Query("type") != "" && c.Query("start") != "" && c.Query("stop") != "" && c.Query("order") != "" {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbHashData := []model.HashData{}
		opts := options.Find()
		order, _ := strconv.Atoi((c.Query("order")))
		var cursor *mongo.Cursor
		var err error
		var limit int
		var SearchQuery string =c.Params("search")
		var filter primitive.D

		// var filter primitive.D
		//  type : time , views , unanswered  , vote
		switch c.Query("type") {
		case "alpha":

			opts.SetSort(bson.D{{Key: "Name", Value: order}}) // 1 for ascending order, -1 for descending order

			// filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			filter = bson.D{
				{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}},
				{Key:"Name",Value: bson.M{"$regex": SearchQuery, "$options": "i"}},
			}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoHashCol.Find(ctx, filter, opts)
			logs.Error("Error while finding blog data using alphabetial search", err)


			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching Hash alphabetical data search",
				})
			}

		case "popular":

			opts.SetSort(bson.D{{Key: "NumOfQuestion", Value: order}})

			// filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}

			filter = bson.D{
				{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}},
				{Key:"Name",Value: bson.M{"$regex": SearchQuery, "$options": "i"}},
			}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoHashCol.Find(ctx, filter, opts)
			logs.Error("Error while finding blog data using view", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching Hash popular data search",
				})
			}

		case "follower":

			opts.SetSort(bson.D{{Key: "NumOfFollower", Value: order}})
			// filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}

			filter = bson.D{
				{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}},
				{Key:"Name",Value: bson.M{"$regex": SearchQuery, "$options": "i"}},
			}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoHashCol.Find(ctx, filter, opts)
            logs.Error("🚀 ~ file: HashList.go ~ line 78 ~ switchc.Query ~ err : ", err)

			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching hash follower data",
				})
			}
	
			
		default:
			opts.SetSort(bson.D{{Key: "Name", Value: order}}) // 1 for ascending order, -1 for descending order

			// filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}}}
			filter = bson.D{
				{Key: "_id", Value: bson.D{{Key: "$ne", Value: 0}}},
				{Key:"Name",Value: bson.M{"$regex": SearchQuery, "$options": "i"}},
			}
			limit, _ = strconv.Atoi(c.Query("stop"))
			opts.SetLimit(int64(limit))
			cursor, err = H.MongoHashCol.Find(ctx, filter, opts)
			logs.Error("Error while finding blog data using alphabetial search", err)


			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while fetching Hash alphabetical data search",
				})
			}

		}

		defer cursor.Close(ctx)
		NumOfHash := cursor.RemainingBatchLength()
		start, _ := strconv.Atoi(c.Query("start"))
        fmt.Println("🚀 ~ file: SearchHashListQueryWithMetadata.go ~ line 128 ~ ifc.Params ~ start : ", start)
		for i := 0; i < start; i++ {
			if !cursor.Next(context.Background()) {
				fmt.Println("skipping", i)
			}
		}

		for i := start; i < limit; i++ {
			if !cursor.Next(context.Background()) {
				var HashArrWithMetadata HashArrWithMetadataType

				HashArrWithMetadata.Metadata.Length = NumOfHash
				HashArrWithMetadata.HashList = dbHashData
                fmt.Println("🚀 ~ file: SearchHashListQueryWithMetadata.go ~ line 141 ~ if!cursor.Next ~ HashArrWithMetadata : ", HashArrWithMetadata)
				return c.Status(fiber.StatusOK).JSON(HashArrWithMetadata)
			}

			var hash model.HashData
			if err := cursor.Decode(&hash); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"message": "Internal Server Error while decoding hash data search",
				})
			}

			dbHashData = append(dbHashData, hash)
		}

		var HashArrWithMetadata HashArrWithMetadataType

		HashArrWithMetadata.Metadata.Length = NumOfHash
		HashArrWithMetadata.HashList = dbHashData
        fmt.Println("🚀 ~ file: SearchHashListQueryWithMetadata.go ~ line 159 ~ ifc.Params ~ HashArrWithMetadata : ", HashArrWithMetadata)

		return c.Status(fiber.StatusOK).JSON(HashArrWithMetadata)
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})

	}
}




// type MetadataType struct {
// 	Length int `json:"Length" bson:"Length"`
// }