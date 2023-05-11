package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type BlogData struct {
	ID          int    `json:"ID" bson:"_id" `
	Title       string `json:"Title" bson:"Title" validate:"required"`
	Image 	 string `json:"Image" bson:"Image" validate:"omitempty"`
	Description string `json:"Description" bson:"Description" validate:"required"`
	WrittenBy 	[]string  `json:"WrittenBy" bson:"WrittenBy" validate:"required"`  // vsn be written mode than one perosn
	WrittenTime time.Time `json:"WrittenTime" bson:"WrittenTime" validate:"omitempty"`
	EditedBy string
	EditedTime time.Time
	Tags 	  []string `json:"Tags" bson:"Tags" validate:"omitempty"`
	Comments  []string `json:"Comments" bson:"Comments" validate:"omitempty"` // list of comment id 
	Upvote 	  int `json:"Upvote" bson:"Upvote" validate:"omitempty"`
	Downvote   int `json:"Downvote" bson:"Downvote" validate:"omitempty"`
	Votes 	  [][]string `json:"Votes" bson:"Votes" validate:"omitempty"`
	Views 	  int `json:"Views" bson:"Views" validate:"omitempty"` // number of views
	ViewedBy  []string `json:"ViewedBy" bson:"ViewedBy" validate:"omitempty"`
}

type NewsModel struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	NewsID        string             `json:"NewsID" bson:"NewsID"`
	Title         string             `json:"Title" bson:"Title"`
	BannerImg     string             `json:"BannerImg" bson:"BannerImg"`
	Date          string             `json:"Date" bson:"Date"`
	WrittenBy     string             `json:"WrittenBy" bson:"WrittenBy"`
	Detail        string             `json:"Detail" bson:"Detail"`
	FooterComment string             `json:"FooterComment" bson:"FooterComment"`
}
