package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	repository "inventory-management/backend/internal/repository/mock"
	"testing"
)

func TestCustomerService_FindAll(t *testing.T) {
	testCases := []struct {
		name                             string
		expectedCustomerRepoFindAll      []*model.Customer
		expectedCustomerRepoFindAllError error
		expectedSvc                      []*response.CustomerResponse
		expectedSvcError                 error
	}{
		{
			name: "Number of customers more than 1",
			expectedCustomerRepoFindAll: []*model.Customer{
				{
					ID:   1,
					Code: "WDWDARFSYH",
					Name: "Widdy Arfiansyah",
				},
				{
					ID:   2,
					Code: "HYSFRADWDW",
					Name: "Arfiansyah",
				},
			},
			expectedSvc: []*response.CustomerResponse{
				{
					ID:        1,
					Code:      "WDWDARFSYH",
					Name:      "Widdy Arfiansyah",
					CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
				},
				{
					ID:        2,
					Code:      "HYSFRADWDW",
					Name:      "Arfiansyah",
					CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
				},
			},
			expectedCustomerRepoFindAllError: nil,
			expectedSvcError:                 nil,
		},
		{
			name:                             "Number of customers is 0 or null",
			expectedCustomerRepoFindAll:      nil,
			expectedSvc:                      nil,
			expectedCustomerRepoFindAllError: nil,
			expectedSvcError:                 nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.CustomerRepositoryMock
			repo.On("FindAll", ctx).Return(tc.expectedCustomerRepoFindAll, tc.expectedCustomerRepoFindAllError)
			svc := NewCustomerService(&repo)
			result, err := svc.FindAll(ctx)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
			assert.Equal(t, len(tc.expectedSvc), len(result))
		})
	}
}

