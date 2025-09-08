package main

import (
	"os"

	"w2s-backend/config"
	"w2s-backend/database"
	"w2s-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	_ = godotenv.Load()

	config.Load()
	database.Connect()

	r := gin.Default()

	r.StaticFile("/swagger.json", "./docs/swagger.json")

	url := ginSwagger.URL("/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	routes.UserRoutes(r)
	routes.AuthRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
