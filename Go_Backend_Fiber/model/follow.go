package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type HashViewData struct {
	ID  primitive.ObjectID `json:"ID" bson:"_id"`
	UserID primitive.ObjectID `json:"UserID" bson:"UserID"`
	Image string `json:"Image" bson:"Image"`
	Name string `json:"Name" bson:"Name"`
	FollowStatus bool `json:"FollowStatus" bson:"FollowStatus"`
}