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
	"inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
	service "inventory-management/backend/internal/service/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCustomerController_FindAll(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus string
		expectedBody   []*response.CustomerResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Number of customers more than 1",
			expectedStatus: "OK",
			expectedBody: []*response.CustomerResponse{
				{
					ID:        1,
					Code:      "KKSJIDNA",
					Name:      "Widdy Arfiansyah",
					CreatedAt: "2021-01-01 00:00:00",
					UpdatedAt: "2021-01-01 00:00:00",
				},
				{
					ID:        2,
					Code:      "MMAJSMWD",
					Name:      "arfn",
					CreatedAt: "2021-01-01 00:00:00",
					UpdatedAt: "2021-01-01 00:00:00",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Number of customers is 0 or null",
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

			var svc service.CustomerServiceMock
			svc.On("CountAll", ctx).Return(int64(2), nil)
			svc.On("FindAll", ctx, 0, 10).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewCustomerController(&svc, route)
			app.Get("/api/customers", ctrl.FindAll)

			req := httptest.NewRequest(http.MethodGet, "/api/customers", nil)
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

func TestCustomerController_FindByCode(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedBody   *response.CustomerResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Customer exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: "OK",
			expectedBody: &response.CustomerResponse{
				ID:        1,
				Code:      "KKSJIDNA",
				Name:      "Widdy Arfiansyah",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Customer doesnt exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: response.ErrorNotFound,
			expectedBody:   nil,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.ErrorNotFound),
		},
		{
			name:           "Service getting an error",
			request:        "KKSJIDNA",
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

			var svc service.CustomerServiceMock
			svc.On("FindByCode", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewCustomerController(&svc, route)
			app.Get("/api/customers/:code", ctrl.FindByCode)

			url := fmt.Sprintf("/api/customers/%s", tc.request)
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

func TestCustomerController_Create(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.CreateCustomerRequest
		expectedStatus string
		expectedBody   *response.CustomerResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Create customer with required fields",
			request: &request.CreateCustomerRequest{
				Name: "Widdy Arfiansyah",
			},
			expectedStatus: "created",
			expectedBody: &response.CustomerResponse{
				ID:        1,
				Code:      "KKSJIDNA",
				Name:      "Widdy Arfiansyah",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			expectedCode:  http.StatusCreated,
			expectedError: nil,
		},
		{
			name:           "[missing] Create customer with missing name field",
			request:        &request.CreateCustomerRequest{},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Name' field"),
		},
		{
			name: "[missing] Create customer with does not meet the validation requirements with the 'min' tag at name field.",
			request: &request.CreateCustomerRequest{
				Name: "Wi",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'min' for 'Name' field"),
		},
		{
			name: "[missing] Create customer with does not meet the validation requirements with the 'max' tag at name field.",
			request: &request.CreateCustomerRequest{
				Name: "Widddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Name' field"),
		},
		{
			name: "Service getting an error",
			request: &request.CreateCustomerRequest{
				Name: "Widdy Arfiansyah",
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

			var svc service.CustomerServiceMock
			svc.On("Create", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewCustomerController(&svc, route)
			app.Post("/api/customers", ctrl.Create)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			req := httptest.NewRequest(http.MethodPost, "/api/customers", bodyRequest)
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

func TestCustomerController_Update(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.UpdateCustomerRequest
		expectedStatus string
		expectedBody   *response.CustomerResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Update customer with required fields",
			request: &request.UpdateCustomerRequest{
				Code: "KKSJIDNA",
				Name: "Widdy Arfiansyah",
			},
			expectedStatus: "updated",
			expectedBody: &response.CustomerResponse{
				ID:        1,
				Code:      "KKSJIDNA",
				Name:      "Widdy Arfiansyah",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "[missing] Update customer with missing name field",
			request: &request.UpdateCustomerRequest{
				Code: "KKSJIDNA",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Name' field"),
		},
		{
			name: "[missing] Update customer with does not meet the validation requirements with the 'min' tag at name field.",
			request: &request.UpdateCustomerRequest{
				Code: "KKSJIDNA",
				Name: "Wi",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'min' for 'Name' field"),
		},
		{
			name: "[missing] Update customer with does not meet the validation requirements with the 'max' tag at name field.",
			request: &request.UpdateCustomerRequest{
				Code: "KKSJIDNA",
				Name: "Widddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddddd",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Name' field"),
		},
		{
			name: "Service getting an error",
			request: &request.UpdateCustomerRequest{
				Code: "KKSJIDNA",
				Name: "Widdy Arfiansyah",
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

			var svc service.CustomerServiceMock
			svc.On("Update", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewCustomerController(&svc, route)
			app.Patch("/api/customers/:code", ctrl.Update)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			url := fmt.Sprintf("/api/customers/%s", tc.request.Code)
			req := httptest.NewRequest(http.MethodPatch, url, bodyRequest)
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
			assert.Contains(t, responseBody.Status, tc.expectedStatus)
		})
	}
}

func TestCustomerController_Delete(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Customer exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: "deleted",
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "Customer doesnt exists with given Code when deleting data",
			request:        "KKSJIDNA",
			expectedStatus: response.ErrorNotFound,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.ErrorNotFound),
		},
		{
			name:           "Service getting an error",
			request:        "KKSJIDNA",
			expectedStatus: "getting an error",
			expectedCode:   http.StatusInternalServerError,
			expectedError:  errors.New("getting an error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New(middleware.FiberConfig())

			ctx := context.Background()

			var svc service.CustomerServiceMock
			svc.On("Delete", ctx, tc.request).Return(tc.expectedError)

			route := app.Group("/api")
			ctrl := NewCustomerController(&svc, route)
			app.Delete("/api/customers/:code", ctrl.Delete)

			url := fmt.Sprintf("/api/customers/%s", tc.request)
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
