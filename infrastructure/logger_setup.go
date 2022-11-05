package infrastructure

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
)

var (
	logger     zerolog.Logger
	loggerOnce sync.Once
)

func InitLogger() {
	loggerOnce.Do(func() {
		logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Str("service-name", "auth-service").
			Logger()
	})
}
