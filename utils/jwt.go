package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sssyrbu/test-task-medos/models"
)

var jwtKey = os.Getenv("JWT_SECRET_KEY")

func GenerateJWT(userID, ip string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &models.TokenClaims{
		UserID: userID,
		IP:     ip,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString(jwtKey)
}

func VerifyJWT(tokenString string) (*models.TokenClaims, error) {
	claims := &models.TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	return claims, nil
}
