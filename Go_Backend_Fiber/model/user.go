package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type UserData struct {
	ID   primitive.ObjectID `json:"ID" bson:"_id"`
	UserID string `json:"UserID" bson:"UserID" validate:"omitempty,min=4,max=64,UserIDcontainsany=~!@#$%^&*()–_=+[]{}\|;:‘“,./<>?"`
	UserName string `json:"UserName" bson:"UserName" validate:"omitempty,min=4,max=64,UserName"` 
	Email    string `json:"Email" bson:"Email" validate:"required,min=4,max=100,Email"`
	Password string `json:"Password"	bson:"Password" validate:"omitempty,min=4,max=100"`
	// UserTitle string `json:"UserTitle"`
	UserImage string `json:"UserImage" bson:"UserImage" validator:"omitempty" `
	Badges    struct {
		Reputation int `json:Reputation bson:"Reputation" validator:"omitempty,numeric"`
		Gold       int `json:"Gold" bson:"Gold" validator:"omitempty,numeric"`
		Silver     int `json:"Silver" bson:"Silver" validator:"omitempty,numeric"`
		Bronze     int `json:"Bronze" 		bson:"Bronze" validator:"omitempty,numeric"`
	} `json:"Badges" bson:"Badges"`
	Follower []string `json:"Follower" json:"Follower" validator:"omitempty"`
	Following []string `json:"Following" json:"Following" validator:"omitempty"`
	// Badges map[string]int
	Location       string   `json:"Location" bson:"Location" validator:"omitempty"`
	MembershipTime primitive.Timestamp   `json:"MembershipTime" bson:"MembershipTime" validator:"omitempty"`
	LastSeen       primitive.Timestamp   `json:"LastSeen" bson:"LastSeen" validator:"omitempty"`
	Aboutme        string   `json:"Aboutme" bson:"Aboutme" validator:"omitempty"`
	Mysite         string   `json:"Mysite" 		bson:"Mysite" validator:"omitempty"`
	Github         string   `json:"Github" bson:"Github" validator:"omitempty"`
	Twitter        string   `json:"Twitter" bson:"Twitter" validator:"omitempty"`
	Linkedin       string   `json:"Linkedin" bson:"Linkedin" validator:"omitempty"`
	TopTags        []string `json:"TopTags" bson:"TopTags" validator:"omitempty"`
	TopTagsPercent []int    `json:"TopTagsPercent" bson:"TopTagsPercent" validator:"omitempty"`
	SelectedPanel  string   `json:"SelectedPanel" bson:"SelectedPanel" validator:"omitempty"`
	Accounttype    string   `json:"Accounttype" bson:"Accounttype" validator:"omitempty"`
}

