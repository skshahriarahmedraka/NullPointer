package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type Comment struct {
	ID primitive.ObjectID `json:"ID" bson:"_id"`
	Upvote int `json:"Upvote" bson:"Upvote"`
	Downvote int `json:"Downvote" bson:"Downvote"`
	UserID primitive.ObjectID  `json:"UserID" bson:"UserID"`
	UserName string `json:"UserName" bson:"UserName"`
	CommentTime time.Time `json:"CommentTime" bson:"CommentTime"`
	Comment string;
}