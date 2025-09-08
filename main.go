package main

import (
	"os"

	"w2s-backend/config"
	"w2s-backend/database"
	"w2s-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	config.Load()
	database.Connect()

	r := gin.Default()

	routes.UserRoutes(r)
	routes.AuthRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
