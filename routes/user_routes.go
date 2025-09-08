package routes

import (
	"w2s-backend/handlers"
	"w2s-backend/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/api/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/profile", handlers.GetProfile)
		user.PUT("/update-profile", handlers.UpdateProfile)
		user.PUT("/update-email", handlers.UpdateEmail)
	}
}
