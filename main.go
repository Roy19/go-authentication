package main

import (
	"context"
	"go-authentication/controllers"
	"go-authentication/infrastructure"
	controllerInterfaces "go-authentication/interfaces/controllers"
	infraInterfaces "go-authentication/interfaces/infrastructures"
	"go-authentication/repositories"
	"go-authentication/services"
	"net/http"

	internalMiddlewares "go-authentication/middlewares"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	userController controllerInterfaces.IUserController
	router         *chi.Mux
	logger         infraInterfaces.ILogger
)

func addMiddlewares(mux *chi.Mux) {
	mux.Use(middleware.Compress(5, "gzip"))
	mux.Use(middleware.Recoverer)
	mux.Use(internalMiddlewares.HTTPanicRecovery)
}

func addPaths(mux *chi.Mux) {
	mux.Post("/register", userController.RegisterNewUser)
	mux.Post("/login", userController.Login)
}

func initInfraServices() {
	router = chi.NewRouter()
	infrastructure.InitLogger()
	infrastructure.InitDBConnection(infrastructure.POSTGRES)
	logger = infrastructure.NewLogger()
}

func initDependencies() {
	// initialize the repositories
	userRepo := &repositories.UserRepository{}
	userRepo.SetDependencies(infrastructure.DB, logger)

	// initialize the services
	jwtService := &services.JwtService{}
	jwtService.SetDependencies(logger)
	userService := &services.UserService{}
	userService.SetDependencies(userRepo, logger, jwtService)

	// initialize the controllers
	userController = &controllers.UserController{}
	userController.SetDependencies(userService, logger)
}

func main() {
	initInfraServices()
	initDependencies()
	addMiddlewares(router)
	addPaths(router)
	if err := http.ListenAndServe(":3000", router); err != nil {
		logger.Err(context.Background(), "Failed to start server", err)
	}
}
