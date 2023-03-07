package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type BlogData struct {
	ID          int    `json:"ID" bson:"_id" `
	Title       string `json:"Title" bson:"Title" validate:"required"`
	Image 	 string `json:"Image" bson:"Image" validate:"omitempty"`
	Description string `json:"Description" bson:"Description" validate:"required"`
	WrittenBy 	[]string  `json:"WrittenBy" bson:"WrittenBy" validate:"required"`  
	WrittenTime primitive.Timestamp `json:"WrittenTime" bson:"WrittenTime" validate:"omitempty"`
	Tags 	  []string `json:"Tags" bson:"Tags" validate:"omitempty"`
	Comments  []string `json:"Comments" bson:"Comments" validate:"omitempty"`
	Upvote 	  int `json:"Upvote" bson:"Upvote" validate:"omitempty"`
	Downvote   int `json:"Downvote" bson:"Downvote" validate:"omitempty"`
	Votes 	  [][]string `json:"Votes" bson:"Votes" validate:"omitempty"`
	Views 	  int `json:"Views" bson:"Views" validate:"omitempty"`
	ViewedBy  []string `json:"ViewedBy" bson:"ViewedBy" validate:"omitempty"`
}