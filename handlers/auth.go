package handlers

import (
	"fmt"
	"net/http"
	"test-task-medos/db"
	"test-task-medos/utils"

	"github.com/gin-gonic/gin"
)

func GenerateAccessAndRefreshTokens(c *gin.Context) {
	userID := c.Query("user_id")
	ip := c.ClientIP()

	accessToken, err := utils.GenerateJWT(userID, ip)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации access токена"})
		return
	}

	refreshToken := utils.GenerateRefreshToken()
	hashedRefreshToken, err := utils.HashToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хэшировании refresh токена"})
		return
	}

	err = db.SaveRefreshToken(userID, hashedRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранее refresh токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c *gin.Context) {
	var request struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неправильный запрос"})
		return
	}

	ip := c.ClientIP()

	claims, err := utils.VerifyJWT(request.AccessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный access токен"})
		return
	}

	hashedRefreshToken, err := db.GetRefreshToken(claims.UserID)
	if err != nil || !utils.CompareTokens(request.RefreshToken, hashedRefreshToken) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неправильный refresh токен"})
		return
	}

	if claims.IP != ip {
		sendEmailWarning(claims.UserID)
	}

	newAccessToken, err := utils.GenerateJWT(claims.UserID, ip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации нового access токена"})
		return
	}

	newRefreshToken := utils.GenerateRefreshToken()
	newHashedRefreshToken, err := utils.HashToken(newRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хэшировании нового refresh токена"})
		return
	}

	err = db.UpdateRefreshToken(claims.UserID, newHashedRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении refresh токена"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}

func sendEmailWarning(userID string) {
	println("Мокаем отправку варнинга на email пользователя с айди", userID)
}
