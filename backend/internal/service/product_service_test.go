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

func TestProductService_FindAll(t *testing.T) {
	testCases := []struct {
		name                            string
		expectedProductRepoFindAll      []*model.Product
		expectedProductRepoFindAllError error
		expectedSvc                     []*response.ProductResponse
		expectedSvcError                error
	}{
		{
			name: "Number of products more than 1",
			expectedProductRepoFindAll: []*model.Product{
				{
					ID:                  1,
					Code:                "KKDJALS",
					Name:                "Shark",
					UnitMassAcronym:     "kg",
					UnitMassDescription: "kilogram",
				},
				{
					ID:                  2,
					Code:                "WIJNWJS",
					Name:                "Shrimp",
					UnitMassAcronym:     "g",
					UnitMassDescription: "gram",
				},
			},
			expectedSvc: []*response.ProductResponse{
				{
					ID:                  1,
					Code:                "KKDJALS",
					Name:                "Shark",
					UnitMassAcronym:     "kg",
					UnitMassDescription: "kilogram",
					CreatedAt:           "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt:           "0001-01-01 00:00:00 +0000 UTC",
				},
				{
					ID:                  2,
					Code:                "WIJNWJS",
					Name:                "Shrimp",
					UnitMassAcronym:     "g",
					UnitMassDescription: "gram",
					CreatedAt:           "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt:           "0001-01-01 00:00:00 +0000 UTC",
				},
			},
			expectedProductRepoFindAllError: nil,
			expectedSvcError:                nil,
		},
		{
			name:                            "Number of products is 0 or null",
			expectedProductRepoFindAll:      nil,
			expectedSvc:                     nil,
			expectedProductRepoFindAllError: nil,
			expectedSvcError:                nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.ProductRepositoryMock
			repo.On("FindAll", ctx).Return(tc.expectedProductRepoFindAll, tc.expectedProductRepoFindAllError)
			svc := NewProductService(&repo)
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

func TestProductService_FindByCode(t *testing.T) {
	testCases := []struct {
		name                               string
		request                            string
		expectedProductRepoFindByCode      *model.Product
		expectedProductRepoFindByCodeError error
		expectedSvc                        *response.ProductResponse
		expectedSvcError                   error
	}{
		{
			name:    "Product exists with given Code",
			request: "KKDJALS",
			expectedProductRepoFindByCode: &model.Product{
				ID:                  1,
				Code:                "KKDJALS",
				Name:                "Shark",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
			},
			expectedSvc: &response.ProductResponse{
				ID:                  1,
				Code:                "KKDJALS",
				Name:                "Shark",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				CreatedAt:           "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt:           "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedProductRepoFindByCodeError: nil,
			expectedSvcError:                   nil,
		},
		{
			name:                               "Product doesnt exists with given Code",
			request:                            "KKDJALS",
			expectedProductRepoFindByCode:      nil,
			expectedSvc:                        nil,
			expectedProductRepoFindByCodeError: errors.New(response.NotFound),
			expectedSvcError:                   errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.ProductRepositoryMock
			repo.On("FindByCode", ctx, tc.request).Return(tc.expectedProductRepoFindByCode, tc.expectedProductRepoFindByCodeError)
			svc := NewProductService(&repo)
			result, err := svc.FindByCode(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestProductService_Create(t *testing.T) {
	testCases := []struct {
		name                           string
		request                        *request.CreateProductRequest
		expectedProductRepoCreate      *model.Product
		expectedProductRepoCreateError error
		expectedSvc                    *response.ProductResponse
		expectedSvcError               error
	}{
		{
			name: "Create product with required fields",
			request: &request.CreateProductRequest{
				Name:                "Shark",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
			},
			expectedProductRepoCreate: &model.Product{
				ID:                  1,
				Code:                "KKDJALS",
				Name:                "Shark",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
			},
			expectedSvc: &response.ProductResponse{
				ID:                  1,
				Code:                "KKDJALS",
				Name:                "Shark",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				CreatedAt:           "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt:           "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedProductRepoCreateError: nil,
			expectedSvcError:               nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.ProductRepositoryMock
			repo.On("Create", ctx, mock.Anything).Return(tc.expectedProductRepoCreate, tc.expectedProductRepoCreateError)
			svc := NewProductService(&repo)
			result, err := svc.Create(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestProductService_Update(t *testing.T) {
	testCases := []struct {
		name                               string
		request                            *request.UpdateProductRequest
		requestProductRepoFindByCode       string
		expectedProductRepoFindByCode      *model.Product
		expectedProductRepoFindByCodeError error
		expectedProductRepoUpdate          *model.Product
		expectedProductRepoUpdateError     error
		expectedSvc                        *response.ProductResponse
		expectedSvcError                   error
	}{
		{
			name: "Update product with required fields",
			request: &request.UpdateProductRequest{
				Code:                "KKJANSM",
				Name:                "Shrimp",
				UnitMassAcronym:     "g",
				UnitMassDescription: "gram",
				ProductQualities: []*request.UpdateProductQualityRequest{
					{
						ID:       1,
						Quality:  "Fresh",
						Price:    250000,
						Quantity: 4.5,
						Type:     "Increase",
					},
				},
			},
			requestProductRepoFindByCode: "KKJANSM",
			expectedProductRepoFindByCode: &model.Product{
				ID:                  1,
				Code:                "KKJANSM",
				Name:                "Shark",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
				ProductQualities: []*model.ProductQuality{
					{
						ID:       1,
						Quality:  "Very Fresh",
						Price:    150000,
						Quantity: 3.5,
						Type:     "Increase",
					},
				},
			},
			expectedProductRepoUpdate: &model.Product{
				ID:                  1,
				Code:                "KKJANSM",
				Name:                "Shrimp",
				UnitMassAcronym:     "g",
				UnitMassDescription: "gram",
				ProductQualities: []*model.ProductQuality{
					{
						ID:       1,
						Quality:  "Fresh",
						Price:    250000,
						Quantity: 4.5,
						Type:     "Increase",
					},
				},
			},
			expectedSvc: &response.ProductResponse{
				ID:                  1,
				Code:                "KKJANSM",
				Name:                "Shrimp",
				UnitMassAcronym:     "g",
				UnitMassDescription: "gram",
				CreatedAt:           "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt:           "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedProductRepoUpdateError:     nil,
			expectedSvcError:                   nil,
			expectedProductRepoFindByCodeError: nil,
		},
		{
			name: "Product doesnt exists with given Code when updating data",
			request: &request.UpdateProductRequest{
				Code:                "KKDJALS",
				Name:                "Shark",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
			},
			requestProductRepoFindByCode:       "KKDJALS",
			expectedProductRepoFindByCode:      nil,
			expectedProductRepoUpdate:          nil,
			expectedSvc:                        nil,
			expectedProductRepoUpdateError:     errors.New("getting an error"),
			expectedSvcError:                   errors.New(response.NotFound),
			expectedProductRepoFindByCodeError: errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.ProductRepositoryMock
			repo.On("FindByCode", ctx, tc.requestProductRepoFindByCode).Return(tc.expectedProductRepoFindByCode, tc.expectedProductRepoFindByCodeError)
			repo.On("Update", ctx, mock.Anything).Return(tc.expectedProductRepoUpdate, tc.expectedProductRepoUpdateError)
			svc := NewProductService(&repo)
			result, err := svc.Update(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)

			// really changed
			if err == nil {
				assert.Equal(t, tc.expectedProductRepoFindByCode.Name, result.Name)
			}
		})
	}
}

func TestProductService_Delete(t *testing.T) {
	testCases := []struct {
		name                               string
		request                            string
		expectedProductRepoFindByCode      *model.Product
		expectedProductRepoFindByCodeError error
		expectedProductRepoUpdateError     error
		expectedSvcError                   error
	}{
		{
			name:    "Product exists with given Code",
			request: "KKDJALS",
			expectedProductRepoFindByCode: &model.Product{
				ID:                  1,
				Code:                "KKDJALS",
				Name:                "Shark",
				UnitMassAcronym:     "kg",
				UnitMassDescription: "kilogram",
			},
			expectedProductRepoUpdateError:     nil,
			expectedSvcError:                   nil,
			expectedProductRepoFindByCodeError: nil,
		},
		{
			name:                               "Product doesnt exists with given Code when deleting data",
			request:                            "KKDJALS",
			expectedProductRepoFindByCode:      nil,
			expectedProductRepoUpdateError:     errors.New("getting an error"),
			expectedSvcError:                   errors.New(response.NotFound),
			expectedProductRepoFindByCodeError: errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.ProductRepositoryMock
			repo.On("FindByCode", ctx, tc.request).Return(tc.expectedProductRepoFindByCode, tc.expectedProductRepoFindByCodeError)
			repo.On("Delete", ctx, tc.request).Return(tc.expectedProductRepoUpdateError)
			svc := NewProductService(&repo)
			err := svc.Delete(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}
		})
	}
}
