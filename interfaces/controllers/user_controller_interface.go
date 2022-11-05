package interfaces

import (
	infraInterfaces "go-authentication/interfaces/infrastructures"
	servicesInterfaces "go-authentication/interfaces/services"
	"net/http"
)

type IUserController interface {
	SetDependencies(servicesInterfaces.IUserService, infraInterfaces.ILogger)
	RegisterNewUser(http.ResponseWriter, *http.Request)
	Login(res http.ResponseWriter, req *http.Request)
}
