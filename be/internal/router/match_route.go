package router

import (
	handler "realtime-score/internal/handlers"
	"realtime-score/internal/middleware"

	"github.com/gin-gonic/gin"
)

func MatchRouter(app *gin.Engine, matchHandler *handler.MatchHandler) {
	matchRoute := app.Group("/match")
	{
		matchRoute.PATCH("/:match_id", middleware.AuthMiddleware(), middleware.OnlyAdmin("admin"), matchHandler.UpdateMatchScore)
	}
}
