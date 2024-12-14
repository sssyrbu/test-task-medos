package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sssyrbu/test-task-medos/db"
	"github.com/sssyrbu/test-task-medos/handlers"
)

func main() {
	db.InitializeDatabase()

	r := gin.Default()

	r.POST("/auth/token", handlers.GenerateToken)
	r.POST("/auth/refresh", handlers.RefreshToken)

	log.Fatal(r.Run(":8888"))
}
