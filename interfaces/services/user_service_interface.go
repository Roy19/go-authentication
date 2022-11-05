package interfaces

import (
	"context"
	"go-authentication/dtos"
	infraInterfaces "go-authentication/interfaces/infrastructures"
	repositoryInterfaces "go-authentication/interfaces/repository"
)

type IUserService interface {
	SetDependencies(repositoryInterfaces.IUserRepository, infraInterfaces.ILogger,
		IJwtService)
	CreateUser(context.Context,
		dtos.CreateUserDtoRequest) (dtos.CreateUserDtoResponse, error)
	LoginUser(ctx context.Context,
		loginRequest dtos.LoginRequest) (dtos.LoginResponse, error)
}
