package database

import (
	"os"

	"app/handler"
	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseInit(db *mongo.Database) handler.DatabaseCollections {

	c := handler.DatabaseCollections{
		MongoDB:          db,
		MongoUserCol:     db.Collection(os.Getenv("MONGO_USER_COL")),
		MongoQuestionCol: db.Collection(os.Getenv("MONGO_QUESTION_COL")),
		MongoAnswerCol:   db.Collection(os.Getenv("MONGO_ANSWER_COL")),
		MongoBlogCol:     db.Collection(os.Getenv("MONGO_BLOG_COL")),
		MongoGroupCol:    db.Collection(os.Getenv("MONGO_GROUP_COL")),
		MongoChannelCol:  db.Collection(os.Getenv("MONGO_CHANNEL_COL")),
	}
	return c
}
