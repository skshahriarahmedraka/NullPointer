package handler

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseCollections struct {
	MongoDB *mongo.Database
	MongoUserCol *mongo.Collection
	MongoQuestionCol *mongo.Collection
	MongoAnswerCol *mongo.Collection
	MongoBlogCol *mongo.Collection
	MongoGroupCol *mongo.Collection
	MongoChannelCol *mongo.Collection
}