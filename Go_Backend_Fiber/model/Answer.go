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
	Votes []AnsVoteData `json:"Votes" bson:"Votes"`
	Bookmark []AnsBookmarkData `json:"Bookmark" bson:"Bookmark"`
	Accepted bool `json:"Accepted" bson:"Accepted"`
	Description string `json:"Description" bson:"Description"`
	Comments []Comment `json:"Comments" bson:"Comments"`
};

type AnsVoteData struct {
	ID          primitive.ObjectID   `json:"ID" bson:"_id" `
	UserID	  primitive.ObjectID `json:"UserID" bson:"UserID" validate:"required"`
	Vote int `json:"Vote" bson:"Vote" validate:"required"`
	VoteTime time.Time `json:"VoteTime" bson:"VoteTime" validate:"omitempty"`
	
}

type AnsBookmarkData struct {
	ID 		primitive.ObjectID   `json:"ID" bson:"_id" `
	UserID	  primitive.ObjectID `json:"UserID" bson:"UserID" validate:"required"`
	BookmarkTime time.Time `json:"BookmarkTime" bson:"BookmarkTime" validate:"omitempty"`
}
