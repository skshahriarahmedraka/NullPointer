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
	QuesBookmark    int      `json:"QuesBookmark" bson:"QuesBookmark" validate:"omitempty"`
	QuesTags        []string `json:"QuesTags" bson:"QuesTags" validate:"omitempty"`
	QuesGroup 	 []string `json:"QuesGroup" bson:"QuesGroup" validate:"omitempty"`
	QuesAnsAccepted string   `json:"QuesAnsAccepted" bson:"QuesAnsAccepted" validate:"required"` // accepted answer id
	
	
	QuesAskedBy     primitive.ObjectID   `json:"QuesAskedBy" bson:"QuesAskedBy" validate:"required"`
	QuesAskedTime  time.Time `json:"QuesAskedTime" bson:"QuesAskedTime" validate:"omitempty"`
	

	QuesEditedBy primitive.ObjectID `json:"QuesEditedBy" bson:"QuesEditedBy" validate:"omitempty"`
	QuesEditedTime time.Time                 `json:"QuesEditedTime" bson:"QuesEditedTime" validate:"omitempty"`

	QuesDescription     string                `json:"QuesDescription" bson:"QuesDescription" validate:"required,min=10,max=10000"`
	QuesComment         []string  `json:"QuesComment" bson:"QuesComment" validate:"omitempty"`
	Answers                 []struct{ 
        ID  string `json:"ID" validate:"required"`
		UpVote int `json:"UpVote" validate:"omitempty"`
		DownVote int `json:"DownVote" validate:"omitempty"`
		AnsweredBy primitive.ObjectID `json:"AnsweredBy" validate:"required"`
        Comment  []string `json:"Comment" validate:"omitempty"`
    }        `json:"Answers" validate:"omitempty"`
}

