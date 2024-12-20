package handler

import (
	"net/http"
	"realtime-score/internal/dto"
	"realtime-score/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	log         *logrus.Logger
	userService services.UserService
}

func NewUser(log *logrus.Logger, userService services.UserService) UserHandler {
	return UserHandler{
		log:         log,
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Registration failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Registration successful",
		"data":    result,
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.userService.Login(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Login failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"data":    result,
	})
}

func (h *UserHandler) Me(c *gin.Context) {
	user_id := c.MustGet("user_id").(string)

	result, err := h.userService.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed get user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get user",
		"data":    result,
	})
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	result, err := h.userService.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "failed get all user",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success get all user",
		"data":    result,
	})
}
