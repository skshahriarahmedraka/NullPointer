package database

import (
	"os"

	"app/handler"

	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

func DatabaseInit(db *mongo.Database , minioClient *minio.Client) handler.DatabaseCollections {

	c := handler.DatabaseCollections{
		MongoDB:          db,
		MongoUserCol:     db.Collection(os.Getenv("MONGO_USER_COL")),
		MongoQuestionCol: db.Collection(os.Getenv("MONGO_QUESTION_COL")),
		MongoAnswerCol:   db.Collection(os.Getenv("MONGO_ANSWER_COL")),
		MongoBlogCol:     db.Collection(os.Getenv("MONGO_BLOG_COL")),
		MongoGroupCol:    db.Collection(os.Getenv("MONGO_GROUP_COL")),
		MongoHashCol:  db.Collection(os.Getenv("MONGO_HASH_COL")),
		MinioClient:      minioClient,
	}
	return c
}
