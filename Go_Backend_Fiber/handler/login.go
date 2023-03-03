package handler

import (
	"app/logs"
	"app/model"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
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
	err = H.MongoUserCol.FindOne(ctx, model.UserData{Email: loginData.Email}).Decode(&dbUser)
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
		if err == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid credentials",
			})
		} else {
			expirationTime := time.Now().Add(1000 * time.Hour)
			claims := model
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Login successful",
				"User":    dbUser,
			})
		}
	}
}
