package utils

import (
	"app/model"
	"encoding/base64"
	"encoding/json"

	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "os"
	// "time"
	// "github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v4"
)

func GenerateHttpCookie(UserData model.UserData) string {

	cookieData := struct {
		Email       string `json:"Email"`
		Name        string `json:"Name"`
		UserID     string `json:"UserID"`
		UserURL 	  string `json:"UserURL"`
		AccountType string `json:"AccountType"`
	}{
		Email:       UserData.Email,
		Name:        UserData.UserName,
		UserID:      UserData.UserID.Hex(),
		UserURL : 	  UserData.UserURL,
		AccountType: UserData.AccountType,
	}
	jsonData, err := json.Marshal(cookieData)
	if err != nil {
		return ""
	}

	base64Data := base64.StdEncoding.EncodeToString(jsonData)
	return base64Data
}
