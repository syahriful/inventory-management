package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/model"
	repository "inventory-management/backend/internal/repository/mock"
	"testing"
)

func TestSupplierService_FindAll(t *testing.T) {
	testCases := []struct {
		name                             string
		expectedSupplierRepoFindAll      []*model.Supplier
		expectedSupplierRepoFindAllError error
		expectedSvc                      []*response.SupplierResponse
		expectedSvcError                 error
	}{
		{
			name: "Number of suppliers more than 1",
			expectedSupplierRepoFindAll: []*model.Supplier{
				{
					ID:      1,
					Code:    "WDWDARFSYH",
					Name:    "Widdy Arfiansyah",
					Address: "Sukabumi",
					Phone:   "082291832488",
				},
				{
					ID:      2,
					Code:    "HYSFRADWDW",
					Name:    "Arfiansyah",
					Address: "Cisaat",
					Phone:   "089922234214",
				},
			},
			expectedSvc: []*response.SupplierResponse{
				{
					ID:        1,
					Code:      "WDWDARFSYH",
					Name:      "Widdy Arfiansyah",
					Address:   "Sukabumi",
					Phone:     "082291832488",
					CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
				},
				{
					ID:        2,
					Code:      "HYSFRADWDW",
					Name:      "Arfiansyah",
					Address:   "Cisaat",
					Phone:     "089922234214",
					CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
				},
			},
			expectedSupplierRepoFindAllError: nil,
			expectedSvcError:                 nil,
		},
		{
			name:                             "Number of suppliers is 0 or null",
			expectedSupplierRepoFindAll:      nil,
			expectedSvc:                      nil,
			expectedSupplierRepoFindAllError: nil,
			expectedSvcError:                 nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.SupplierRepositoryMock
			repo.On("FindAll", ctx).Return(tc.expectedSupplierRepoFindAll, tc.expectedSupplierRepoFindAllError)
			svc := NewSupplierService(&repo)
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

func TestSupplierService_FindByCode(t *testing.T) {
	testCases := []struct {
		name                                string
		request                             string
		expectedSupplierRepoFindByCode      *model.Supplier
		expectedSupplierRepoFindByCodeError error
		expectedSvc                         *response.SupplierResponse
		expectedSvcError                    error
	}{
		{
			name:    "Supplier exists with given code",
			request: "WDWDARFSYH",
			expectedSupplierRepoFindByCode: &model.Supplier{
				ID:      1,
				Code:    "WDWDARFSYH",
				Name:    "Widdy Arfiansyah",
				Address: "Sukabumi",
				Phone:   "082291832488",
			},
			expectedSvc: &response.SupplierResponse{
				ID:        1,
				Code:      "WDWDARFSYH",
				Name:      "Widdy Arfiansyah",
				Address:   "Sukabumi",
				Phone:     "082291832488",
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedSupplierRepoFindByCodeError: nil,
			expectedSvcError:                    nil,
		},
		{
			name:                                "Supplier doesnt exists with given code",
			request:                             "WDWDARFSYH",
			expectedSupplierRepoFindByCode:      nil,
			expectedSvc:                         nil,
			expectedSupplierRepoFindByCodeError: errors.New(response.NotFound),
			expectedSvcError:                    errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.SupplierRepositoryMock
			repo.On("FindByCodeWithAssociations", ctx, tc.request).Return(tc.expectedSupplierRepoFindByCode, tc.expectedSupplierRepoFindByCodeError)
			svc := NewSupplierService(&repo)
			result, err := svc.FindByCode(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestSupplierService_Create(t *testing.T) {

	testCases := []struct {
		name                            string
		request                         *request.CreateSupplierRequest
		expectedSupplierRepoCreate      *model.Supplier
		expectedSupplierRepoCreateError error
		expectedSvc                     *response.SupplierResponse
		expectedSvcError                error
	}{
		{
			name: "Create supplier with required fields",
			request: &request.CreateSupplierRequest{
				Name:    "Widdy Arfiansyah",
				Address: "Sukabumi",
				Phone:   "082291832488",
			},
			expectedSupplierRepoCreate: &model.Supplier{
				ID:      1,
				Code:    "WDWDARFSYH",
				Name:    "Widdy Arfiansyah",
				Address: "Sukabumi",
				Phone:   "082291832488",
			},
			expectedSvc: &response.SupplierResponse{
				ID:        1,
				Code:      "WDWDARFSYH",
				Name:      "Widdy Arfiansyah",
				Address:   "Sukabumi",
				Phone:     "082291832488",
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedSupplierRepoCreateError: nil,
			expectedSvcError:                nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.SupplierRepositoryMock
			repo.On("Create", ctx, mock.Anything).Return(tc.expectedSupplierRepoCreate, tc.expectedSupplierRepoCreateError)
			svc := NewSupplierService(&repo)
			result, err := svc.Create(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestSupplierService_Update(t *testing.T) {
	testCases := []struct {
		name                                string
		request                             *request.UpdateSupplierRequest
		requestSupplierRepoFindByCode       string
		expectedSupplierRepoFindByCode      *model.Supplier
		expectedSupplierRepoFindByCodeError error
		expectedSupplierRepoUpdate          *model.Supplier
		expectedSupplierRepoUpdateError     error
		expectedSvc                         *response.SupplierResponse
		expectedSvcError                    error
	}{
		{
			name: "Update supplier with required fields",
			request: &request.UpdateSupplierRequest{
				Code:    "WDWDARFSYH",
				Name:    "Arfian",
				Address: "Cisaat",
				Phone:   "089922238421",
			},
			requestSupplierRepoFindByCode: "WDWDARFSYH",
			expectedSupplierRepoFindByCode: &model.Supplier{
				ID:      1,
				Code:    "WDWDARFSYH",
				Name:    "Widdy Arfiansyah",
				Address: "Sukabumi",
				Phone:   "082291832488",
			},
			expectedSupplierRepoUpdate: &model.Supplier{
				ID:      1,
				Code:    "WDWDARFSYH",
				Name:    "Arfian",
				Address: "Cisaat",
				Phone:   "089922238421",
			},
			expectedSvc: &response.SupplierResponse{
				ID:        1,
				Code:      "WDWDARFSYH",
				Name:      "Arfian",
				Address:   "Cisaat",
				Phone:     "089922238421",
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedSupplierRepoFindByCodeError: nil,
			expectedSupplierRepoUpdateError:     nil,
			expectedSvcError:                    nil,
		},
		{
			name: "Supplier doesnt exists with given Code when updating data",
			request: &request.UpdateSupplierRequest{
				Code:    "WDWDARFSYH",
				Name:    "Arfian",
				Address: "Cisaat",
				Phone:   "089922238421",
			},
			requestSupplierRepoFindByCode:       "WDWDARFSYH",
			expectedSupplierRepoFindByCode:      nil,
			expectedSupplierRepoUpdate:          nil,
			expectedSvc:                         nil,
			expectedSupplierRepoUpdateError:     errors.New("getting an error"),
			expectedSupplierRepoFindByCodeError: errors.New(response.NotFound),
			expectedSvcError:                    errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.SupplierRepositoryMock
			repo.On("FindByCode", ctx, tc.requestSupplierRepoFindByCode).Return(tc.expectedSupplierRepoFindByCode, tc.expectedSupplierRepoFindByCodeError)
			repo.On("Update", ctx, mock.Anything).Return(tc.expectedSupplierRepoUpdate, tc.expectedSupplierRepoUpdateError)
			svc := NewSupplierService(&repo)
			result, err := svc.Update(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)

			// really changed
			if err == nil {
				assert.Equal(t, tc.expectedSupplierRepoFindByCode.Name, result.Name)
			}
		})
	}
}

func TestSupplierService_Delete(t *testing.T) {
	testCases := []struct {
		name                                string
		request                             string
		expectedSupplierRepoFindByCode      *model.Supplier
		expectedSupplierRepoFindByCodeError error
		expectedSupplierRepoDeleteError     error
		expectedSvcError                    error
	}{
		{
			name:    "Supplier exists with given Code",
			request: "WDWDARFSYH",
			expectedSupplierRepoFindByCode: &model.Supplier{
				ID:      1,
				Code:    "WDWDARFSYH",
				Name:    "Widdy Arfiansyah",
				Address: "Sukabumi",
				Phone:   "082291832488",
			},
			expectedSupplierRepoDeleteError:     nil,
			expectedSvcError:                    nil,
			expectedSupplierRepoFindByCodeError: nil,
		},
		{
			name:                                "Supplier doesnt exists with given Code when deleting data",
			request:                             "WDWDARFSYH",
			expectedSupplierRepoFindByCode:      nil,
			expectedSupplierRepoDeleteError:     errors.New("getting an error"),
			expectedSvcError:                    errors.New(response.NotFound),
			expectedSupplierRepoFindByCodeError: errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.SupplierRepositoryMock
			repo.On("FindByCode", ctx, tc.request).Return(tc.expectedSupplierRepoFindByCode, tc.expectedSupplierRepoFindByCodeError)
			repo.On("Delete", ctx, tc.request).Return(tc.expectedSupplierRepoDeleteError)
			svc := NewSupplierService(&repo)
			err := svc.Delete(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}
		})
	}
}
