package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"inventory-management/backend/internal/http/middleware"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	service "inventory-management/backend/internal/service/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUserController_FindAll(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus string
		expectedBody   []*response.UserResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Number of users more than 1",
			expectedStatus: "OK",
			expectedBody: []*response.UserResponse{
				{
					ID:        1,
					Username:  "wdyarfn",
					Name:      "Widdy Arfiansyah",
					CreatedAt: "2021-01-01 00:00:00",
					UpdatedAt: "2021-01-01 00:00:00",
				},
				{
					ID:        2,
					Username:  "Arfian",
					Name:      "arfn",
					CreatedAt: "2021-01-01 00:00:00",
					UpdatedAt: "2021-01-01 00:00:00",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Number of users is 0 or null",
			expectedStatus: "OK",
			expectedBody:   nil,
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "Service getting an error",
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
			svc.On("FindAll", ctx).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewUserController(&svc, route)
			app.Get("/api/users", ctrl.FindAll)

			req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
			req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)

			res, err := app.Test(req, -1)
			assert.Nil(t, err)

			var responseBody response.ApiResponse
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			assert.Nil(t, err)

			assert.Equal(t, responseBody.Code, tc.expectedCode)
			assert.Equal(t, responseBody.Status, tc.expectedStatus)
		})
	}
}

func TestUserController_FindByID(t *testing.T) {
	testCases := []struct {
		name           string
		request        int64
		expectedStatus string
		expectedBody   *response.UserResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "User exists with given ID",
			request:        1,
			expectedStatus: "OK",
			expectedBody: &response.UserResponse{
				ID:        1,
				Username:  "wdyarfn",
				Name:      "Widdy Arfiansyah",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "User doesnt exists with given ID",
			request:        1,
			expectedStatus: response.NotFound,
			expectedBody:   nil,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.NotFound),
		},
		{
			name:           "Service getting an error",
			request:        1,
			expectedStatus: "getting an error",
			expectedBody:   nil,
			expectedCode:   http.StatusInternalServerError,
			expectedError:  errors.New("getting an error"),
		},
		{
			name:           "Wrong Params ID",
			request:        1,
			expectedStatus: "failed to convert: strconv.Atoi: parsing",
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("failed to convert: strconv.Atoi: parsing"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New(middleware.FiberConfig())

			ctx := context.Background()

			var svc service.UserServiceMock
			svc.On("FindByID", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewUserController(&svc, route)
			app.Get("/api/users/:id", ctrl.FindByID)

			url := fmt.Sprintf("/api/users/%d", tc.request)
			if tc.name == "Wrong Params ID" {
				url = "/api/users/asd"
			}
			req := httptest.NewRequest(http.MethodGet, url, nil)
			req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)

			res, err := app.Test(req, -1)
			assert.Nil(t, err)

			var responseBody response.ApiResponse
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			assert.Nil(t, err)

			assert.Equal(t, responseBody.Code, tc.expectedCode)
			assert.Contains(t, responseBody.Status, tc.expectedStatus)
		})
	}
}

func TestUserController_Create(t *testing.T) {
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
			ctrl := NewUserController(&svc, route)
			app.Post("/api/users", ctrl.Create)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			req := httptest.NewRequest(http.MethodPost, "/api/users", bodyRequest)
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

func TestUserController_Update(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.UpdateUserRequest
		expectedStatus string
		expectedBody   *response.UserResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Update user with required fields",
			request: &request.UpdateUserRequest{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Password: "12345678910",
			},
			expectedStatus: "updated",
			expectedBody: &response.UserResponse{
				ID:        1,
				Username:  "wdyarfn",
				Name:      "Widdy Arfiansyah",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "Update user with missing name field",
			request: &request.UpdateUserRequest{
				Password: "12345678910",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Name' field"),
		},
		{
			name: "Update user with missing minimum characters password field",
			request: &request.UpdateUserRequest{
				Name:     "Widdy Arfiansyah",
				Password: "12",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'min' for 'Password' field"),
		},
		{
			name: "Service getting an error",
			request: &request.UpdateUserRequest{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Password: "12345678910",
			},
			expectedStatus: "getting an error",
			expectedBody:   nil,
			expectedCode:   http.StatusInternalServerError,
			expectedError:  errors.New("getting an error"),
		},
		{
			name: "Wrong Params ID",
			request: &request.UpdateUserRequest{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Password: "12345678910",
			},
			expectedStatus: "failed to convert: strconv.Atoi: parsing",
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("failed to convert: strconv.Atoi: parsing"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New(middleware.FiberConfig())

			ctx := context.Background()

			var svc service.UserServiceMock
			svc.On("Update", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewUserController(&svc, route)
			app.Patch("/api/users/:id", ctrl.Update)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			url := fmt.Sprintf("/api/users/%d", tc.request.ID)
			if tc.name == "Wrong Params ID" {
				url = "/api/users/asd"
			}
			req := httptest.NewRequest(http.MethodPatch, url, bodyRequest)
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
			assert.Contains(t, responseBody.Status, tc.expectedStatus)
		})
	}
}

func TestUserController_Delete(t *testing.T) {
	testCases := []struct {
		name           string
		request        int64
		expectedStatus string
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "User exists with given ID",
			request:        1,
			expectedStatus: "deleted",
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "User doesnt exists with given ID when deleting data",
			request:        1,
			expectedStatus: response.NotFound,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.NotFound),
		},
		{
			name:           "Service getting an error",
			request:        1,
			expectedStatus: "getting an error",
			expectedCode:   http.StatusInternalServerError,
			expectedError:  errors.New("getting an error"),
		},
		{
			name:           "Wrong Params ID",
			request:        1,
			expectedStatus: "failed to convert: strconv.Atoi: parsing",
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("failed to convert: strconv.Atoi: parsing"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New(middleware.FiberConfig())

			ctx := context.Background()

			var svc service.UserServiceMock
			svc.On("Delete", ctx, tc.request).Return(tc.expectedError)

			route := app.Group("/api")
			ctrl := NewUserController(&svc, route)
			app.Delete("/api/users/:id", ctrl.Delete)

			url := fmt.Sprintf("/api/users/%d", tc.request)
			if tc.name == "Wrong Params ID" {
				url = "/api/users/asd"
			}
			req := httptest.NewRequest(http.MethodDelete, url, nil)
			req.Header.Set("Content-Type", fiber.MIMEApplicationJSON)

			res, err := app.Test(req, -1)
			assert.Nil(t, err)

			var responseBody response.ApiResponse
			err = json.NewDecoder(res.Body).Decode(&responseBody)
			assert.Nil(t, err)

			assert.Equal(t, responseBody.Code, tc.expectedCode)
			assert.Contains(t, responseBody.Status, tc.expectedStatus)
		})
	}
}
