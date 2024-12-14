package models

import "github.com/dgrijalva/jwt-go"

type TokenClaims struct {
	UserID string `json:"user_id"`
	IP     string `json:"ip"`
	jwt.StandardClaims
}