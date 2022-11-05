package services

import (
	"context"
	"errors"
	"go-authentication/dtos"
	infraInterfaces "go-authentication/interfaces/infrastructures"
	repositoryInterfaces "go-authentication/interfaces/repository"
	servicesInterfaces "go-authentication/interfaces/services"
	"go-authentication/models"
)

type UserService struct {
	userRepo   repositoryInterfaces.IUserRepository
	logger     infraInterfaces.ILogger
	jwtService servicesInterfaces.IJwtService
}

func (us *UserService) SetDependencies(userRepo repositoryInterfaces.IUserRepository,
	logger infraInterfaces.ILogger, jwtService servicesInterfaces.IJwtService) {
	us.userRepo = userRepo
	us.logger = logger
	us.jwtService = jwtService
}

func (us *UserService) CreateUser(ctx context.Context,
	createUserDto dtos.CreateUserDtoRequest) (dtos.CreateUserDtoResponse, error) {
	if us.userRepo == nil {
		us.logger.Err(ctx, "invalid dependency: userRepo in UserService",
			errors.New("invalid dependency: userRepo in UserService"))
	}
	user := models.User{
		Email:    createUserDto.Email,
		UserName: createUserDto.UserName,
		Password: createUserDto.Password,
	}
	result, err := us.userRepo.CreateUser(ctx, &user)
	if result {
		return dtos.CreateUserDtoResponse{
			ID:       user.ID,
			Email:    user.Email,
			UserName: user.UserName,
		}, nil
	} else {
		return dtos.CreateUserDtoResponse{}, err
	}
}

func (us *UserService) LoginUser(ctx context.Context,
	loginRequest dtos.LoginRequest) (dtos.LoginResponse, error) {
	if us.userRepo == nil {
		us.logger.Err(ctx, "invalid dependency: userRepo in UserService",
			errors.New("invalid dependency: userRepo in UserService"))
	}
	user := models.User{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	}
	result, err := us.userRepo.FindUser(ctx, &user)
	if result {
		token, err := us.jwtService.CreateNewToken(ctx, user.ID)
		if err != nil {
			us.logger.Err(ctx, "Failed to generate token", err)
			return dtos.LoginResponse{}, err
		} else {
			return dtos.LoginResponse{
				UserName: user.UserName,
				Email:    user.Email,
				Token:    token,
			}, nil
		}
	} else {
		return dtos.LoginResponse{}, err
	}
}
