package model

import "github.com/golang-jwt/jwt/v4"

type JwtAuth1 struct {
	Email   string `json:"Email" validate:"required,email"`
	Name   string `json:"Name"  validate:"required"`
	UserID string `json:"UserID"  validate:"required"`
	AccountType string `json:"AccountType"  validate:"required"`
	UUID string `json:"UUID"  validate:"required"`
	jwt.StandardClaims
}