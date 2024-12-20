package services

import (
	"realtime-score/internal/dto"
	"realtime-score/internal/models"
	"realtime-score/internal/pkg"
	"realtime-score/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUser(userRepo repository.UserRepository) UserService {
	return UserService{userRepo: userRepo}
}

func (s *UserService) Register(req dto.UserCreateRequest) (dto.UserResponse, error) {
	isEmailExist := s.userRepo.IsEmailExist(req.Email)
	if isEmailExist {
		return dto.UserResponse{}, dto.EmailAlreadyExist
	}

	isUsernameExist := s.userRepo.IsUsernameExist(req.Username)
	if isUsernameExist {
		return dto.UserResponse{}, dto.UsernameAlreadyExist
	}
	hashedPassword, err := pkg.HashPassword(req.Password)
	data := models.User{
		Username:   req.Username,
		Email:      req.Email,
		Password:   hashedPassword,
		IsVerified: true,
		Role:       "user",
	}

	register, err := s.userRepo.CreateUser(data)
	if err != nil {
		return dto.UserResponse{}, dto.CantCreateUser

	}
	return dto.UserResponse{
		Username: register.Username,
		Email:    register.Email,
	}, nil
}

func (s *UserService) Login(req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	checkUserExist, err := s.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return dto.UserLoginResponse{}, dto.InvalidCredentials
	}

	checkPass := pkg.CheckPassword(req.Password, checkUserExist.Password)
	if !checkPass {
		return dto.UserLoginResponse{}, dto.InvalidCredentials
	}

	token, err := pkg.GenerateToken(checkUserExist.ID, checkUserExist.Role)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	return dto.UserLoginResponse{
		Token: token,
		Role:  checkUserExist.Role,
	}, nil
}

func (s *UserService) GetUserByID(userid string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(userid)
	if err != nil {
		return dto.UserResponse{}, dto.UserNotFound
	}
	return dto.UserResponse{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *UserService) GetAllUser() (dto.GetAllUserResponse, error) {
	users, err := s.userRepo.GetAllUser()
	if err != nil {
		return dto.GetAllUserResponse{}, dto.ErrGetAllUser
	}
	return users, nil
}
