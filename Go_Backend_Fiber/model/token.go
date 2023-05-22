package model

import (
	"github.com/golang-jwt/jwt/v4"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtAuth1 struct {
	Email   string `json:"Email" validate:"required,email"`
	Name   string `json:"Name"  validate:"required"`
	UserURL string `json:"UserURL"  validate:"required"`
	AccountType string `json:"AccountType"  validate:"required"`
	UserID string `json:"UserID"  validate:"required"`
	jwt.StandardClaims
}