package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuesData struct {
	ID                   primitive.ObjectID      `json:"ID" bson:"_id" validate:"required"`
	QuesTitle       string    `json:"QuesTitle" bson:"QuesTitle" validate:"required,min=4,max=200"`
	
	QuesViewed      int      `json:"QuesViewed" bson:"QuesViewed" validate:"omitempty"`
	QuesUpvote      int      `json:"QuesUpvote" bson:"QuesUpvote" validate:"omitempty"`
	QuesDownvote    int      `json:"QuesDownvote" bson:"QuesDownvote" validate:"omitempty"`
	QuesVotes 	 []QuesVoteData `json:"QuesVotes" bson:"QuesVotes" validate:"omitempty"`
	QuesBookmark    []QuesBookmarkData      `json:"QuesBookmark" bson:"QuesBookmark" validate:"omitempty"`
	QuesTags        []string `json:"QuesTags" bson:"QuesTags" validate:"omitempty"`
	QuesGroup 	 []string `json:"QuesGroup" bson:"QuesGroup" validate:"omitempty"`
	QuesAnsAccepted string   `json:"QuesAnsAccepted" bson:"QuesAnsAccepted" validate:"required"` // accepted answer id
	
	
	QuesAskedBy     primitive.ObjectID   `json:"QuesAskedBy" bson:"QuesAskedBy" validate:"required"`
	QuesAskedTime  time.Time `json:"QuesAskedTime" bson:"QuesAskedTime" validate:"omitempty"`
	

	QuesEditedBy primitive.ObjectID `json:"QuesEditedBy" bson:"QuesEditedBy" validate:"omitempty"`
	QuesEditedTime time.Time                 `json:"QuesEditedTime" bson:"QuesEditedTime" validate:"omitempty"`

	QuesDescription     string                `json:"QuesDescription" bson:"QuesDescription" validate:"required,min=10,max=10000"`
	QuesComment         []Comment 		   `json:"QuesComment" bson:"QuesComment" validate:"omitempty"`
	Answers                 []QuesAnsIndicatorData `json:"Answers" bson:"Answers" validate:"omitempty"`
}
type QuesVoteData struct {
	ID          primitive.ObjectID   `json:"ID" bson:"_id" `
	UserID	  primitive.ObjectID `json:"UserID" bson:"UserID" validate:"required"`
	Vote int `json:"Vote" bson:"Vote" validate:"required"`
	VoteTime time.Time `json:"VoteTime" bson:"VoteTime" validate:"omitempty"`
	
}
type QuesBookmarkData struct {
	ID 		primitive.ObjectID   `json:"ID" bson:"_id" `
	UserID	  primitive.ObjectID `json:"UserID" bson:"UserID" validate:"required"`
	BookmarkTime time.Time `json:"BookmarkTime" bson:"BookmarkTime" validate:"omitempty"`
}


type QuesAnsIndicatorData struct{ 
	ID  primitive.ObjectID `json:"ID" validate:"required"`
	UpVote int `json:"UpVote" validate:"omitempty"`
	DownVote int `json:"DownVote" validate:"omitempty"`
	AnsweredBy primitive.ObjectID `json:"AnsweredBy" validate:"required"`
	Comment  []string `json:"Comment" validate:"omitempty"`
}



type QuesFlair struct {
	ID primitive.ObjectID `json:"ID" bson:"_id"`
	QuesTitle       string    `json:"QuesTitle" bson:"QuesTitle" validate:"required,min=4,max=200"`
	QuesAskedBy     primitive.ObjectID   `json:"QuesAskedBy" bson:"QuesAskedBy" validate:"required"`
	QuesUpvote      int      `json:"QuesUpvote" bson:"QuesUpvote" validate:"omitempty"`
	QuesDownvote    int      `json:"QuesDownvote" bson:"QuesDownvote" validate:"omitempty"`
	QuesAnsAccepted string   `json:"QuesAnsAccepted" bson:"QuesAnsAccepted" validate:"required"` // accepted answer id
	Answers                 []QuesAnsIndicatorData `json:"Answers" bson:"Answers" validate:"omitempty"`

}