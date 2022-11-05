package infrastructure

import (
	"context"
	interfaces "go-authentication/interfaces/infrastructures"

	"github.com/rs/zerolog"
	"gorm.io/gorm/utils"
)

type Logger struct {
	logLevel interfaces.LogLevel
	logger   zerolog.Logger
}

func NewLogger() interfaces.ILogger {
	return &Logger{
		logLevel: interfaces.Info,
		logger:   logger,
	}
}

func (l *Logger) LogMode(logLevel interfaces.LogLevel) interfaces.ILogger {
	l.logLevel = logLevel
	return l
}

func (l *Logger) Info(ctx context.Context, message string, params ...interface{}) {
	if l.logLevel <= interfaces.Info {
		l.logger.Info().Msgf("%v:%v", message, append([]interface{}{utils.FileWithLineNum()}, params))
	}
}

func (l *Logger) Err(ctx context.Context, message string, err error, params ...interface{}) {
	if l.logLevel <= interfaces.Error {
		l.logger.Err(err).Msgf("%v:%v", message, append([]interface{}{utils.FileWithLineNum()}, params))
	}
}

func (l *Logger) Debug(ctx context.Context, message string, params ...interface{}) {
	if l.logLevel <= interfaces.Debug {
		// l.logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		// 	c.Str("correlationId", ctx.Value("correlationId").(string))
		// 	return c
		// })
		l.logger.Debug().Msgf("%v:%v", message, append([]interface{}{utils.FileWithLineNum()}, params))
	}
}

func (l *Logger) Warn(ctx context.Context, message string, params ...interface{}) {
	if l.logLevel <= interfaces.Warn {
		l.logger.Warn().Msgf("%v:%v", message, append([]interface{}{utils.FileWithLineNum()}, params))
	}
}
