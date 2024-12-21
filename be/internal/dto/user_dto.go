package dto

import (
	"errors"
	"realtime-score/internal/models"
)

var (
	EmailAlreadyExist    = errors.New("email already exist")
	UsernameAlreadyExist = errors.New("username already exist")
	CantCreateUser       = errors.New("cant create user")
	InvalidCredentials   = errors.New("wrong password or email")
	UserNotFound         = errors.New("user not found")
	ErrGetAllUser        = errors.New("error get all user")
)

type (
	UserCreateRequest struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	UserResponse struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	UserLoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	UserLoginResponse struct {
		Token string `json:"token"`
		Role  string `json:"role"`
	}
	GetAllUserResponse struct {
		Users []models.User
	}
	UserUpdateRequest struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
)
