package interfaces

import (
	"context"
)

type LogLevel int

const (
	Debug LogLevel = iota + 1
	Info
	Warn
	Error
)

type ILogger interface {
	LogMode(LogLevel) ILogger
	Info(context.Context, string, ...interface{})
	Err(context.Context, string, error, ...interface{})
	Warn(context.Context, string, ...interface{})
	Debug(context.Context, string, ...interface{})
}
