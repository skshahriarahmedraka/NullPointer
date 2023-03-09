package model

type QuesData struct {
	ID                  int      `json:"ID" bson:"_id" validate:"required"`
	QuesTitle       string   `json:"QuesTitle" bson:"QuesTitle" validate:"required,min=4,max=200"`
	QuesitonAskedTime   string   `json:"QuesAskedTime" bson:"QuesAskedTime" validate:"omitempty"`
	QuesModified    string   `json:"QuesModified" bson:"QuesModified" validate:"omitempty"`
	QuesViewed      int      `json:"QuesViewed" bson:"QuesViewed" validate:"omitempty"`
	QuesUpvote      int      `json:"QuesUpvote" bson:"QuesUpvote" validate:"omitempty"`
	QuesDownvote    int      `json:"QuesDownvote" bson:"QuesDownvote" validate:"omitempty"`
	QuesBookmark    int      `json:"QuesBookmark" bson:"QuesBookmark" validate:"omitempty"`
	QuesTags        []string `json:"QuesTags" bson:"QuesTags" validate:"omitempty"`
	QuesAskedBy     string   `json:"QuesAskedBy" bson:"QuesAskedBy" validate:"required"`
	QuesAnsAccepted string   `json:"QuesAnsAccepted" bson:"QuesAnsAccepted" validate:"required"`

	// QuesAskedByName    string   `json:"QuesAskedByName" bson:"QuesAskedByName" validate:"required"`
	// QuesAskedByBadge    struct {
	// 	Reputation int `json:Reputation bson:"Reputation" validate:"omitempty, numeric"`
	// 	Gold       int `json:"Gold" bson:"Gold" validate:"omitempty,numeric"`
	// 	Silver     int `json:"Silver" bson:"Silver" validate:"omitempty,numeric"`
	// 	Bronze     int `json:"Bronze" 		bson:"Bronze" validate:"omitempty,numeric"`
	// } `json:"QuesAskedByBadge" bson:"QuesAskedByBadge"`
	QuesAskedTimeExact string `json:"QuesAskedTimeExact" bson:"QuesAskedTimeExact" validate:"omitempty"`
	QuesAskedByImage   string `json:"QuesAskedByImage" bson:"QuesAskedByImage" validate:"omitempty"`

	QuesEditedBy string `json:"QuesEditedBy" bson:"QuesEditedBy" validate:"omitempty"`

	QuesEditedTimeExact string                `json:"QuesEditedTimeExact" bson:"QuesEditedTimeExact" validate:"omitempty"`
	QuesDescription     string                `json:"QuesDescription" bson:"QuesDescription" validate:"required,min=10,max=10000"`
	QuesComment         [][]string  `json:"QuesComment" bson:"QuesComment" validate:"omitempty"`
	Answers                 [][]string            `json:"Answers" validate:"omitempty"`
}
