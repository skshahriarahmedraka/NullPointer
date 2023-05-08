package model

import "go.mongodb.org/mongo-driver/bson/primitive"




type AnswerData struct {

	ID        primitive.ObjectID `json:"ID" bson:"_id"`
	AnsweredBy  string `json:"AnsweredBy" bson:"AnsweredBy"`
	AnsweredTime string
	EditedBy string
	EditedTime string;
	Upvote int
	Downvote int
	Bookmark int
	Accepted bool
	Description string
	Comment []struct {
		ID primitive.ObjectID `json:"ID" bson:"_id"`
		Upvote int
		Downvote int
		UserID string 
		UserName string;
		CommentTime string;
		Comment string;
	}
};