package handler

import (
	"net/http"
	"realtime-score/internal/dto"
	"realtime-score/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MatchHandler struct {
	log          *logrus.Logger
	MatchService *services.MatchService
}

func NewMatchHandler(matchService *services.MatchService, log *logrus.Logger) *MatchHandler {
	return &MatchHandler{
		log:          log,
		MatchService: matchService,
	}
}

func (h *MatchHandler) UpdateMatchScore(c *gin.Context) {
	match_id := c.Param("match_id")
	var update dto.UpdateScore
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed update score",
			"message": err.Error(),
		})
		return
	}
	match, err := h.MatchService.UpdateMatchScore(match_id, update.ScoreA, update.ScoreB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed update score",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Score updated",
		"match":   match,
	})
}
