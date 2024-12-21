package router

import (
	handler "realtime-score/internal/handlers"
	"realtime-score/internal/middleware"

	"github.com/gin-gonic/gin"
)

func User(app *gin.Engine, userHandler handler.UserHandler) {
	userRoutes := app.Group("/users")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.GET("/", middleware.AuthMiddleware(), userHandler.Me)
		userRoutes.GET("/all", middleware.AuthMiddleware(), middleware.OnlyAdmin("admin"), userHandler.GetAllUser)
		userRoutes.PATCH("/", middleware.AuthMiddleware(), userHandler.UpdateUser)
		userRoutes.DELETE("/", middleware.AuthMiddleware(), userHandler.DeleteUser)
	}
}
