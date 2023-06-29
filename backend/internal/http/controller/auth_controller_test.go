package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"inventory-management/backend/internal/http/middleware"
	"inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
	service "inventory-management/backend/internal/service/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthController_Login(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.LoginUserRequest
		expectedStatus string
		expectedBody   *response.UserLoginResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "User exists with given username",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
				Password: "12345678910",
			},
			expectedStatus: "OK",
			expectedBody: &response.UserLoginResponse{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc3NjUxMDQsInVzZXJuYW1lIjoiYXJmaWFuIn0.xchUDFf3xBJlfQWUzq4FCBjAHcYSAEcAiruZuuGAJBk",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "[missing] User login with missing username field",
			request: &request.LoginUserRequest{
				Password: "12345678910",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Username' field"),
		},
		{
			name: "[missing] User login with missing password field",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Password' field"),
		},
		{
			name: "User doesnt exists with given username",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
				Password: "12345678910",
			},
			expectedStatus: response.ErrorNotFound,
			expectedBody:   nil,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.ErrorNotFound),
		},
		{
			name: "Password doesnt match with given password input",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
				Password: "12345678910",
			},
			expectedStatus: response.ErrorInvalidPassword,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New(response.ErrorInvalidPassword),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New(middleware.FiberConfig())

			ctx := context.Background()

			var svc service.UserServiceMock
			svc.On("VerifyLogin", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewAuthController(&svc, route)
			app.Post("/api/login", ctrl.Login)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			req := httptest.NewRequest(http.MethodPost, "/api/login", bodyRequest)
			req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)

			res, err := app.Test(req, -1)
			assert.Nil(t, err)

			if strings.Contains(tc.name, "[missing]") {
				var responseBody response.ErrorValidationResponse
				err = json.NewDecoder(res.Body).Decode(&responseBody)
				assert.Nil(t, err)

				assert.Equal(t, responseBody.Code, tc.expectedCode)
				assert.Equal(t, responseBody.Status, tc.expectedStatus)
				assert.NotNil(t, responseBody.Error)
				assert.Equal(t, responseBody.Error[0].Value, tc.expectedError.Error())
				return
			}

			var responseBody response.ApiResponse
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			assert.Nil(t, err)

			assert.Equal(t, responseBody.Code, tc.expectedCode)
			assert.Equal(t, responseBody.Status, tc.expectedStatus)
		})
	}
}

func TestAuthController_Register(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.CreateUserRequest
		expectedStatus string
		expectedBody   *response.UserResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Create user with required fields",
			request: &request.CreateUserRequest{
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: "12345678910",
			},
			expectedStatus: "created",
			expectedBody: &response.UserResponse{
				ID:        1,
				Username:  "wdyarfn",
				Name:      "Widdy Arfiansyah",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			expectedCode:  http.StatusCreated,
			expectedError: nil,
		},
		{
			name: "Create user with missing name field",
			request: &request.CreateUserRequest{
				Username: "wdyarfn",
				Password: "12345678910",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Name' field"),
		},
		{
			name: "Create user with missing username field",
			request: &request.CreateUserRequest{
				Name:     "Widdy Arfiansyah",
				Password: "12345678910",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Username' field"),
		},
		{
			name: "Create user with missing password field",
			request: &request.CreateUserRequest{
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Password' field"),
		},
		{
			name: "Service getting an error",
			request: &request.CreateUserRequest{
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: "12345678910",
			},
			expectedStatus: "getting an error",
			expectedBody:   nil,
			expectedCode:   http.StatusInternalServerError,
			expectedError:  errors.New("getting an error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New(middleware.FiberConfig())

			ctx := context.Background()

			var svc service.UserServiceMock
			svc.On("Create", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewAuthController(&svc, route)
			app.Post("/api/register", ctrl.Register)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			req := httptest.NewRequest(http.MethodPost, "/api/register", bodyRequest)
			req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)

			res, err := app.Test(req, -1)
			assert.Nil(t, err)

			if strings.Contains(tc.name, "missing") {
				var responseBody response.ErrorValidationResponse
				err = json.NewDecoder(res.Body).Decode(&responseBody)
				assert.Nil(t, err)

				assert.Equal(t, responseBody.Code, tc.expectedCode)
				assert.Equal(t, responseBody.Status, tc.expectedStatus)
				assert.NotNil(t, responseBody.Error)
				assert.Equal(t, responseBody.Error[0].Value, tc.expectedError.Error())
				return
			}

			var responseBody response.ApiResponse
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			assert.Nil(t, err)

			assert.Equal(t, responseBody.Code, tc.expectedCode)
			assert.Equal(t, responseBody.Status, tc.expectedStatus)
		})
	}
}
