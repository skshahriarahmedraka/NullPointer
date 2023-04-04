package utils

import (
	"app/model"
	"time"

	// "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateHttpCookie(UserData model.UserData) string {
	expirationTime := time.Now().Add(1000 * time.Hour)
			claims := model.JwtAuth1{
				Email:        UserData.Email,
				Name:         UserData.UserName,
				UserID:       UserData.UserID,
				AccountType:  UserData.AccountType,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString([]byte("COOKIE_SECRET_NOAUTH1"))
			if err != nil {
				return ""
			}
			return tokenString
}