package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type UserActivity struct {
	ID primitive.ObjectID `json:"ID" bson:"_id"`
	// UserID primitive.ObjectID `json:"UserID" bson:"UserID"`
	FollowingHash []HashFlair `json:"FollowingHash" bson:"FollowingHash"`
	FollowingUser []UserDataFlair `json:"FollowingUser" bson:"FollowingUser"`
	FollowingQuestion []QuesFlair `json:"FollowingQuestion" bson:"FollowingQuestion"`
}



