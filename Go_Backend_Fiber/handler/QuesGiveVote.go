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

	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

func (H *DatabaseCollections) QuesGiveVote(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var reqData model.QuesVoteGive
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&reqData)
    fmt.Println("🚀 ~ file: QuesGiveVote.go ~ line 25 ~ func ~ reqData : ", reqData)
    logs.Error("🚀 ~ file: QuesGiveVote.go ~ line 25 ~ func ~ err : ", err)
   
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}
	// var HashFlair model.HashFlair 
	// HashFlair.ID = reqData.ID
	// HashFlair.Name = reqData.Name
	// HashFlair.Image = reqData.Image


	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	///////////////////
	filter := bson.M{
		"_id": reqData.ID,
		
	}

	// count, err := H.MongoQuestionCol.CountDocuments(ctx, filter)
	// fmt.Println("🚀 ~ file: QuesVoteCheck .go ~ line 51 ~ func ~ count : ", count)

	var result model.QuesData

	err = H.MongoQuestionCol.FindOne(ctx, filter).Decode(&result)
    fmt.Println("🚀 ~ file: QuesVoteCheck .go ~ line 54 ~ func ~ result : ", result)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal Server Error while decoding result",
		})
	}
	myVote :=0
	for _, v := range result.QuesVotes {
		if v.UserID == reqData.UserID {
			myVote= v.Vote
		}
	}
	fmt.Println("🚀 ~ file: QuesGiveVote.go ~ line 61 ~ func ~ myVote : ", myVote)
	//////////////////////
	var update primitive.M

	if myVote == 0 {
		opts := options.Update().SetUpsert(true)
		filter := bson.D{{Key: "_id", Value: reqData.ID}}

		var VoteData model.QuesVoteData

		VoteData.ID =reqData.ID
		VoteData.UserID = reqData.UserID
		VoteData.Vote= reqData.Vote
		VoteData.VoteTime = time.Now().UTC()
		if reqData.Vote == 1 {

			update = bson.M{
				"$inc": bson.M{"QuesUpvote": 1},
				"$push": bson.M{
					"QuesVotes": VoteData,
				},
			}
		}else if reqData.Vote == -1 {
			update = bson.M{
				"$inc": bson.M{"QuesDownvote": 1},
				"$push": bson.M{
					"QuesVotes": VoteData,
				},
			}

		}
	
		res , err := H.MongoQuestionCol.UpdateOne(ctx, filter, update,opts)
        fmt.Println("🚀 ~ file: QuesGiveVote.go ~ line 86 ~ func ~ res : ", res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while update",
			})
		}
		return c.Status(fiber.StatusOK).JSON(res)
	}else {
		opts := options.Update().SetUpsert(true)
		filter := bson.D{{Key: "_id", Value: reqData.ID}}
		var VoteData model.QuesVoteData

		VoteData.ID =reqData.ID
		VoteData.UserID = reqData.UserID
		VoteData.Vote= reqData.Vote
		VoteData.VoteTime = time.Now().UTC()
		
		if reqData.Vote == 1 {
			update = bson.M{
				"$inc": bson.M{"QuesUpvote": 1, "QuesDownvote": -1},
				"$set": bson.M{
					"QuesVotes.$.Vote": reqData.Vote,
				},
			}
		}else if reqData.Vote == -1 {
			update = bson.M{
				"$inc": bson.M{"QuesUpvote": -1, "QuesDownvote": 1},
				"$set": bson.M{
					"QuesVotes.$.Vote": reqData.Vote,
				},
			}
		}
		// update := bson.M{
		// 	"$set": bson.M{
		// 		"QuesVotes.O.Vote": bson.M{"_id": reqData.Vote},
		// 	},
		// }
		res , err := H.MongoQuestionCol.UpdateOne(ctx, filter, update,opts)
        fmt.Println("🚀 ~ file: QuesGiveVote.go ~ line 102 ~ func ~ res : ", res)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal Server Error while set update",
			})
		}
		return c.Status(fiber.StatusOK).JSON(res)
		

		// opts := options.Update().SetUpsert(true)
		// filter := bson.D{{Key: "_id", Value: reqData.UserID}}


		// update := bson.M{
		// 	"$push": bson.M{
		// 		"FollowingHash": HashFlair,
		// 	},
		// }
	
		// res , err := H.MongoUserActivityCol.UpdateOne(ctx, filter, update,opts)
		// fmt.Println("🚀 ~ file: HashFollow.go ~ line 52 ~ func ~ res : ", res)
		// if err != nil {
		// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		// 		"message": "Internal Server Error while pushing user activity",
		// 	})
		// }
		// return c.Status(fiber.StatusOK).JSON(fiber.Map{
		// 	"message": true,
		// })
		
		
	}

	
}
