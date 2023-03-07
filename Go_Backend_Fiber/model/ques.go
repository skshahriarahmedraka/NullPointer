package model


type QuestionData struct {
	ID 					int      `json:"ID" bson:"_id" validate:"required"`
	QuestionTitle          string   `json:"QuestionTitle" bson:"QuestionTitle" validate:"required,min=4,max=200"`
	QuesitonAskedTime      string   `json:"QuestionAskedTime" bson:"QuestionAskedTime" validate:"omitempty"`
	QuestionModified       string   `json:"QuestionModified" bson:"QuestionModified" validate:"omitempty"`
	QuestionViewed         int      `json:"QuestionViewed" bson:"QuestionViewed" validate:"omitempty"`
	QuestionUpvote         int      `json:"QuestionUpvote" bson:"QuestionUpvote" validate:"omitempty"`
	QuestionDownvote       int      `json:"QuestionDownvote" bson:"QuestionDownvote" validate:"omitempty"`
	QuestionBookmark       int      `json:"QuestionBookmark" bson:"QuestionBookmark" validate:"omitempty"`
	QuestionTags           []string `json:"QuestionTags" bson:"QuestionTags" validate:"omitempty"`
	QuestionAskedBy        string   `json:"QuestionAskedBy" bson:"QuestionAskedBy" validate:"required"`
	// QuestionAskedByName    string   `json:"QuestionAskedByName" bson:"QuestionAskedByName" validate:"required"`
	// QuestionAskedByBadge    struct {
	// 	Reputation int `json:Reputation bson:"Reputation" validate:"omitempty, numeric"`
	// 	Gold       int `json:"Gold" bson:"Gold" validate:"omitempty,numeric"`
	// 	Silver     int `json:"Silver" bson:"Silver" validate:"omitempty,numeric"`
	// 	Bronze     int `json:"Bronze" 		bson:"Bronze" validate:"omitempty,numeric"`
	// } `json:"QuestionAskedByBadge" bson:"QuestionAskedByBadge"`   
	QuestionAskedTimeExact string   `json:"QuestionAskedTimeExact" bson:"QuestionAskedTimeExact" validate:"omitempty"`
	QuestionAskedByImage   string   `json:"QuestionAskedByImage" bson:"QuestionAskedByImage" validate:"omitempty"`

	QuestionEditedBy        string          `json:"QuestionEditedBy" bson:"QuestionEditedBy" validate:"omitempty"`

	QuestionEditedTimeExact string          `json:"QuestionEditedTimeExact" bson:"QuestionEditedTimeExact" validate:"omitempty"`
	QuestionDescription     string          `json:"QuestionDescription" bson:"QuestionDescription" validate:"required,min=10,max=10000"`
	QuestionComment         [][]map[string]string      `json:"QuestionComment" bson:"QuestionComment" validate:"omitempty"`
	Answers                 [][]string  `json:"Answers" validate:"omitempty"`
}
