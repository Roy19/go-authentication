package controllers

import (
	"encoding/json"
	"go-authentication/dtos"
	infraInterfaces "go-authentication/interfaces/infrastructures"
	servicesInterfaces "go-authentication/interfaces/services"
	"go-authentication/utils"
	"io"
	"net/http"
)

type UserController struct {
	userService servicesInterfaces.IUserService
	logger      infraInterfaces.ILogger
}

func (uc *UserController) SetDependencies(userService servicesInterfaces.IUserService,
	logger infraInterfaces.ILogger) {
	uc.userService = userService
	uc.logger = logger
}

func (uc *UserController) RegisterNewUser(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	var response dtos.Response
	if err != nil {
		uc.logger.Err(req.Context(), "Failed to parse request body for UserController.RegisterNewUser", err)
		response = dtos.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}
	}

	var requestDto dtos.CreateUserDtoRequest
	err = json.Unmarshal(body, &requestDto)

	if err != nil {
		uc.logger.Err(req.Context(), "Failed to parse request body for UserController.RegisterNewUser", err)
		response = dtos.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}
	}

	serviceResponse, err := uc.userService.CreateUser(req.Context(), requestDto)
	if err != nil {
		response = dtos.Response{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		}
	} else {
		response = dtos.Response{
			StatusCode: http.StatusCreated,
			Value:      serviceResponse,
		}
	}

	err = utils.WriteResponse(res, response)

	if err != nil {
		uc.logger.Err(req.Context(), "Failed to write response back", err)
	}
}

func (uc *UserController) Login(res http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	var response dtos.Response
	if err != nil {
		uc.logger.Err(req.Context(), "Failed to parse request body for UserController.Login", err)
		response = dtos.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}
	}

	var loginRequest dtos.LoginRequest
	err = json.Unmarshal(body, &loginRequest)

	if err != nil {
		uc.logger.Err(req.Context(), "Failed to parse request body for UserController.Login", err)
		response = dtos.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}
	}

	serviceResponse, err := uc.userService.LoginUser(req.Context(), loginRequest)

	if err != nil {
		response = dtos.Response{
			StatusCode: http.StatusNotFound,
			Error:      err.Error(),
		}
	} else {
		response = dtos.Response{
			StatusCode: http.StatusOK,
			Value:      serviceResponse,
		}
	}

	err = utils.WriteResponse(res, response)

	if err != nil {
		uc.logger.Err(req.Context(), "Failed to write response back", err)
	}
}
