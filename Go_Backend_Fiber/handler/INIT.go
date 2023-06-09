package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/minio/minio-go/v7"
)

type DatabaseCollections struct {
	MongoDB *mongo.Database
	MongoUserCol *mongo.Collection
	MongoQuestionCol *mongo.Collection
	MongoAnswerCol *mongo.Collection
	MongoBlogCol *mongo.Collection
	MongoGroupCol *mongo.Collection
	MongoHashCol *mongo.Collection
	MongoUserActivityCol *mongo.Collection
	MinioClient *minio.Client
}