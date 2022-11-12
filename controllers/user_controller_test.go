package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"go-authentication/dtos"
	infraInterfaces "go-authentication/interfaces/infrastructures"
	repositoryInterfaces "go-authentication/interfaces/repository"
	serviceInterfaces "go-authentication/interfaces/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakeUserService struct {
	returnResp any
	returnErr  any
}

func (*FakeUserService) SetDependencies(repositoryInterfaces.IUserRepository,
	infraInterfaces.ILogger,
	serviceInterfaces.IJwtService) {
}

func (fs *FakeUserService) CreateUser(ctx context.Context,
	req dtos.CreateUserDtoRequest) (dtos.CreateUserDtoResponse, error) {
	res, _ := fs.returnResp.(dtos.CreateUserDtoResponse)
	err, _ := fs.returnErr.(error)
	return res, err
}

func (fs *FakeUserService) LoginUser(ctx context.Context,
	loginRequest dtos.LoginRequest) (dtos.LoginResponse, error) {
	res, _ := fs.returnResp.(dtos.LoginResponse)
	err, _ := fs.returnErr.(error)
	return res, err
}

type FakeLogger struct{}

func (fk *FakeLogger) LogMode(infraInterfaces.LogLevel) infraInterfaces.ILogger {
	return fk
}
func (*FakeLogger) Info(context.Context, string, ...interface{})       {}
func (*FakeLogger) Err(context.Context, string, error, ...interface{}) {}
func (*FakeLogger) Warn(context.Context, string, ...interface{})       {}
func (*FakeLogger) Debug(context.Context, string, ...interface{})      {}

func TestSetDependencies(t *testing.T) {
	uc := &UserController{}
	uc.SetDependencies(&FakeUserService{}, &FakeLogger{})

	if uc.logger == nil {
		t.Errorf("Failed to initialize logger in UserController")
	}

	if uc.userService == nil {
		t.Errorf("Failed to initialize UserService inside UserController")
	}
}

type RegisterUserTest struct {
	input         dtos.CreateUserDtoRequest
	serviceOutput dtos.CreateUserDtoResponse
	serviceError  error
	returnStatus  int
	returnValue   dtos.Response
}

func TestRegisterNewUser(t *testing.T) {
	testTable := []RegisterUserTest{
		{
			input: dtos.CreateUserDtoRequest{
				Email:    "test@test.com",
				UserName: "test",
				Password: "abcd1234",
			},
			serviceOutput: dtos.CreateUserDtoResponse{
				ID:       1,
				Email:    "test@test.com",
				UserName: "test",
			},
			serviceError: nil,
			returnStatus: http.StatusCreated,
			returnValue: dtos.Response{
				StatusCode: http.StatusCreated,
				Value: dtos.CreateUserDtoResponse{
					ID:       1,
					Email:    "test@test.com",
					UserName: "test",
				},
			},
		},
	}

	for _, test := range testTable {
		uc := &UserController{}
		uc.SetDependencies(&FakeUserService{
			returnResp: test.serviceOutput,
			returnErr:  test.serviceError,
		}, &FakeLogger{})
		data, _ := json.Marshal(test.input)
		toSend := bytes.NewBuffer(data)
		req := httptest.NewRequest(http.MethodPost, "/register", toSend)
		rr := httptest.NewRecorder()
		uc.RegisterNewUser(rr, req)

		expected, got := test.returnStatus, rr.Code
		if got != expected {
			t.Errorf("Failed to get a valid return status. Expected: %v, Got: %v", expected, got)
		}

		var gotBody dtos.Response
		expectedBody := test.returnValue
		if err := json.NewDecoder(rr.Body).Decode(&gotBody); err != nil {
			t.Errorf("Failed to decode response body")
		}

		if !checkResponses(expectedBody, gotBody) {
			t.Errorf("Received the wrong response. Expected: %v, Got: %v", expectedBody, gotBody)
		}
	}
}

func checkResponses(val1, val2 dtos.Response) bool {
	if val1.StatusCode != val2.StatusCode {
		return false
	}
	// need better comparison here
	// if val1.Value != val2.Value {
	// 	return false
	// }
	// if val1.Error != val2.Error {
	// 	return false
	// }
	return true
}
