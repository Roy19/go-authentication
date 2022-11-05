package services

import (
	"context"
	infraInterfaces "go-authentication/interfaces/infrastructures"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtService struct {
	logger infraInterfaces.ILogger
}

var secret = []byte("jakdjskjasdhklashu446s4d64asdsa")

func (js *JwtService) SetDependencies(logger infraInterfaces.ILogger) {
	js.logger = logger
}

func (js *JwtService) CreateNewToken(ctx context.Context, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:  strconv.FormatUint(uint64(userId), 10),
		IssuedAt: jwt.NewNumericDate(time.Now()),
	})
	tokenString, err := token.SignedString(secret)
	if err != nil {
		js.logger.Err(ctx, "Failed to sign token", err)
	}
	return tokenString, err
}
