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
	"inventory-management/backend/internal/http/response"
	service "inventory-management/backend/internal/service/mock"
	"inventory-management/backend/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTransactionController_FindAll(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus string
		expectedBody   []*response.TransactionResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Number of transactions more than 1",
			expectedStatus: "OK",
			expectedBody: []*response.TransactionResponse{
				{
					ID:               1,
					Code:             "WDWDARFSYH",
					ProductQualityID: 1,
					SupplierCode:     util.ToPointerString("SUP001"),
					Description:      util.ToPointerString("Pembelian barang dari supplier"),
					Quantity:         23,
					Type:             "IN",
					UnitMassAcronym:  "kg",
					CreatedAt:        "2021-01-01 07:00:00",
					UpdatedAt:        "2021-01-01 07:00:00",
				},
				{
					ID:               2,
					Code:             "ARFNSYH",
					ProductQualityID: 1,
					CustomerCode:     util.ToPointerString("CUS001"),
					Description:      util.ToPointerString("Pembelian barang dari supplier"),
					Quantity:         3,
					Type:             "OUT",
					UnitMassAcronym:  "kg",
					CreatedAt:        "2021-01-01 07:00:00",
					UpdatedAt:        "2021-01-01 07:00:00",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Number of transactions is 0 or null",
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

			var svc service.TransactionServiceMock
			svc.On("CountAll", ctx).Return(int64(2), nil)
			svc.On("FindAll", ctx, 0, 10).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewTransactionController(&svc, route)
			app.Get("/api/transactions", ctrl.FindAll)

			req := httptest.NewRequest(http.MethodGet, "/api/transactions", nil)
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

func TestTransactionController_FindAllBySupplierByCode(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedBody   []*response.TransactionResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Number of transactions more than 1",
			request:        "WDWDARFSYH",
			expectedStatus: "OK",
			expectedBody: []*response.TransactionResponse{
				{
					ID:               1,
					Code:             "WDWDARFSYH",
					ProductQualityID: 1,
					SupplierCode:     util.ToPointerString("SUP001"),
					Description:      util.ToPointerString("Pembelian barang dari supplier"),
					Quantity:         23,
					Type:             "IN",
					UnitMassAcronym:  "kg",
					CreatedAt:        "2021-01-01 07:00:00",
					UpdatedAt:        "2021-01-01 07:00:00",
				},
				{
					ID:               2,
					Code:             "ARFNSYH",
					ProductQualityID: 1,
					SupplierCode:     util.ToPointerString("SUP001"),
					Description:      util.ToPointerString("Dapet barang dari supplier"),
					Quantity:         3,
					Type:             "OUT",
					UnitMassAcronym:  "kg",
					CreatedAt:        "2021-01-01 07:00:00",
					UpdatedAt:        "2021-01-01 07:00:00",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Number of transactions is 0 or null",
			request:        "WDWDARFSYH",
			expectedStatus: "OK",
			expectedBody:   nil,
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "Service getting an error",
			request:        "WDWDARFSYH",
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

			var svc service.TransactionServiceMock
			svc.On("FindAllBySupplierCode", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewTransactionController(&svc, route)
			app.Get("/api/transactions/:code/supplier", ctrl.FindAllBySupplierCode)

			url := fmt.Sprintf("/api/transactions/%s/supplier", tc.request)
			req := httptest.NewRequest(http.MethodGet, url, nil)
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

func TestTransactionController_FindAllByCustomerByCode(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedBody   []*response.TransactionResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Number of transactions more than 1",
			request:        "WDWDARFSYH",
			expectedStatus: "OK",
			expectedBody: []*response.TransactionResponse{
				{
					ID:               1,
					Code:             "WDWDARFSYH",
					ProductQualityID: 1,
					CustomerCode:     util.ToPointerString("CUS001"),
					Description:      util.ToPointerString("Pembelian barang dari customer"),
					Quantity:         23,
					Type:             "IN",
					UnitMassAcronym:  "kg",
					CreatedAt:        "2021-01-01 07:00:00",
					UpdatedAt:        "2021-01-01 07:00:00",
				},
				{
					ID:               2,
					Code:             "ARFNSYH",
					ProductQualityID: 1,
					CustomerCode:     util.ToPointerString("CUS001"),
					Description:      util.ToPointerString("Dapet barang dari customer"),
					Quantity:         3,
					Type:             "OUT",
					UnitMassAcronym:  "kg",
					CreatedAt:        "2021-01-01 07:00:00",
					UpdatedAt:        "2021-01-01 07:00:00",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Number of transactions is 0 or null",
			request:        "WDWDARFSYH",
			expectedStatus: "OK",
			expectedBody:   nil,
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "Service getting an error",
			request:        "WDWDARFSYH",
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

			var svc service.TransactionServiceMock
			svc.On("FindAllByCustomerCode", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewTransactionController(&svc, route)
			app.Get("/api/transactions/:code/customer", ctrl.FindAllByCustomerCode)

			url := fmt.Sprintf("/api/transactions/%s/customer", tc.request)
			req := httptest.NewRequest(http.MethodGet, url, nil)
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

func TestTransactionController_FindByCode(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedBody   *response.TransactionResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Transaction customer is exists with given code with type IN",
			request:        "WDWDARFSYH",
			expectedStatus: "OK",
			expectedBody: &response.TransactionResponse{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				ProductQuality: &response.ProductQualityResponse{
					ID:          1,
					ProductCode: "SHRMP",
					Quality:     "Fresh",
					Quantity:    23,
					Price:       100000,
					Type:        "increase",
					Product: &response.ProductResponse{
						ID:                  1,
						Code:                "SHRMP",
						Name:                "Shrimp",
						UnitMassAcronym:     "kg",
						UnitMassDescription: "kilogram",
						CreatedAt:           "0001-01-01 07:00:00 +0700 +07",
						UpdatedAt:           "0001-01-01 07:00:00 +0700 +07",
					},
				},
				SupplierCode: util.ToPointerString("SUP001"),
				Supplier: &response.SupplierResponse{
					ID:        1,
					Code:      "SUP001",
					Name:      "PT. Sumber Makmur",
					Address:   "Jl. Raya Bogor",
					Phone:     "08123456789",
					CreatedAt: "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
				},
				Description:     util.ToPointerString("Menambahkan stok barang dari supplier"),
				Quantity:        23,
				Type:            "IN",
				UnitMassAcronym: "kg",
				CreatedAt:       "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt:       "0001-01-01 07:00:00 +0700 +07",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Transaction customer is exists with given code with type OUT",
			request:        "WDWDARFSYH",
			expectedStatus: "OK",
			expectedBody: &response.TransactionResponse{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				ProductQuality: &response.ProductQualityResponse{
					ID:          1,
					ProductCode: "SHRMP",
					Quality:     "Fresh",
					Quantity:    23,
					Price:       100000,
					Type:        "increase",
					Product: &response.ProductResponse{
						ID:                  1,
						Code:                "SHRMP",
						Name:                "Shrimp",
						UnitMassAcronym:     "kg",
						UnitMassDescription: "kilogram",
						CreatedAt:           "0001-01-01 07:00:00 +0700 +07",
						UpdatedAt:           "0001-01-01 07:00:00 +0700 +07",
					},
				},
				CustomerCode: util.ToPointerString("CUS001"),
				Customer: &response.CustomerResponse{
					ID:        1,
					Code:      "CUS001",
					Name:      "PT. Sumber Makmur",
					CreatedAt: "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
				},
				Description:     util.ToPointerString("Pembelian barang dari customer"),
				Quantity:        23,
				Type:            "OUT",
				UnitMassAcronym: "kg",
				CreatedAt:       "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt:       "0001-01-01 07:00:00 +0700 +07",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Transaction exists with given code with type TRANSFER",
			request:        "WDWDARFSYH",
			expectedStatus: "OK",
			expectedBody: &response.TransactionResponse{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				ProductQuality: &response.ProductQualityResponse{
					ID:          1,
					ProductCode: "SHRMP",
					Quality:     "Fresh",
					Quantity:    23,
					Price:       100000,
					Type:        "increase",
					Product: &response.ProductResponse{
						ID:                  1,
						Code:                "SHRMP",
						Name:                "Shrimp",
						UnitMassAcronym:     "kg",
						UnitMassDescription: "kilogram",
						CreatedAt:           "0001-01-01 07:00:00 +0700 +07",
						UpdatedAt:           "0001-01-01 07:00:00 +0700 +07",
					},
				},
				ProductQualityIDTransferred: util.ToPointerInt64(2),
				ProductQualityTransferred: &response.ProductQualityResponse{
					ID:          2,
					ProductCode: "SHRMP",
					Quality:     "Expired",
					Quantity:    2,
					Price:       1000,
					Type:        "decrease",
					Product: &response.ProductResponse{
						ID:                  1,
						Code:                "SHRMP",
						Name:                "Shrimp",
						UnitMassAcronym:     "kg",
						UnitMassDescription: "kilogram",
						CreatedAt:           "0001-01-01 07:00:00 +0700 +07",
						UpdatedAt:           "0001-01-01 07:00:00 +0700 +07",
					},
				},
				Description:     util.ToPointerString("Barang sudah expired dan dibuang"),
				Quantity:        20,
				Type:            "TRANSFER",
				UnitMassAcronym: "kg",
				CreatedAt:       "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt:       "0001-01-01 07:00:00 +0700 +07",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Transaction doesnt exists with given Code",
			request:        "WDWDARFSYH",
			expectedStatus: response.ErrorNotFound,
			expectedBody:   nil,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.ErrorNotFound),
		},
		{
			name:           "Service getting an error",
			request:        "WDWDARFSYH",
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

			var svc service.TransactionServiceMock
			svc.On("FindByCode", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewTransactionController(&svc, route)
			app.Get("/api/transactions/:code", ctrl.FindByCode)

			url := fmt.Sprintf("/api/transactions/%s", tc.request)
			req := httptest.NewRequest(http.MethodGet, url, nil)
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

func TestTransactionController_Create(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.CreateTransactionRequest
		expectedStatus string
		expectedBody   *response.TransactionResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Create transaction customer with required fields",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("CUS001"),
				Description:      util.ToPointerString("Pembelian barang dari customer"),
				Quantity:         23,
				Type:             "OUT",
				UnitMassAcronym:  "kg",
			},
			expectedStatus: "created",
			expectedBody: &response.TransactionResponse{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("CUS001"),
				Description:      util.ToPointerString("Pembelian barang dari customer"),
				Quantity:         23,
				Type:             "OUT",
				UnitMassAcronym:  "kg",
				CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
			},
			expectedCode:  http.StatusCreated,
			expectedError: nil,
		},
		{
			name: "Create transaction supplier with required fields",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
			},
			expectedStatus: "created",
			expectedBody: &response.TransactionResponse{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
				CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
			},
			expectedCode:  http.StatusCreated,
			expectedError: nil,
		},
		{
			name: "Transaction doesnt exists with given Product Quality ID when creating data",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
			},
			expectedStatus: response.ErrorNotFound,
			expectedBody:   nil,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.ErrorNotFound),
		},
		{
			name: "[missing] Create transaction with missing product quality id field",
			request: &request.CreateTransactionRequest{
				CustomerCode:    util.ToPointerString("SUP001"),
				Description:     util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:        23,
				Type:            "IN",
				UnitMassAcronym: "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'ProductQualityID' field"),
		},
		{
			name: "[missing] Create transaction with does not meet the validation requirements with the 'max' tag at customer code field.",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"),
				Description:      util.ToPointerString("Pembelian barang dari customer"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'CustomerCode' field"),
		},
		{
			name: "[missing] Create transaction with does not meet the validation requirements with the 'max' tag at supplier code field.",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				SupplierCode:     util.ToPointerString("SUP001sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'SupplierCode' field"),
		},
		{
			name: "[missing] Create transaction with does not meet the validation requirements with the 'max' tag at description field.",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				SupplierCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Description' field"),
		},
		{
			name: "[missing] Create transaction with missing quantity field",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Type:             "IN",
				UnitMassAcronym:  "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Quantity' field"),
		},
		{
			name: "[missing] Create transaction with missing type field",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				UnitMassAcronym:  "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Type' field"),
		},
		{
			name: "[missing] Create transaction with does not meet the validation requirements with the 'oneof' tag at type field.",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "FREE",
				UnitMassAcronym:  "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'oneof' for 'Type' field"),
		},
		{
			name: "[missing] Create transaction with missing unit mass acronym field",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'UnitMassAcronym' field"),
		},
		{
			name: "[missing] Create transaction with does not meet the validation requirements with the 'oneof' tag at unit mass acronym field.",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "km",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'oneof' for 'UnitMassAcronym' field"),
		},
		{
			name: "Failed create transaction",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
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

			var svc service.TransactionServiceMock
			svc.On("Create", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewTransactionController(&svc, route)
			app.Post("/api/transactions", ctrl.Create)

			body, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/transactions", bytes.NewBuffer(body))
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

func TestTransactionController_Update(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.UpdateTransactionRequest
		expectedStatus string
		expectedBody   *response.TransactionResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Update transaction customer with required fields",
			request: &request.UpdateTransactionRequest{
				Code:            "WDWDARFSYH",
				CustomerCode:    util.ToPointerString("CUS001"),
				Description:     util.ToPointerString("Ubah pembelian barang dari customer"),
				Quantity:        23,
				UnitMassAcronym: "kg",
			},
			expectedStatus: "updated",
			expectedBody: &response.TransactionResponse{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("CUS001"),
				Description:      util.ToPointerString("Ubah pembelian barang dari customer"),
				Quantity:         23,
				Type:             "OUT",
				UnitMassAcronym:  "kg",
				CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "Update transaction supplier with required fields",
			request: &request.UpdateTransactionRequest{
				Code:            "WDWDARFSYH",
				SupplierCode:    util.ToPointerString("SUP001"),
				Description:     util.ToPointerString("Ubah pembelian barang dari supplier"),
				Quantity:        23,
				UnitMassAcronym: "kg",
			},
			expectedStatus: "updated",
			expectedBody: &response.TransactionResponse{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Ubah pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
				CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "[missing] Update transaction with does not meet the validation requirements with the 'max' tag at customer code field.",
			request: &request.UpdateTransactionRequest{
				Code:            "WDWDARFSYH",
				CustomerCode:    util.ToPointerString("SUP001sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"),
				Description:     util.ToPointerString("Pembelian barang dari customer"),
				Quantity:        23,
				UnitMassAcronym: "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'CustomerCode' field"),
		},
		{
			name: "[missing] Update transaction with does not meet the validation requirements with the 'max' tag at supplier code field.",
			request: &request.UpdateTransactionRequest{
				Code:            "WDWDARFSYH",
				SupplierCode:    util.ToPointerString("SUP001sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"),
				Description:     util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:        23,
				UnitMassAcronym: "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'SupplierCode' field"),
		},
		{
			name: "[missing] Update transaction with does not meet the validation requirements with the 'max' tag at description field.",
			request: &request.UpdateTransactionRequest{
				Code:            "WDWDARFSYH",
				SupplierCode:    util.ToPointerString("SUP001"),
				Description:     util.ToPointerString("Pembelian barang dari sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"),
				Quantity:        23,
				UnitMassAcronym: "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Description' field"),
		},
		{
			name: "[missing] Update transaction with missing quantity field",
			request: &request.UpdateTransactionRequest{
				Code:            "WDWDARFSYH",
				CustomerCode:    util.ToPointerString("SUP001"),
				Description:     util.ToPointerString("Pembelian barang dari supplier"),
				UnitMassAcronym: "kg",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Quantity' field"),
		},
		{
			name: "[missing] Update transaction with missing unit mass acronym field",
			request: &request.UpdateTransactionRequest{
				Code:         "WDWDARFSYH",
				CustomerCode: util.ToPointerString("SUP001"),
				Description:  util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:     23,
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'UnitMassAcronym' field"),
		},
		{
			name: "[missing] Update transaction with does not meet the validation requirements with the 'oneof' tag at unit mass acronym field.",
			request: &request.UpdateTransactionRequest{
				Code:            "WDWDARFSYH",
				CustomerCode:    util.ToPointerString("SUP001"),
				Description:     util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:        23,
				UnitMassAcronym: "km",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'oneof' for 'UnitMassAcronym' field"),
		},
		{
			name: "Transaction doesnt exists with given Code when updating data",
			request: &request.UpdateTransactionRequest{
				Code:            "WDWDARFSYH",
				CustomerCode:    util.ToPointerString("CUS001"),
				Description:     util.ToPointerString("Ubah pembelian barang dari customer"),
				Quantity:        23,
				UnitMassAcronym: "kg",
			},
			expectedStatus: response.ErrorNotFound,
			expectedBody:   nil,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New(middleware.FiberConfig())

			ctx := context.Background()

			var svc service.TransactionServiceMock
			svc.On("Update", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewTransactionController(&svc, route)
			app.Patch("/api/transactions/:code", ctrl.Update)

			body, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			url := fmt.Sprintf("/api/transactions/%s", tc.request.Code)
			req := httptest.NewRequest(http.MethodPatch, url, bytes.NewBuffer(body))
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

func TestTransactionController_TransferStock(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.TransferStockTransactionRequest
		expectedStatus string
		expectedBody   *response.TransactionResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Transfer stock with required fields",
			request: &request.TransferStockTransactionRequest{
				ProductQualityID:            1,
				ProductQualityIDTransferred: 2,
				Description:                 util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:                    23,
			},
			expectedStatus: "transferred",
			expectedBody: &response.TransactionResponse{
				ID:                          1,
				Code:                        "WDWDARFSYH",
				ProductQualityID:            1,
				ProductQualityIDTransferred: util.ToPointerInt64(2),
				Description:                 util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:                    23,
				Type:                        "TRANSFER",
				UnitMassAcronym:             "kg",
				CreatedAt:                   "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt:                   "0001-01-01 07:00:00 +0700 +07",
			},
			expectedCode:  http.StatusCreated,
			expectedError: nil,
		},
		{
			name: "Transaction doesnt exists with given ProductQualityID when transferring stock",
			request: &request.TransferStockTransactionRequest{
				ProductQualityID:            1,
				ProductQualityIDTransferred: 2,
				Description:                 util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:                    23,
			},
			expectedStatus: response.ErrorNotFound,
			expectedBody:   nil,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.ErrorNotFound),
		},
		{
			name: "[missing] Transfer stock transaction with missing product quality ID",
			request: &request.TransferStockTransactionRequest{
				ProductQualityIDTransferred: 2,
				Description:                 util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:                    23,
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'ProductQualityID' field"),
		},
		{
			name: "[missing] Transfer stock transaction with missing product quality ID transferred",
			request: &request.TransferStockTransactionRequest{
				ProductQualityID: 1,
				Description:      util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:         23,
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'ProductQualityIDTransferred' field"),
		},
		{
			name: "[missing] Transfer stock transaction with does not meet the validation requirements with the 'max' tag at description field.",
			request: &request.TransferStockTransactionRequest{
				ProductQualityID:            1,
				ProductQualityIDTransferred: 2,
				Description:                 util.ToPointerString("Pembelian barang dari sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss"),
				Quantity:                    23,
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Description' field"),
		},
		{
			name: "[missing] Transfer stock transaction with missing quantity field",
			request: &request.TransferStockTransactionRequest{
				ProductQualityID:            1,
				ProductQualityIDTransferred: 2,
				Description:                 util.ToPointerString("Pembelian barang dari supplier"),
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Quantity' field"),
		},
		{
			name: "Transaction doesnt exists with given ProductQualityIDTransferred when transferring stock",
			request: &request.TransferStockTransactionRequest{
				ProductQualityID:            1,
				ProductQualityIDTransferred: 2,
				Description:                 util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:                    23,
			},
			expectedStatus: response.ErrorNotFound,
			expectedBody:   nil,
			expectedCode:   http.StatusNotFound,
			expectedError:  errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			app := fiber.New(middleware.FiberConfig())

			ctx := context.Background()

			var svc service.TransactionServiceMock
			svc.On("TransferStock", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewTransactionController(&svc, route)
			app.Patch("/api/transactions/transfer", ctrl.TransferStock)

			body, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			req := httptest.NewRequest(http.MethodPost, "/api/transactions/transfer", bytes.NewBuffer(body))
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

func TestTransactionController_Delete(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Transaction exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: "deleted",
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "Transaction doesnt exists with given Code when deleting data",
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

			var svc service.TransactionServiceMock
			svc.On("Delete", ctx, tc.request).Return(tc.expectedError)

			route := app.Group("/api")
			ctrl := NewTransactionController(&svc, route)
			app.Delete("/api/transactions/:code", ctrl.Delete)

			url := fmt.Sprintf("/api/transactions/%s", tc.request)
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
