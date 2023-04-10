package model

type QuesData struct {
	ID                  int      `json:"ID" bson:"_id" validate:"required"`
	QuesTitle       string    `json:"QuesTitle" bson:"QuesTitle" validate:"required,min=4,max=200"`
	
	QuesViewed      int      `json:"QuesViewed" bson:"QuesViewed" validate:"omitempty"`
	QuesUpvote      int      `json:"QuesUpvote" bson:"QuesUpvote" validate:"omitempty"`
	QuesDownvote    int      `json:"QuesDownvote" bson:"QuesDownvote" validate:"omitempty"`
	QuesBookmark    int      `json:"QuesBookmark" bson:"QuesBookmark" validate:"omitempty"`
	QuesTags        []string `json:"QuesTags" bson:"QuesTags" validate:"omitempty"`
	QuesAnsAccepted string   `json:"QuesAnsAccepted" bson:"QuesAnsAccepted" validate:"required"`
	
	
	QuesAskedBy     string   `json:"QuesAskedBy" bson:"QuesAskedBy" validate:"required"`
	QuesAskedTime string `json:"QuesAskedTime" bson:"QuesAskedTime" validate:"omitempty"`
	

	QuesEditedBy string `json:"QuesEditedBy" bson:"QuesEditedBy" validate:"omitempty"`
	QuesEditedTime string                `json:"QuesEditedTime" bson:"QuesEditedTime" validate:"omitempty"`

	QuesDescription     string                `json:"QuesDescription" bson:"QuesDescription" validate:"required,min=10,max=10000"`
	QuesComment         []string  `json:"QuesComment" bson:"QuesComment" validate:"omitempty"`
	Answers                 []struct{ 
        ID  string `json:"ID" validate:"required"`
        Vote int `json:"Vote" validate:"omitempty"`
        Comment  []string `json:"Comment" validate:"omitempty"`
    }        `json:"Answers" validate:"omitempty"`
}

