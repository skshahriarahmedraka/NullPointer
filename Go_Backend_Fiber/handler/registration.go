package handler

import (
	"app/logs"
	"app/model"
	"app/utils"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	// "hash"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (H *DatabaseCollections) Register(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var RegData model.UserData
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&RegData)
	logs.Error("Request decoding error", err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	fmt.Println("User Request :", RegData)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var dbUser model.UserData
	err = H.MongoUserCol.FindOne(ctx, bson.D{{Key: "Email", Value: RegData.Email}}).Decode(&dbUser)
	if err == nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "User Already Exists",
		})
	} else {
		// Check if the password is correct

		RegData.UserID = primitive.NewObjectID()
		RegData.UserURL = "/"+RegData.UserID.Hex()
		hash, err := bcrypt.GenerateFromPassword([]byte(RegData.Password), bcrypt.DefaultCost)
		RegData.Password = string(hash)
		// UserTitle string `json:"UserTitle"`
		RegData.UserImage = ""
		RegData.Badges = struct {
			Reputation int `json:Reputation bson:"Reputation" validate:"omitempty, numeric"`
			Gold       int `json:"Gold" bson:"Gold" validate:"omitempty,numeric"`
			Silver     int `json:"Silver" bson:"Silver" validate:"omitempty,numeric"`
			Bronze     int `json:"Bronze" 		bson:"Bronze" validate:"omitempty,numeric"`
		}{
			Reputation: 1,
			Gold:       0,
			Silver:     0,
			Bronze:     0,
		}
		RegData.Follower = []string{}
		RegData.Following = []string{}
		// Badges map[string]int
		RegData.Location = ""
		RegData.MembershipTime =time.Now().UTC()
		RegData.LastSeen = time.Now().UTC()
		RegData.Aboutme = ""
		RegData.Mysite = ""
		RegData.Github = ""
		RegData.Twitter = ""
		RegData.Linkedin = ""
		RegData.TopTags = []string{}
		RegData.TopTagsPercent = map[string]int{}
		RegData.SelectedPanel = "profile"
		RegData.AccountType = "normal"

		res, err := H.MongoUserCol.InsertOne(ctx, RegData)
		logs.Error("Error while inserting user", err)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error while inserting user",
			})
		}
		fmt.Println("🚀 inserted UserData : ", res)

	
		tokenString := utils.GenerateHttpOnlyJWT(RegData)
		if tokenString == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error while creating token",
			})
		}
		
		c.Cookie(&fiber.Cookie{
			Name:     "Auth1",
			Value:    tokenString,
			Path:     "/",
			Domain:   os.Getenv("IP"),
			Expires:  time.Now().Add(1000 * time.Hour),
			HTTPOnly: true,

			SameSite: "lax",
		})

		tokenString2 := utils.GenerateHttpCookie(RegData)
		if tokenString2 == "" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Error while creating token2",
			})
		}
		c.Cookie(&fiber.Cookie{
			Name:     "Info1",
			Value:    tokenString2,
			Path:     "/",
			Domain:   os.Getenv("IP"),
			Expires:  time.Now().Add(1000 * time.Hour),
			HTTPOnly: false,

			SameSite: "lax",
		})
		RegData.Password=""
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Register successful",
			"User":    RegData,
		})

	}
}
