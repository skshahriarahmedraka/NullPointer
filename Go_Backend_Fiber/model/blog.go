package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type BlogData struct {
	ID          primitive.ObjectID   `json:"ID" bson:"_id" `
	Title       string `json:"Title" bson:"Title" validate:"required"`
	MetaTitle  string `json:"MetaTitle" bson:"MetaTitle" validate:"required"`
	Image 	 string `json:"Image" bson:"Image" validate:"omitempty"`
	Description string `json:"Description" bson:"Description" validate:"required"`
	WrittenBy 	[]primitive.ObjectID  `json:"WrittenBy" bson:"WrittenBy" validate:"required"`  // vsn be written mode than one perosn
	WrittenTime time.Time `json:"WrittenTime" bson:"WrittenTime" validate:"omitempty"`
	EditedBy string `json:"EditedBy" bson:"EditedBy" validate:"omitempty"`
	EditedTime time.Time `json:"EditedTime" bson:"EditedTime" validate:"omitempty"`
	Tags 	  []string `json:"Tags" bson:"Tags" validate:"omitempty"`
	Comments  []Comment `json:"Comments" bson:"Comments" validate:"omitempty"`
	Upvote 	  int `json:"Upvote" bson:"Upvote" validate:"omitempty"`
	Downvote   int `json:"Downvote" bson:"Downvote" validate:"omitempty"`
	Votes 	  []BlogVoteData `json:"Votes" bson:"Votes" validate:"omitempty"`
	Bookmark  []BlogBookmarkData `json:"Bookmark" bson:"Bookmark" validate:"omitempty"`
	Views 	  int `json:"Views" bson:"Views" validate:"omitempty"` // number of views
	ViewedBy  []string `json:"ViewedBy" bson:"ViewedBy" validate:"omitempty"`
}

type BlogVoteData struct {
	ID          primitive.ObjectID   `json:"ID" bson:"_id" `
	UserID	  primitive.ObjectID `json:"UserID" bson:"UserID" validate:"required"`
	Vote int `json:"Vote" bson:"Vote" validate:"omitempty"`
	VoteTime time.Time `json:"VoteTime" bson:"VoteTime" validate:"omitempty"`

}
type BlogBookmarkData struct {
	ID 		primitive.ObjectID   `json:"ID" bson:"_id" `
	UserID	  primitive.ObjectID `json:"UserID" bson:"UserID" validate:"required"`
	BookmarkTime time.Time `json:"BookmarkTime" bson:"BookmarkTime" validate:"omitempty"`
}

type BlogFlairData struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	NewsID        string             `json:"NewsID" bson:"NewsID"`
	Title         string             `json:"Title" bson:"Title"`
	BannerImg     string             `json:"BannerImg" bson:"BannerImg"`
	Date          string             `json:"Date" bson:"Date"`
	WrittenBy     string             `json:"WrittenBy" bson:"WrittenBy"`
	Detail        string             `json:"Detail" bson:"Detail"`
	FooterComment string             `json:"FooterComment" bson:"FooterComment"`
}
