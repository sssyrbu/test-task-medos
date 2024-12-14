package main

import (
	"log"

	"test-task-medos/db"
	"test-task-medos/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDatabase()

	r := gin.Default()

	r.POST("/auth/token", handlers.GenerateAccessAndRefreshTokens)
	r.POST("/auth/refresh", handlers.RefreshToken)

	log.Fatal(r.Run(":8888"))
}
