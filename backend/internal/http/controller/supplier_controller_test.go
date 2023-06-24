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

func TestSupplierController_FindAll(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus string
		expectedBody   []*response.SupplierResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Number of suppliers more than 1",
			expectedStatus: "OK",
			expectedBody: []*response.SupplierResponse{
				{
					ID:        1,
					Code:      "KKSJIDNA",
					Name:      "Widdy Arfiansyah",
					Address:   "Bhayangkara",
					Phone:     "089911182399",
					CreatedAt: "2021-01-01 00:00:00",
					UpdatedAt: "2021-01-01 00:00:00",
				},
				{
					ID:        2,
					Code:      "MMAJSMWD",
					Name:      "arfn",
					Address:   "Cisaat",
					Phone:     "089911183725",
					CreatedAt: "2021-01-01 00:00:00",
					UpdatedAt: "2021-01-01 00:00:00",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Number of suppliers is 0 or null",
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

			var svc service.SupplierServiceMock
			svc.On("FindAll", ctx).Return(tc.expectedBody, tc.expectedError)

			ctrl := NewSupplierController(&svc)
			app.Get("/api/suppliers", ctrl.FindAll)

			req := httptest.NewRequest(http.MethodGet, "/api/suppliers", nil)
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

func TestSupplierController_FindByCode(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedBody   *response.SupplierResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Supplier exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: "OK",
			expectedBody: &response.SupplierResponse{
				ID:        1,
				Code:      "KKSJIDNA",
				Name:      "Widdy Arfiansyah",
				Address:   "Bhayangkara",
				Phone:     "089911182399",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Supplier doesnt exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: response.NotFound,
			expectedBody:   nil,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.NotFound),
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

			var svc service.SupplierServiceMock
			svc.On("FindByCode", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			ctrl := NewSupplierController(&svc)
			app.Get("/api/suppliers/:code", ctrl.FindByCode)

			url := fmt.Sprintf("/api/suppliers/%s", tc.request)
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

func TestSupplierController_Create(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.CreateSupplierRequest
		expectedStatus string
		expectedBody   *response.SupplierResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Create supplier with required fields",
			request: &request.CreateSupplierRequest{
				Name:    "Widdy Arfiansyah",
				Address: "Bhayangkara",
				Phone:   "089911182399",
			},
			expectedStatus: "created",
			expectedBody: &response.SupplierResponse{
				ID:        1,
				Code:      "KKSJIDNA",
				Name:      "Widdy Arfiansyah",
				Address:   "Bhayangkara",
				Phone:     "089911182399",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			expectedCode:  http.StatusCreated,
			expectedError: nil,
		},
		{
			name: "Create supplier with missing name field",
			request: &request.CreateSupplierRequest{
				Address: "Bhayangkara",
				Phone:   "089911182399",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Name' field"),
		},
		{
			name: "Create supplier with missing address field",
			request: &request.CreateSupplierRequest{
				Name:  "Widdy Arfiansyah",
				Phone: "089911182399",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Address' field"),
		},
		{
			name: "Create supplier with missing phone field",
			request: &request.CreateSupplierRequest{
				Name:    "Widdy Arfiansyah",
				Address: "Bhayangkara",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Phone' field"),
		},
		{
			name: "Service getting an error",
			request: &request.CreateSupplierRequest{
				Name:    "Widdy Arfiansyah",
				Address: "Bhayangkara",
				Phone:   "089911182399",
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

			var svc service.SupplierServiceMock
			svc.On("Create", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			ctrl := NewSupplierController(&svc)
			app.Post("/api/suppliers", ctrl.Create)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			req := httptest.NewRequest(http.MethodPost, "/api/suppliers", bodyRequest)
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

func TestSupplierController_Update(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.UpdateSupplierRequest
		expectedStatus string
		expectedBody   *response.SupplierResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Update supplier with required fields",
			request: &request.UpdateSupplierRequest{
				Code:    "KKSJIDNA",
				Name:    "Widdy Arfiansyah",
				Address: "Bhayangkara",
				Phone:   "089911182399",
			},
			expectedStatus: "updated",
			expectedBody: &response.SupplierResponse{
				ID:        1,
				Code:      "KKSJIDNA",
				Name:      "Widdy Arfiansyah",
				Address:   "Bhayangkara",
				Phone:     "089911182399",
				CreatedAt: "2021-01-01 00:00:00",
				UpdatedAt: "2021-01-01 00:00:00",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "Update supplier with missing name field",
			request: &request.UpdateSupplierRequest{
				Code:    "KKSJIDNA",
				Address: "Bhayangkara",
				Phone:   "089911182399",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Name' field"),
		},
		{
			name: "Update supplier with missing address field",
			request: &request.UpdateSupplierRequest{
				Code:  "KKSJIDNA",
				Name:  "Widdy Arfiansyah",
				Phone: "089911182399",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Address' field"),
		},
		{
			name: "Update supplier with missing phone field",
			request: &request.UpdateSupplierRequest{
				Code:    "KKSJIDNA",
				Name:    "Widdy Arfiansyah",
				Address: "Bhayangkara",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Phone' field"),
		},
		{
			name: "Service getting an error",
			request: &request.UpdateSupplierRequest{
				Code:    "KKSJIDNA",
				Name:    "Widdy Arfiansyah",
				Address: "Bhayangkara",
				Phone:   "089911182399",
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

			var svc service.SupplierServiceMock
			svc.On("Update", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			ctrl := NewSupplierController(&svc)
			app.Patch("/api/suppliers/:code", ctrl.Update)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			url := fmt.Sprintf("/api/suppliers/%s", tc.request.Code)
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

func TestSupplierController_Delete(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Supplier exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: "deleted",
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "Supplier doesnt exists with given Code when deleting data",
			request:        "KKSJIDNA",
			expectedStatus: response.NotFound,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.NotFound),
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

			var svc service.SupplierServiceMock
			svc.On("Delete", ctx, tc.request).Return(tc.expectedError)

			ctrl := NewSupplierController(&svc)
			app.Delete("/api/suppliers/:code", ctrl.Delete)

			url := fmt.Sprintf("/api/suppliers/%s", tc.request)
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
