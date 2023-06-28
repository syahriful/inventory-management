package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	response "inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	repository "inventory-management/backend/internal/repository/mock"
	"testing"
)

func TestProductQualityService_FindAll(t *testing.T) {
	testCases := []struct {
		name                                   string
		expectedProductQualityRepoFindAll      []*model.ProductQuality
		expectedProductQualityRepoFindAllError error
		expectedSvc                            []*response.ProductQualityResponse
		expectedSvcError                       error
	}{
		{
			name: "Number of products more than 1",
			expectedProductQualityRepoFindAll: []*model.ProductQuality{
				{
					ID:          1,
					ProductCode: "KKDJALS",
					Quality:     "Very Fresh",
					Price:       250000,
					Quantity:    5.5,
					Type:        "Increase",
				},
				{
					ID:          2,
					ProductCode: "KKDJALS",
					Quality:     "Fresh",
					Price:       150000,
					Quantity:    3.5,
					Type:        "Increase",
				},
			},
			expectedSvc: []*response.ProductQualityResponse{
				{
					ID:          1,
					ProductCode: "KKDJALS",
					Quality:     "Very Fresh",
					Price:       250000,
					Quantity:    5.5,
					Type:        "Increase",
				},
				{
					ID:          2,
					ProductCode: "KKDJALS",
					Quality:     "Fresh",
					Price:       150000,
					Quantity:    3.5,
					Type:        "Increase",
				},
			},
			expectedProductQualityRepoFindAllError: nil,
			expectedSvcError:                       nil,
		},
		{
			name:                                   "Number of products is 0 or null",
			expectedProductQualityRepoFindAll:      nil,
			expectedSvc:                            nil,
			expectedProductQualityRepoFindAllError: nil,
			expectedSvcError:                       nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoPQ repository.ProductQualityRepositoryMock
			var repP repository.ProductRepositoryMock
			repoPQ.On("FindAll", ctx).Return(tc.expectedProductQualityRepoFindAll, tc.expectedProductQualityRepoFindAllError)
			svc := NewProductQualityService(&repoPQ, &repP)
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

func TestProductQualityService_FindAllByProductCode(t *testing.T) {
	testCases := []struct {
		name                                                string
		request                                             string
		expectedProductRepoFindByCode                       *model.Product
		expectedProductRepoFindByCodeError                  error
		expectedProductQualityRepoFindAllByProductCode      []*model.ProductQuality
		expectedProductQualityRepoFindAllByProductCodeError error
		expectedSvc                                         *response.ProductQualityWithOwnProductResponse
		expectedSvcError                                    error
	}{
		{
			name:    "Product quality exists with given Code",
			request: "KKDJALS",
			expectedProductRepoFindByCode: &model.Product{
				ID:                  1,
				Code:                "KKDJALS",
				Name:                "Shrimp",
				UnitMassAcronym:     "g",
				UnitMassDescription: "gram",
			},
			expectedProductQualityRepoFindAllByProductCode: []*model.ProductQuality{
				{
					ID:          1,
					ProductCode: "KKDJALS",
					Quality:     "Very Fresh",
					Price:       250000,
					Quantity:    5.5,
					Type:        "Increase",
				},
				{
					ID:          2,
					ProductCode: "KKDJALS",
					Quality:     "Fresh",
					Price:       150000,
					Quantity:    3.5,
					Type:        "Increase",
				},
			},
			expectedSvc: &response.ProductQualityWithOwnProductResponse{
				Product: &response.ProductResponse{
					ID:                  1,
					Code:                "KKDJALS",
					Name:                "Shrimp",
					UnitMassAcronym:     "g",
					UnitMassDescription: "gram",
					CreatedAt:           "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt:           "0001-01-01 00:00:00 +0000 UTC",
				},
				ProductQualities: []*response.ProductQualityResponse{
					{
						ID:          1,
						ProductCode: "KKDJALS",
						Quality:     "Very Fresh",
						Price:       250000,
						Quantity:    5.5,
						Type:        "Increase",
					},
					{
						ID:          2,
						ProductCode: "KKDJALS",
						Quality:     "Fresh",
						Price:       150000,
						Quantity:    3.5,
						Type:        "Increase",
					},
				},
			},
			expectedProductRepoFindByCodeError:                  nil,
			expectedProductQualityRepoFindAllByProductCodeError: nil,
			expectedSvcError: nil,
		},
		{
			name:                          "Product quality doesnt exists with given Code",
			request:                       "KKDJALS",
			expectedProductRepoFindByCode: nil,
			expectedProductQualityRepoFindAllByProductCode: nil,
			expectedSvc:                        nil,
			expectedProductRepoFindByCodeError: errors.New(response.ErrorNotFound),
			expectedProductQualityRepoFindAllByProductCodeError: errors.New(response.ErrorNotFound),
			expectedSvcError: errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoPQ repository.ProductQualityRepositoryMock
			var repP repository.ProductRepositoryMock
			repP.On("FindByCodeWithAssociations", ctx, tc.request).Return(tc.expectedProductRepoFindByCode, tc.expectedProductRepoFindByCodeError)
			repoPQ.On("FindAllByProductCode", ctx, tc.request).Return(tc.expectedProductQualityRepoFindAllByProductCode, tc.expectedProductQualityRepoFindAllByProductCodeError)
			svc := NewProductQualityService(&repoPQ, &repP)
			result, err := svc.FindAllByProductCode(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestProductQualityService_FindByID(t *testing.T) {
	testCases := []struct {
		name                                    string
		request                                 int64
		expectedProductQualityRepoFindByID      *model.ProductQuality
		expectedProductQualityRepoFindByIDError error
		expectedSvc                             *response.ProductQualityResponse
		expectedSvcError                        error
	}{
		{
			name:    "Product quality exists with given ID",
			request: 1,
			expectedProductQualityRepoFindByID: &model.ProductQuality{
				ID:          1,
				ProductCode: "KKDJALS",
				Quality:     "Very Fresh",
				Price:       250000,
				Quantity:    5.5,
				Type:        "Increase",
				Product: &model.Product{
					ID:                  1,
					Code:                "KKDJALS",
					Name:                "Shrimp",
					UnitMassAcronym:     "g",
					UnitMassDescription: "gram",
				},
			},
			expectedSvc: &response.ProductQualityResponse{
				ID:          1,
				ProductCode: "KKDJALS",
				Quality:     "Very Fresh",
				Price:       250000,
				Quantity:    5.5,
				Type:        "Increase",
				Product: &response.ProductResponse{
					ID:                  1,
					Code:                "KKDJALS",
					Name:                "Shrimp",
					UnitMassAcronym:     "g",
					UnitMassDescription: "gram",
					CreatedAt:           "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt:           "0001-01-01 00:00:00 +0000 UTC",
				},
			},
			expectedProductQualityRepoFindByIDError: nil,
			expectedSvcError:                        nil,
		},
		{
			name:                                    "Product quality doesnt exists with given ID",
			request:                                 1,
			expectedProductQualityRepoFindByID:      nil,
			expectedProductQualityRepoFindByIDError: errors.New(response.ErrorNotFound),
			expectedSvc:                             nil,
			expectedSvcError:                        errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoPQ repository.ProductQualityRepositoryMock
			var repP repository.ProductRepositoryMock
			repoPQ.On("FindByIDWithAssociations", ctx, tc.request).Return(tc.expectedProductQualityRepoFindByID, tc.expectedProductQualityRepoFindByIDError)
			svc := NewProductQualityService(&repoPQ, &repP)
			result, err := svc.FindByID(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestProductQualityService_Delete(t *testing.T) {
	testCases := []struct {
		name                                    string
		request                                 int64
		expectedProductQualityRepoFindByID      *model.ProductQuality
		expectedProductQualityRepoFindByIDError error
		expectedProductQualityRepoDeleteError   error
		expectedSvcError                        error
	}{
		{
			name:    "Product quality exists with given Code",
			request: 1,
			expectedProductQualityRepoFindByID: &model.ProductQuality{
				ID:          1,
				ProductCode: "KKDJALS",
				Quality:     "Very Fresh",
				Price:       250000,
				Quantity:    5.5,
				Type:        "Increase",
			},
			expectedProductQualityRepoFindByIDError: nil,
			expectedProductQualityRepoDeleteError:   nil,
			expectedSvcError:                        nil,
		},
		{
			name:                                    "Product quality doesnt exists with given Code when deleting data",
			request:                                 1,
			expectedProductQualityRepoFindByID:      nil,
			expectedProductQualityRepoDeleteError:   errors.New("getting an error"),
			expectedProductQualityRepoFindByIDError: errors.New(response.ErrorNotFound),
			expectedSvcError:                        errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoPQ repository.ProductQualityRepositoryMock
			var repP repository.ProductRepositoryMock
			repoPQ.On("FindByID", ctx, tc.request).Return(tc.expectedProductQualityRepoFindByID, tc.expectedProductQualityRepoFindByIDError)
			repoPQ.On("Delete", ctx, tc.request).Return(tc.expectedProductQualityRepoDeleteError)
			svc := NewProductQualityService(&repoPQ, &repP)
			err := svc.Delete(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}
		})
	}
}
