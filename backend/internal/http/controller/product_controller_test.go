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
	request "inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
	service "inventory-management/backend/internal/service/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProductController_FindAll(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus string
		expectedBody   []*response.ProductResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Number of products more than 1",
			expectedStatus: "OK",
			expectedBody: []*response.ProductResponse{
				{
					ID:                  1,
					Code:                "KKSJIDNA",
					Name:                "Shrimp",
					UnitMassAcronym:     "kg",
					UnitMassDescription: "kilogram",
					CreatedAt:           "2021-01-01 07:00:00",
					UpdatedAt:           "2021-01-01 07:00:00",
				},
				{
					ID:                  2,
					Code:                "MMAJSMWD",
					Name:                "Shark",
					UnitMassAcronym:     "g",
					UnitMassDescription: "gram",
					CreatedAt:           "2021-01-01 07:00:00",
					UpdatedAt:           "2021-01-01 07:00:00",
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Number of products is 0 or null",
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

			var svc service.ProductServiceMock
			svc.On("CountAll", ctx).Return(int64(2), nil)
			svc.On("FindAll", ctx, 0, 10).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewProductController(&svc, route)
			app.Get("/api/products", ctrl.FindAll)

			req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
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

func TestProductController_FindByCode(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedBody   *response.ProductResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Product exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: "OK",
			expectedBody: &response.ProductResponse{
				ID:                  1,
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				CreatedAt:           "2021-01-01 07:00:00",
				UpdatedAt:           "2021-01-01 07:00:00",
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name:           "Product doesnt exists with given Code",
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

			var svc service.ProductServiceMock
			svc.On("FindByCode", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewProductController(&svc, route)
			app.Get("/api/products/:code", ctrl.FindByCode)

			url := fmt.Sprintf("/api/products/%s", tc.request)
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

func TestProductController_Create(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.CreateProductRequest
		expectedStatus string
		expectedBody   *response.ProductResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Create product with required fields",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: "created",
			expectedBody: &response.ProductResponse{
				ID:                  1,
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				CreatedAt:           "2021-01-01 07:00:00",
				UpdatedAt:           "2021-01-01 07:00:00",
				ProductQualities: []*response.ProductQualityResponse{
					{
						ID:          1,
						ProductCode: "KKSJIDNA",
						Quality:     "Very Fresh",
						Price:       20000,
						Quantity:    5.5,
						Type:        "increase",
					},
					{
						ID:          2,
						ProductCode: "KKSJIDNA",
						Quality:     "Fresh",
						Price:       15000,
						Quantity:    4.5,
						Type:        "increase",
					},
				},
			},
			expectedCode:  http.StatusCreated,
			expectedError: nil,
		},
		{
			name: "[missing] Create product with missing name field",
			request: &request.CreateProductRequest{
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Name' field"),
		},
		{
			name: "[missing] Create product with does not meet the validation requirements with the 'max' tag at name field.",
			request: &request.CreateProductRequest{
				Name:                "ProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProduk",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Name' field"),
		},
		{
			name: "[missing] Create product with missing unit mass acronym field",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'UnitMassAcronym' field"),
		},
		{
			name: "[missing] Create product with does not meet the validation requirements with the 'max' tag at unit mass acronym field.",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'UnitMassAcronym' field"),
		},
		{
			name: "[missing] Create product with missing unit mass description field",
			request: &request.CreateProductRequest{
				Name:            "Shrimp",
				UnitMassAcronym: "kg",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'UnitMassDescription' field"),
		},
		{
			name: "[missing] Create product with does not meet the validation requirements with the 'max' tag at unit mass description field.",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogramkilogramkilogramkilogramkilogramkilogramkilogramkilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'UnitMassDescription' field"),
		},
		{
			name: "[missing] Create product with missing product qualities field",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'ProductQualities' field"),
		},
		{
			name: "[missing] Create product with does not meet the validation requirements with the 'required' tag at dive quality field.",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Quality' field"),
		},
		{
			name: "[missing] Create product with does not meet the validation requirements with the 'max' tag at dive quality field.",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "FreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Quality' field"),
		},
		{
			name: "[missing] Create product with does not meet the validation requirements with the 'required' tag at dive price field.",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Fresh",
						Quantity: 5.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Price' field"),
		},
		{
			name: "[missing] Create product with does not meet the validation requirements with the 'required' tag at dive type field.",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Fresh",
						Price:    20000,
						Quantity: 5.5,
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Type' field"),
		},
		{
			name: "[missing] Create product with does not meet the validation requirements with the 'max' tag at dive type field.",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increaseincreaseincreaseincreaseincreaseincrease",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Type' field"),
		},
		{
			name: "Service getting an error",
			request: &request.CreateProductRequest{
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.CreateProductQualityRequest{
					{
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
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

			var svc service.ProductServiceMock
			svc.On("Create", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewProductController(&svc, route)
			app.Post("/api/products", ctrl.Create)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			req := httptest.NewRequest(http.MethodPost, "/api/products", bodyRequest)
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

func TestProductController_Update(t *testing.T) {
	testCases := []struct {
		name           string
		request        *request.UpdateProductRequest
		expectedStatus string
		expectedBody   *response.ProductResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name: "Update product with required fields",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						ID:       2,
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: "updated",
			expectedBody: &response.ProductResponse{
				ID:                  1,
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				CreatedAt:           "2021-01-01 07:00:00",
				UpdatedAt:           "2021-01-01 07:00:00",
				ProductQualities: []*response.ProductQualityResponse{
					{
						ID:          1,
						ProductCode: "KKSJIDNA",
						Quality:     "Very Fresh",
						Price:       20000,
						Quantity:    5.5,
						Type:        "increase",
					},
					{
						ID:          2,
						ProductCode: "KKSJIDNA",
						Quality:     "Fresh",
						Price:       15000,
						Quantity:    4.5,
						Type:        "increase",
					},
				},
			},
			expectedCode:  http.StatusOK,
			expectedError: nil,
		},
		{
			name: "[missing] Update product with missing name field",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						ID:       2,
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Name' field"),
		},
		{
			name: "[missing] Update product with does not meet the validation requirements with the 'max' tag at name field.",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "ProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProdukProduk",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						ID:       2,
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Name' field"),
		},
		{
			name: "[missing] Update product with missing unit mass acronym field",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						ID:       2,
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'UnitMassAcronym' field"),
		},
		{
			name: "[missing] Update product with does not meet the validation requirements with the 'max' tag at unit mass acronym field.",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkgkg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						ID:       2,
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'UnitMassAcronym' field"),
		},
		{
			name: "[missing] Update product with missing unit mass description field",
			request: &request.UpdateProductRequest{
				Code:            "KKSJIDNA",
				Name:            "Shrimp",
				UnitMassAcronym: "kg",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						ID:       2,
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'UnitMassDescription' field"),
		},
		{
			name: "[missing] Update product with does not meet the validation requirements with the 'max' tag at unit mass description field.",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogramkilogramkilogramkilogramkilogramkilogramkilogramkilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						ID:       2,
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'UnitMassDescription' field"),
		},
		{
			name: "[missing] Update product with missing product qualities field",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'ProductQualities' field"),
		},
		{
			name: "[missing] Update product with does not meet the validation requirements with the 'required' tag at dive quality field.",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Quality' field"),
		},
		{
			name: "[missing] Update product with does not meet the validation requirements with the 'max' tag at dive quality field.",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "FreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFreshFresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Quality' field"),
		},
		{
			name: "[missing] Update product with does not meet the validation requirements with the 'required' tag at dive price field.",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Fresh",
						Quantity: 5.5,
						Type:     "increase",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Price' field"),
		},
		{
			name: "[missing] Update product with does not meet the validation requirements with the 'required' tag at dive type field.",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Fresh",
						Price:    20000,
						Quantity: 5.5,
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'required' for 'Type' field"),
		},
		{
			name: "[missing] Update product with does not meet the validation requirements with the 'max' tag at dive type field.",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increaseincreaseincreaseincreaseincreaseincrease",
					},
				},
			},
			expectedStatus: response.ErrorValidation,
			expectedBody:   nil,
			expectedCode:   http.StatusBadRequest,
			expectedError:  errors.New("Error validation 'max' for 'Type' field"),
		},
		{
			name: "Service getting an error",
			request: &request.UpdateProductRequest{
				Code:                "KKSJIDNA",
				Name:                "Shrimp",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Very Fresh",
						Price:    20000,
						Quantity: 5.5,
						Type:     "increase",
					},
					{
						ID:       2,
						Quality:  "Fresh",
						Price:    15000,
						Quantity: 4.5,
						Type:     "increase",
					},
				},
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

			var svc service.ProductServiceMock
			svc.On("Update", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			route := app.Group("/api")
			ctrl := NewProductController(&svc, route)
			app.Patch("/api/products/:code", ctrl.Update)

			byteRequest, err := json.Marshal(tc.request)
			assert.Nil(t, err)

			bodyRequest := bytes.NewReader(byteRequest)
			url := fmt.Sprintf("/api/products/%s", tc.request.Code)
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

func TestProductController_Delete(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Product exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: "deleted",
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "Product doesnt exists with given Code when deleting data",
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

			var svc service.ProductServiceMock
			svc.On("Delete", ctx, tc.request).Return(tc.expectedError)

			route := app.Group("/api")
			ctrl := NewProductController(&svc, route)
			app.Delete("/api/products/:code", ctrl.Delete)

			url := fmt.Sprintf("/api/products/%s", tc.request)
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
