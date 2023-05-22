package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)




type AnswerData struct {

	ID        primitive.ObjectID `json:"ID" bson:"_id"`
	QuesID   primitive.ObjectID `json:"QuesID" bson:"QuesID"`
	AnsweredBy  primitive.ObjectID `json:"AnsweredBy" bson:"AnsweredBy"`
	AnsweredTime time.Time `json:"AnsweredTime" bson:"AnsweredTime"`
	EditedBy string `json:"EditedBy" bson:"EditedBy"`
	EditedTime time.Time	`json:"EditedTime" bson:"EditedTime"`
	Upvote int `json:"Upvote" bson:"Upvote"`
	Downvote int `json:"Downvote" bson:"Downvote"`
	Bookmark int 	`json:"Bookmark" bson:"Bookmark"`
	Accepted bool `json:"Accepted" bson:"Accepted"`
	Description string `json:"Description" bson:"Description"`
	Comment []struct {
		ID primitive.ObjectID `json:"ID" bson:"_id"`
		Upvote int `json:"Upvote" bson:"Upvote"`
		Downvote int `json:"Downvote" bson:"Downvote"`
		UserID primitive.ObjectID  `json:"UserID" bson:"UserID"`
		UserName string `json:"UserName" bson:"UserName"`
		CommentTime time.Time `json:"CommentTime" bson:"CommentTime"`
		Comment string;
	}
};