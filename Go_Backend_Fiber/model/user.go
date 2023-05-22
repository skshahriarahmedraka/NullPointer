package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserData struct {
	UserID        primitive.ObjectID `json:"UserID" bson:"_id"`
	UserURL    string             `json:"UserURL" bson:"UserURL" validate:"omitempty,min=4,max=64,UserIDcontainsany=~!@#$%^&*()_=+[]{}\|;:‘“,./<>?"`
	UserName  string             `json:"UserName" bson:"UserName" validate:"omitempty,min=4,max=64,UserName"`
	Email     string             `json:"Email" bson:"Email" validate:"required,min=4,max=100,Email"`
	Password  string             `json:"Password"	bson:"Password" validate:"omitempty,min=4,max=100"`
	UserTitle string             `json:"UserTitle" bson:"UserTitle" validate:"omitempty"`
	UserImage string             `json:"UserImage" bson:"UserImage" validate:"omitempty" `
	Badges    struct {
		Reputation int `json:Reputation bson:"Reputation" validate:"omitempty, numeric"`
		Gold       int `json:"Gold" bson:"Gold" validate:"omitempty,numeric"`
		Silver     int `json:"Silver" bson:"Silver" validate:"omitempty,numeric"`
		Bronze     int `json:"Bronze" 		bson:"Bronze" validate:"omitempty,numeric"`
	} `json:"Badges" bson:"Badges"`
	Follower  []string `json:"Follower" bson:"Follower" validate:"omitempty"`
	Following []string `json:"Following" bson:"Following" validate:"omitempty"`
	// Badges map[string]int
	Location       string         `json:"Location" bson:"Location" validate:"omitempty"`
	MembershipTime time.Time      `json:"MembershipTime" bson:"MembershipTime" validate:"omitempty"`
	LastSeen       time.Time      `json:"LastSeen" bson:"LastSeen" validate:"omitempty"`
	Aboutme        string         `json:"Aboutme" bson:"Aboutme" validate:"omitempty"`
	Mysite         string         `json:"Mysite" 		bson:"Mysite" validate:"omitempty"`
	Github         string         `json:"Github" bson:"Github" validate:"omitempty"`
	Twitter        string         `json:"Twitter" bson:"Twitter" validate:"omitempty"`
	Linkedin       string         `json:"Linkedin" bson:"Linkedin" validate:"omitempty"`
	TopTags        []string       `json:"TopTags" bson:"TopTags" validate:"omitempty"`
	TopTagsPercent map[string]int `json:"TopTagsPercent" bson:"TopTagsPercent" validate:"omitempty"`
	SelectedPanel  string         `json:"SelectedPanel" bson:"SelectedPanel" validate:"omitempty"`
	AccountType    string         `json:"AccountType" bson:"AccountType" validate:"omitempty"`
}

type UserDataFlair struct {
	UserID        primitive.ObjectID `json:"UserID" bson:"_id"`
	UserURL   string             `json:"UserURL" bson:"UserURL" validate:"omitempty,min=4,max=64,UserIDcontainsany=~!@#$%^&*()–_=+[]{}\|;:‘“,./<>?"`
	UserName  string             `json:"UserName" bson:"UserName" validate:"omitempty,min=4,max=64,UserName"`
	UserImage string             `json:"UserImage" bson:"UserImage" validate:"omitempty" `
	Badges    struct {
		Reputation int `json:Reputation bson:"Reputation" validate:"omitempty, numeric"`
		Gold       int `json:"Gold" bson:"Gold" validate:"omitempty,numeric"`
		Silver     int `json:"Silver" bson:"Silver" validate:"omitempty,numeric"`
		Bronze     int `json:"Bronze" 		bson:"Bronze" validate:"omitempty,numeric"`
	} `json:"Badges" bson:"Badges"`
	Location string `json:"Location" bson:"Location" validate:"omitempty"`
	Aboutme  string `json:"Aboutme" bson:"Aboutme" validate:"omitempty"`
}
