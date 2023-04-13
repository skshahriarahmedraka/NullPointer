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
	"time"

	"github.com/gofiber/fiber/v2"

	// "github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (H *DatabaseCollections) Login(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var loginData model.LoginModel
	err := json.NewDecoder(bytes.NewReader(c.Body())).Decode(&loginData)
	logs.Error("Request decoding error", err)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	fmt.Println("Login Request :", loginData)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var dbUser model.UserData
	err = H.MongoUserCol.FindOne(ctx, bson.D{{"Email", loginData.Email}}).Decode(&dbUser)

	
	// Check if the user exists
	logs.Error("User Email not found in database :", err)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	} else {
		// Check if the password is correct
		userPass := []byte(loginData.Password)
		dbPass := []byte(dbUser.Password)
		err := bcrypt.CompareHashAndPassword(dbPass, userPass)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid credentials",
			})
		} else {
			// expirationTime := time.Now().Add(1000 * time.Hour)
			// claims := model.JwtAuth1{
			// 	Email:        dbUser.Email,
			// 	Name:         dbUser.UserName,
			// 	UserID:       dbUser.UserID,
			// 	AccountType:  dbUser.AccountType,
			// 	StandardClaims: jwt.StandardClaims{
			// 		ExpiresAt: expirationTime.Unix(),
			// 	},
			// }
			// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// tokenString, err := token.SignedString([]byte("COOKIE_SECRET_JWT_AUTH1"))
			// if err != nil {
			// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			// 		"message": "Error while creating token",
			// 	})
			// }
			tokenString := utils.GenerateHttpOnlyJWT(dbUser)
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

			tokenString2 := utils.GenerateHttpCookie(dbUser)
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
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Login successful",
				"User":    dbUser,
			})
		}
	}
}
