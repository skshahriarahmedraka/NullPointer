package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type HashData struct {
	ID primitive.ObjectID  `json:"ID" bson:"_id"`
	Name  string `json:"Name" bson:"Name"`
	MetaTitle  string `json:"MetaTitle" bson:"MetaTitle"`
	Image  string `json:"Image" bson:"Image"`
 	BannerImage  string `json:"BannerImage" bson:"BannerImage"`
	NumOfFollower int `json:"NumOfFollower" bson:"NumOfFollower"`
	NumOfQuestion int `json:"NumOfQuestion" bson:"NumOfQuestion"`
	NumOfBlog int `json:"NumOfBlog" bson:"NumOfBlog"`
	NumOfAnswer int `json:"NumOfAnswer" bson:"NumOfAnswer"`
	About  string `json:"About" bson:"About"`

}

type HashFlair struct {
	ID primitive.ObjectID `json:"ID" bson:"_id"`
	Image string `json:"Image" bson:"Image"`
	Name string `json:"Name" bson:"Name"`
}