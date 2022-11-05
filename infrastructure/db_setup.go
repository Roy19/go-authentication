package infrastructure

import (
	"go-authentication/models"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	DBErr  error
	dbOnce sync.Once
)

func InitDBConnection(dbType DBType) {
	dbOnce.Do(func() {
		dsn := "host=localhost user=postgres password=postgres dbname=go_authentication port=5432 sslmode=disable TimeZone=Asia/Kolkata"
		if dbType == POSTGRES {
			DB, DBErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
			logger.Info().Msg("Initialized connection to DB")
		}
		// auto-migrate all models to DB
		DB.AutoMigrate(models.User{})
		logger.Info().Msg("Migrated all models to DB")
	})
}
