package controller

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"inventory-management/backend/internal/http/middleware"
	"inventory-management/backend/internal/http/presenter/response"
	service "inventory-management/backend/internal/service/mock"
	"net/http"
	"net/http/httptest"
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
					Username:  "admin",
					Name:      "Administrator",
					CreatedAt: "2021-01-01 00:00:00",
					UpdatedAt: "2021-01-01 00:00:00",
				},
				{
					ID:        2,
					Username:  "member",
					Name:      "Member",
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
			name:           "Service error",
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

			ctrl := NewUserController(&svc)
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
