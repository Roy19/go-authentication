package repositories

import (
	"context"
	interfaces "go-authentication/interfaces/infrastructures"
	"go-authentication/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	logger interfaces.ILogger
}

func (ur *UserRepository) SetDependencies(dbConnection *gorm.DB, logger interfaces.ILogger) {
	ur.db = dbConnection
	ur.logger = logger
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *models.User) (bool, error) {
	result := ur.db.Create(user)
	if result.Error != nil {
		ur.logger.Err(ctx, "Failed to create an entry in Users table", result.Error)
		return false, result.Error
	} else {
		return true, nil
	}
}

func (ur *UserRepository) FindUser(ctx context.Context, user *models.User) (bool, error) {
	result := ur.db.Where(user).First(user)
	if result.Error != nil {
		ur.logger.Err(ctx, "Failed to find user in DB", result.Error)
		return false, result.Error
	} else {
		return true, nil
	}
}
