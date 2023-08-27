package model

import (

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// "go.mongodb.org/mongo-driver/mongo"
// "gorm.io/gorm"

// type Article struct {
// 	ID      int    `json:"id"`
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }

// var ArticleList = []Article{
// 	{ID: 123, Title: "harry potter", Content: "amazing story book"},
// 	{ID: 15, Title: "One piece", Content: "See anime its awesome "},
// }

// func GetAllTheArticle() []Article {
// 	return ArticleList
// }

// func GetArticleByID(a int) (*Article, error) {
// 	for _, i := range ArticleList {
// 		if i.ID == a {
// 			return &i, nil
// 		}
// 	}
// 	return nil, errors.New("Article was not found ")
// }

type PostgresConfig struct {
	HOST     string
	PORT     string
	PASSWORD string
	USER     string
	DBNAME   string
	SSLMODE  string
	TIMEZONE string
}

type UserData struct {
	ID   primitive.ObjectID `json:"ID" bson:"_id"`
	UserID string `json:"UserID" bson:"UserID"`
	UserName string `json:"UserName" bson:"UserName"`
	Email    string `json:"Email" bson:"Email"`
	Password string `json:"Password"	bson:"Password"`
	// UserTitle string `json:"UserTitle"`
	UserImage string `json:"UserImage" bson:"UserImage"`
	Badges    struct {
		Reputation int `json:Reputation bson:"Reputation"`
		Gold       int `json:"Gold" bson:"Gold"`
		Silver     int `json:"Silver" bson:"Silver"`
		Bronze     int `json:"Bronze" 		bson:"Bronze"`
	} `json:"Badges" bson:"Badges"`
	Follower []string `json:"Follower" json:"Follower"`
	Following []string `json:"Following" json:"Following"`
	// Badges map[string]int
	Location       string   `json:"Location" bson:"Location"`
	MembershipTime primitive.Timestamp   `json:"MembershipTime" bson:"MembershipTime"`
	LastSeen       primitive.Timestamp   `json:"LastSeen" bson:"LastSeen"`
	Aboutme        string   `json:"Aboutme" bson:"Aboutme"`
	Mysite         string   `json:"Mysite" 		bson:"Mysite"`
	Github         string   `json:"Github" bson:"Github"`
	Twitter        string   `json:"Twitter" bson:"Twitter"`
	Linkedin       string   `json:"Linkedin" bson:"Linkedin"`
	TopTags        []string `json:"TopTags" bson:"TopTags"`
	TopTagsPercent []int    `json:"TopTagsPercent" bson:"TopTagsPercent"`
	SelectedPanel  string   `json:"SelectedPanel" bson:"SelectedPanel"`
	Accounttype    string   `json:"Accounttype" bson:"Accounttype"`
}

type QuestionData struct {
	QuestionTitle          string   `json:"QuestionTitle"`
	QuesitonAskedTime      string   `json:"QuestionAskedTime"`
	QuestionModified       string   `json:"QuestionModified"`
	QuestionViewed         int      `json:"QuestionViewed"`
	QuestionUpvote         int      `json:"QuestionUpvote"`
	QuestionDownvote       int      `json:"QuestionDownvote"`
	QuestionBookmark       int      `json:"QuestionBookmark"`
	QuestionTags           []string `json:"QuestionTags"`
	QuestionAskedBy        string   `json:"QuestionAskedBy"`
	QuestionAskedByName    string   `json:"QuestionAskedByName"`
	QuestionAskedByBadge   []int    `json:"QuestionAskedByBadge"`
	QuestionAskedTimeExact string   `json:"QuestionAskedTimeExact"`
	QuestionAskedByImage   string   `json:"QuestionAskedByImage"`

	QuestionEditedBy        string          `json:"QuestionEditedBy"`
	QuestionEditedByName    string          `json:"QuestionEditedByName"`
	QuestionEditedByBadge   []int           `json:"QuestionEditedByBadge"`
	QuestionEditedByImage   string          `json:"QuestionEditedByImage"`
	QuestionEditedTimeExact string          `json:"QuestionEditedTimeExact"`
	QuestionDescription     string          `json:"QuestionDescription"`
	QuestionComment         [][]string      `json:"QuestionComment"`
	Answers                 [][]interface{} `json:"Answers"`
}

type RelatedQuestionList [][]interface{}

type SpaceList [][]string

type FavoriteHash [][]string

type ListOfBlog [][]string
type HotQuestions [][]string

type PublicQues [][]interface{}