func TestCustomerService_FindByCode(t *testing.T) {
	testCases := []struct {
		name                                string
		request                             string
		expectedCustomerRepoFindByCode      *model.Customer
		expectedCustomerRepoFindByCodeError error
		expectedSvc                         *response.CustomerResponse
		expectedSvcError                    error
	}{
		{
			name:    "Customer exists with given code",
			request: "WDWDARFSYH",
			expectedCustomerRepoFindByCode: &model.Customer{
				ID:   1,
				Code: "WDWDARFSYH",
				Name: "Widdy Arfiansyah",
			},
			expectedSvc: &response.CustomerResponse{
				ID:        1,
				Code:      "WDWDARFSYH",
				Name:      "Widdy Arfiansyah",
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedCustomerRepoFindByCodeError: nil,
			expectedSvcError:                    nil,
		},
		{
			name:                                "Customer doesnt exists with given code",
			request:                             "WDWDARFSYH",
			expectedCustomerRepoFindByCode:      nil,
			expectedSvc:                         nil,
			expectedCustomerRepoFindByCodeError: errors.New(response.NotFound),
			expectedSvcError:                    errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.CustomerRepositoryMock
			repo.On("FindByCodeWithAssociations", ctx, tc.request).Return(tc.expectedCustomerRepoFindByCode, tc.expectedCustomerRepoFindByCodeError)
			svc := NewCustomerService(&repo)
			result, err := svc.FindByCode(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestCustomerService_Create(t *testing.T) {

	testCases := []struct {
		name                            string
		request                         *request.CreateCustomerRequest
		expectedCustomerRepoCreate      *model.Customer
		expectedCustomerRepoCreateError error
		expectedSvc                     *response.CustomerResponse
		expectedSvcError                error
	}{
		{
			name: "Create customer with required fields",
			request: &request.CreateCustomerRequest{
				Name: "Widdy Arfiansyah",
			},
			expectedCustomerRepoCreate: &model.Customer{
				ID:   1,
				Code: "WDWDARFSYH",
				Name: "Widdy Arfiansyah",
			},
			expectedSvc: &response.CustomerResponse{
				ID:        1,
				Code:      "WDWDARFSYH",
				Name:      "Widdy Arfiansyah",
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedCustomerRepoCreateError: nil,
			expectedSvcError:                nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.CustomerRepositoryMock
			repo.On("Create", ctx, mock.Anything).Return(tc.expectedCustomerRepoCreate, tc.expectedCustomerRepoCreateError)
			svc := NewCustomerService(&repo)
			result, err := svc.Create(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestCustomerService_Update(t *testing.T) {
	testCases := []struct {
		name                                string
		request                             *request.UpdateCustomerRequest
		requestCustomerRepoFindByCode       string
		expectedCustomerRepoFindByCode      *model.Customer
		expectedCustomerRepoFindByCodeError error
		expectedCustomerRepoUpdate          *model.Customer
		expectedCustomerRepoUpdateError     error
		expectedSvc                         *response.CustomerResponse
		expectedSvcError                    error
	}{
		{
			name: "Update customer with required fields",
			request: &request.UpdateCustomerRequest{
				Code: "WDWDARFSYH",
				Name: "Arfian",
			},
			requestCustomerRepoFindByCode: "WDWDARFSYH",
			expectedCustomerRepoFindByCode: &model.Customer{
				ID:   1,
				Code: "WDWDARFSYH",
				Name: "Widdy Arfiansyah",
			},
			expectedCustomerRepoUpdate: &model.Customer{
				ID:   1,
				Code: "WDWDARFSYH",
				Name: "Arfian",
			},
			expectedSvc: &response.CustomerResponse{
				ID:        1,
				Code:      "WDWDARFSYH",
				Name:      "Arfian",
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedCustomerRepoFindByCodeError: nil,
			expectedCustomerRepoUpdateError:     nil,
			expectedSvcError:                    nil,
		},
		{
			name: "Customer doesnt exists with given Code when updating data",
			request: &request.UpdateCustomerRequest{
				Code: "WDWDARFSYH",
				Name: "Arfian",
			},
			requestCustomerRepoFindByCode:       "WDWDARFSYH",
			expectedCustomerRepoFindByCode:      nil,
			expectedCustomerRepoUpdate:          nil,
			expectedSvc:                         nil,
			expectedCustomerRepoUpdateError:     errors.New("getting an error"),
			expectedCustomerRepoFindByCodeError: errors.New(response.NotFound),
			expectedSvcError:                    errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.CustomerRepositoryMock
			repo.On("FindByCode", ctx, tc.requestCustomerRepoFindByCode).Return(tc.expectedCustomerRepoFindByCode, tc.expectedCustomerRepoFindByCodeError)
			repo.On("Update", ctx, mock.Anything).Return(tc.expectedCustomerRepoUpdate, tc.expectedCustomerRepoUpdateError)
			svc := NewCustomerService(&repo)
			result, err := svc.Update(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)

			// really changed
			if err == nil {
				assert.Equal(t, tc.expectedCustomerRepoFindByCode.Name, result.Name)
			}
		})
	}
}

func TestCustomerService_Delete(t *testing.T) {
	testCases := []struct {
		name                                string
		request                             string
		expectedCustomerRepoFindByCode      *model.Customer
		expectedCustomerRepoFindByCodeError error
		expectedCustomerRepoDeleteError     error
		expectedSvcError                    error
	}{
		{
			name:    "Customer exists with given Code",
			request: "WDWDARFSYH",
			expectedCustomerRepoFindByCode: &model.Customer{
				ID:   1,
				Code: "WDWDARFSYH",
				Name: "Widdy Arfiansyah",
			},
			expectedCustomerRepoDeleteError:     nil,
			expectedSvcError:                    nil,
			expectedCustomerRepoFindByCodeError: nil,
		},
		{
			name:                                "Customer doesnt exists with given Code when deleting data",
			request:                             "WDWDARFSYH",
			expectedCustomerRepoFindByCode:      nil,
			expectedCustomerRepoDeleteError:     errors.New("getting an error"),
			expectedSvcError:                    errors.New(response.NotFound),
			expectedCustomerRepoFindByCodeError: errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.CustomerRepositoryMock
			repo.On("FindByCode", ctx, tc.request).Return(tc.expectedCustomerRepoFindByCode, tc.expectedCustomerRepoFindByCodeError)
			repo.On("Delete", ctx, tc.request).Return(tc.expectedCustomerRepoDeleteError)
			svc := NewCustomerService(&repo)
			err := svc.Delete(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}
		})
	}
}
