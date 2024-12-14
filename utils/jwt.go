package utils

import (
	"os"
	"time"

	"test-task-medos/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))
	if len(jwtKey) == 0 {
		panic("JWT_SECRET_KEY не установлен")
	}
}

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
