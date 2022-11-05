package interfaces

import (
	"context"
	infraInterfaces "go-authentication/interfaces/infrastructures"
)

type IJwtService interface {
	SetDependencies(infraInterfaces.ILogger)
	CreateNewToken(ctx context.Context, userId uint) (string, error)
}
