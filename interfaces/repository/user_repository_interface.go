package interfaces

import (
	"context"
	interfaces "go-authentication/interfaces/infrastructures"
	"go-authentication/models"

	"gorm.io/gorm"
)

type IUserRepository interface {
	SetDependencies(*gorm.DB, interfaces.ILogger)
	CreateUser(context.Context, *models.User) (bool, error)
	FindUser(context.Context, *models.User) (bool, error)
}
