package utils

import (
	"app/model"
	"os"
	"time"

	// "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateHttpOnlyJWT(UserData model.UserData) string {
	expirationTime := time.Now().Add(1000 * time.Hour)
			claims := model.JwtAuth1{
				Email:        UserData.Email,
				Name:         UserData.UserName,
				UserURL:       UserData.UserURL,
				UserID : 	  UserData.UserID.Hex(),
				AccountType:  UserData.AccountType,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte(os.Getenv("COOKIE_SECRET_JWT_AUTH1")))
			if err != nil {
				return ""
			}
			return tokenString
}