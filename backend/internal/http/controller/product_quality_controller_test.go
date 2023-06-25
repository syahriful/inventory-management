package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"inventory-management/backend/internal/http/middleware"
	"inventory-management/backend/internal/http/presenter/response"
	service "inventory-management/backend/internal/service/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProductQualityController_FindAll(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus string
		expectedBody   []*response.ProductQualityResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "Number of products more than 1",
			expectedStatus: "OK",
			expectedBody: []*response.ProductQualityResponse{
				{
					ID:          1,
					ProductCode: "KKSJIDNA",
					Quality:     "Very Fresh",
					Price:       25000,
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

			var svc service.ProductQualityServiceMock
			svc.On("FindAll", ctx).Return(tc.expectedBody, tc.expectedError)

			ctrl := NewProductQualityController(&svc)
			app.Get("/api/product-qualities", ctrl.FindAll)

			req := httptest.NewRequest(http.MethodGet, "/api/product-qualities", nil)
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

func TestProductQualityController_FindAllByProductCode(t *testing.T) {
	testCases := []struct {
		name           string
		request        string
		expectedStatus string
		expectedBody   *response.ProductQualityWithOwnProductResponse
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "ProductQuality exists with given Code",
			request:        "KKSJIDNA",
			expectedStatus: "OK",
			expectedBody: &response.ProductQualityWithOwnProductResponse{
				Product: &response.ProductResponse{
					Code:                "KKSJIDNA",
					Name:                "Shrimp",
					UnitMassAcronym:     "kg",
					UnitMassDescription: "kilogram",
				},
				ProductQualities: []*response.ProductQualityResponse{
					{
						ID:          1,
						ProductCode: "KKSJIDNA",
						Quality:     "Very Fresh",
						Price:       25000,
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
			name:           "ProductQuality doesnt exists with given Code",
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

			var svc service.ProductQualityServiceMock
			svc.On("FindAllByProductCode", ctx, tc.request).Return(tc.expectedBody, tc.expectedError)

			ctrl := NewProductQualityController(&svc)
			app.Get("/api/product-qualities/:code", ctrl.FindAllByProductCode)

			url := fmt.Sprintf("/api/product-qualities/%s", tc.request)
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

func TestProductQualityController_Delete(t *testing.T) {
	testCases := []struct {
		name           string
		request        int64
		expectedStatus string
		expectedCode   int
		expectedError  error
	}{
		{
			name:           "ProductQuality exists with given ID",
			request:        1,
			expectedStatus: "deleted",
			expectedCode:   http.StatusOK,
			expectedError:  nil,
		},
		{
			name:           "ProductQuality doesnt exists with given ID when deleting data",
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

			var svc service.ProductQualityServiceMock
			svc.On("Delete", ctx, tc.request).Return(tc.expectedError)

			ctrl := NewProductQualityController(&svc)
			app.Delete("/api/product-qualities/:id", ctrl.Delete)

			url := fmt.Sprintf("/api/product-qualities/%d", tc.request)
			if tc.name == "Wrong Params ID" {
				url = "/api/product-qualities/asd"
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
