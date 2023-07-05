package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	repository "inventory-management/backend/internal/repository/mock"
	"inventory-management/backend/util"
	"testing"
)

func TestTransactionService_FindAll(t *testing.T) {
	testCases := []struct {
		name                                string
		expectedTransactionRepoFindAll      []*model.Transaction
		expectedTransactionRepoFindAllError error
		expectedSvc                         []*response.TransactionResponse
		expectedSvcError                    error
	}{
		{
			name: "Number of transactions more than 1",
			expectedTransactionRepoFindAll: []*model.Transaction{
				{
					ID:               1,
					Code:             "WDWDARFSYH",
					ProductQualityID: 1,
					SupplierCode:     util.ToPointerString("SUP001"),
					Description:      util.ToPointerString("Pembelian barang dari supplier"),
					Quantity:         23,
					Type:             "IN",
					UnitMassAcronym:  "kg",
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
				},
			},
			expectedSvc: []*response.TransactionResponse{
				{
					ID:               1,
					Code:             "WDWDARFSYH",
					ProductQualityID: 1,
					SupplierCode:     util.ToPointerString("SUP001"),
					Description:      util.ToPointerString("Pembelian barang dari supplier"),
					Quantity:         23,
					Type:             "IN",
					UnitMassAcronym:  "kg",
					CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
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
					CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
				},
			},
			expectedTransactionRepoFindAllError: nil,
			expectedSvcError:                    nil,
		},
		{
			name:                                "Number of transactions is 0 or null",
			expectedTransactionRepoFindAll:      nil,
			expectedSvc:                         nil,
			expectedTransactionRepoFindAllError: nil,
			expectedSvcError:                    nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoT repository.TransactionRepositoryMock
			var repoPQ repository.ProductQualityRepositoryMock
			var repoTx repository.TxTransactionRepositoryMock
			repoT.On("FindAll", ctx, 0, 10).Return(tc.expectedTransactionRepoFindAll, tc.expectedTransactionRepoFindAllError)
			svc := NewTransactionService(&repoT, &repoPQ, &repoTx)
			result, err := svc.FindAll(ctx, 0, 10)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
			assert.Equal(t, len(tc.expectedSvc), len(result))
		})
	}
}

func TestTransactionService_FindAllBySupplierCode(t *testing.T) {
	testCases := []struct {
		name                                              string
		request                                           string
		expectedTransactionRepoFindAllBySupplierCode      []*model.Transaction
		expectedTransactionRepoFindAllBySupplierCodeError error
		expectedSvc                                       []*response.TransactionResponse
		expectedSvcError                                  error
	}{
		{
			name:    "Number of transactions more than 1",
			request: "SUP001",
			expectedTransactionRepoFindAllBySupplierCode: []*model.Transaction{
				{
					ID:               1,
					Code:             "WDWDARFSYH",
					ProductQualityID: 1,
					SupplierCode:     util.ToPointerString("SUP001"),
					Description:      util.ToPointerString("Pembelian barang dari supplier"),
					Quantity:         23,
					Type:             "IN",
					UnitMassAcronym:  "kg",
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
				},
			},
			expectedSvc: []*response.TransactionResponse{
				{
					ID:               1,
					Code:             "WDWDARFSYH",
					ProductQualityID: 1,
					SupplierCode:     util.ToPointerString("SUP001"),
					Description:      util.ToPointerString("Pembelian barang dari supplier"),
					Quantity:         23,
					Type:             "IN",
					UnitMassAcronym:  "kg",
					CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
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
					CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
				},
			},
			expectedTransactionRepoFindAllBySupplierCodeError: nil,
			expectedSvcError: nil,
		},
		{
			name:    "Number of transactions is 0 or null",
			request: "SUP001",
			expectedTransactionRepoFindAllBySupplierCode: nil,
			expectedSvc: nil,
			expectedTransactionRepoFindAllBySupplierCodeError: nil,
			expectedSvcError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoT repository.TransactionRepositoryMock
			var repoPQ repository.ProductQualityRepositoryMock
			var repoTx repository.TxTransactionRepositoryMock
			repoT.On("FindAllBySupplierCode", ctx, tc.request).Return(tc.expectedTransactionRepoFindAllBySupplierCode, tc.expectedTransactionRepoFindAllBySupplierCodeError)
			svc := NewTransactionService(&repoT, &repoPQ, &repoTx)
			result, err := svc.FindAllBySupplierCode(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
			assert.Equal(t, len(tc.expectedSvc), len(result))
		})
	}
}

func TestTransactionService_FindAllByCustomerCode(t *testing.T) {
	testCases := []struct {
		name                                              string
		request                                           string
		expectedTransactionRepoFindAllByCustomerCode      []*model.Transaction
		expectedTransactionRepoFindAllByCustomerCodeError error
		expectedSvc                                       []*response.TransactionResponse
		expectedSvcError                                  error
	}{
		{
			name:    "Number of transactions more than 1",
			request: "CUS001",
			expectedTransactionRepoFindAllByCustomerCode: []*model.Transaction{
				{
					ID:               1,
					Code:             "WDWDARFSYH",
					ProductQualityID: 1,
					CustomerCode:     util.ToPointerString("CUS001"),
					Description:      util.ToPointerString("Pembelian barang dari customer"),
					Quantity:         23,
					Type:             "IN",
					UnitMassAcronym:  "kg",
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
				},
			},
			expectedSvc: []*response.TransactionResponse{
				{
					ID:               1,
					Code:             "WDWDARFSYH",
					ProductQualityID: 1,
					CustomerCode:     util.ToPointerString("CUS001"),
					Description:      util.ToPointerString("Pembelian barang dari customer"),
					Quantity:         23,
					Type:             "IN",
					UnitMassAcronym:  "kg",
					CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
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
					CreatedAt:        "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt:        "0001-01-01 07:00:00 +0700 +07",
				},
			},
			expectedTransactionRepoFindAllByCustomerCodeError: nil,
			expectedSvcError: nil,
		},
		{
			name:    "Number of transactions is 0 or null",
			request: "CUS001",
			expectedTransactionRepoFindAllByCustomerCode: nil,
			expectedSvc: nil,
			expectedTransactionRepoFindAllByCustomerCodeError: nil,
			expectedSvcError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoT repository.TransactionRepositoryMock
			var repoPQ repository.ProductQualityRepositoryMock
			var repoTx repository.TxTransactionRepositoryMock
			repoT.On("FindAllByCustomerCode", ctx, tc.request).Return(tc.expectedTransactionRepoFindAllByCustomerCode, tc.expectedTransactionRepoFindAllByCustomerCodeError)
			svc := NewTransactionService(&repoT, &repoPQ, &repoTx)
			result, err := svc.FindAllByCustomerCode(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
			assert.Equal(t, len(tc.expectedSvc), len(result))
		})
	}
}

func TestTransactionService_FindByCode(t *testing.T) {
	testCases := []struct {
		name                                   string
		request                                string
		expectedTransactionRepoFindByCode      *model.Transaction
		expectedTransactionRepoFindByCodeError error
		expectedSvc                            *response.TransactionResponse
		expectedSvcError                       error
	}{
		{
			name:    "Transaction customer is exists with given code with type IN",
			request: "WDWDARFSYH",
			expectedTransactionRepoFindByCode: &model.Transaction{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				ProductQuality: &model.ProductQuality{
					ID:          1,
					ProductCode: "SHRMP",
					Quality:     "Fresh",
					Quantity:    23,
					Price:       100000,
					Type:        "increase",
					Product: &model.Product{
						ID:                  1,
						Code:                "SHRMP",
						Name:                "Shrimp",
						UnitMassAcronym:     "kg",
						UnitMassDescription: "kilogram",
					},
				},
				SupplierCode: util.ToPointerString("SUP001"),
				Supplier: &model.Supplier{
					ID:      1,
					Code:    "SUP001",
					Name:    "PT. Sumber Makmur",
					Address: "Jl. Raya Bogor",
					Phone:   "08123456789",
				},
				Description:     util.ToPointerString("Menambahkan stok barang dari supplier"),
				Quantity:        23,
				Type:            "IN",
				UnitMassAcronym: "kg",
			},
			expectedSvc: &response.TransactionResponse{
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
			expectedTransactionRepoFindByCodeError: nil,
			expectedSvcError:                       nil,
		},
		{
			name:    "Transaction customer is exists with given code with type OUT",
			request: "WDWDARFSYH",
			expectedTransactionRepoFindByCode: &model.Transaction{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				ProductQuality: &model.ProductQuality{
					ID:          1,
					ProductCode: "SHRMP",
					Quality:     "Fresh",
					Quantity:    23,
					Price:       100000,
					Type:        "increase",
					Product: &model.Product{
						ID:                  1,
						Code:                "SHRMP",
						Name:                "Shrimp",
						UnitMassAcronym:     "kg",
						UnitMassDescription: "kilogram",
					},
				},
				CustomerCode: util.ToPointerString("CUS001"),
				Customer: &model.Customer{
					ID:   1,
					Code: "CUS001",
					Name: "PT. Sumber Makmur",
				},
				Description:     util.ToPointerString("Pembelian barang dari customer"),
				Quantity:        23,
				Type:            "OUT",
				UnitMassAcronym: "kg",
			},
			expectedSvc: &response.TransactionResponse{
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
			expectedTransactionRepoFindByCodeError: nil,
			expectedSvcError:                       nil,
		},
		{
			name:    "Transaction exists with given code with type TRANSFER",
			request: "WDWDARFSYH",
			expectedTransactionRepoFindByCode: &model.Transaction{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				ProductQuality: &model.ProductQuality{
					ID:          1,
					ProductCode: "SHRMP",
					Quality:     "Fresh",
					Quantity:    23,
					Price:       100000,
					Type:        "increase",
					Product: &model.Product{
						ID:                  1,
						Code:                "SHRMP",
						Name:                "Shrimp",
						UnitMassAcronym:     "kg",
						UnitMassDescription: "kilogram",
					},
				},
				ProductQualityIDTransferred: util.ToPointerInt64(2),
				ProductQualityTransferred: &model.ProductQuality{
					ID:          2,
					ProductCode: "SHRMP",
					Quality:     "Expired",
					Quantity:    2,
					Price:       1000,
					Type:        "decrease",
					Product: &model.Product{
						ID:                  1,
						Code:                "SHRMP",
						Name:                "Shrimp",
						UnitMassAcronym:     "kg",
						UnitMassDescription: "kilogram",
					},
				},
				Description:     util.ToPointerString("Barang sudah expired dan dibuang"),
				Quantity:        20,
				Type:            "TRANSFER",
				UnitMassAcronym: "kg",
			},
			expectedSvc: &response.TransactionResponse{
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
			expectedTransactionRepoFindByCodeError: nil,
			expectedSvcError:                       nil,
		},
		{
			name:                                   "Transaction doesnt exists with given code",
			request:                                "WDWDARFSYH",
			expectedTransactionRepoFindByCode:      nil,
			expectedSvc:                            nil,
			expectedTransactionRepoFindByCodeError: errors.New(response.ErrorNotFound),
			expectedSvcError:                       errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoT repository.TransactionRepositoryMock
			var repoPQ repository.ProductQualityRepositoryMock
			var repoTx repository.TxTransactionRepositoryMock
			repoT.On("FindByCodeWithAssociations", ctx, tc.request).Return(tc.expectedTransactionRepoFindByCode, tc.expectedTransactionRepoFindByCodeError)
			svc := NewTransactionService(&repoT, &repoPQ, &repoTx)
			result, err := svc.FindByCode(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestTransactionService_Create(t *testing.T) {
	testCases := []struct {
		name                               string
		request                            *request.CreateTransactionRequest
		expectedTransactionRepoCreate      *model.Transaction
		expectedSvc                        *response.TransactionResponse
		expectedSvcError                   error
		expectedTransactionRepoCreateError error
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
			expectedTransactionRepoCreate: &model.Transaction{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("CUS001"),
				Description:      util.ToPointerString("Pembelian barang dari customer"),
				Quantity:         23,
				Type:             "OUT",
				UnitMassAcronym:  "kg",
			},
			expectedSvc: &response.TransactionResponse{
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
			expectedSvcError:                   nil,
			expectedTransactionRepoCreateError: nil,
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
			expectedTransactionRepoCreate: &model.Transaction{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
			},
			expectedSvc: &response.TransactionResponse{
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
			expectedSvcError:                   nil,
			expectedTransactionRepoCreateError: nil,
		},
		{
			name: "Transaction doesnt exists with given Product Quality ID when creating data",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("CUS001"),
				Description:      util.ToPointerString("Pembelian barang dari customer"),
				Quantity:         23,
				Type:             "OUT",
				UnitMassAcronym:  "kg",
			},
			expectedTransactionRepoCreate:      nil,
			expectedSvc:                        nil,
			expectedSvcError:                   errors.New(response.ErrorNotFound),
			expectedTransactionRepoCreateError: errors.New(response.ErrorNotFound),
		},
		{
			name: "Failed create transaction",
			request: &request.CreateTransactionRequest{
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("CUS001"),
				Description:      util.ToPointerString("Pembelian barang dari customer"),
				Quantity:         23,
				Type:             "OUT",
				UnitMassAcronym:  "kg",
			},
			expectedTransactionRepoCreate:      nil,
			expectedSvc:                        nil,
			expectedSvcError:                   errors.New("getting an error"),
			expectedTransactionRepoCreateError: errors.New("getting an error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoT repository.TransactionRepositoryMock
			var repoPQ repository.ProductQualityRepositoryMock
			var repoTx repository.TxTransactionRepositoryMock
			repoTx.On("Create", ctx, tc.request).Return(tc.expectedTransactionRepoCreate, tc.expectedTransactionRepoCreateError)
			svc := NewTransactionService(&repoT, &repoPQ, &repoTx)
			result, err := svc.Create(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestTransactionService_Update(t *testing.T) {
	testCases := []struct {
		name                               string
		request                            *request.UpdateTransactionRequest
		expectedTransactionRepoUpdate      *model.Transaction
		expectedSvc                        *response.TransactionResponse
		expectedSvcError                   error
		expectedTransactionRepoUpdateError error
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
			expectedTransactionRepoUpdate: &model.Transaction{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("CUS001"),
				Description:      util.ToPointerString("Ubah pembelian barang dari customer"),
				Quantity:         23,
				Type:             "OUT",
				UnitMassAcronym:  "kg",
			},
			expectedSvc: &response.TransactionResponse{
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
			expectedSvcError:                   nil,
			expectedTransactionRepoUpdateError: nil,
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
			expectedTransactionRepoUpdate: &model.Transaction{
				ID:               1,
				Code:             "WDWDARFSYH",
				ProductQualityID: 1,
				CustomerCode:     util.ToPointerString("SUP001"),
				Description:      util.ToPointerString("Ubah pembelian barang dari supplier"),
				Quantity:         23,
				Type:             "IN",
				UnitMassAcronym:  "kg",
			},
			expectedSvc: &response.TransactionResponse{
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
			expectedSvcError:                   nil,
			expectedTransactionRepoUpdateError: nil,
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
			expectedSvc:                        nil,
			expectedSvcError:                   errors.New(response.ErrorNotFound),
			expectedTransactionRepoUpdateError: errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoT repository.TransactionRepositoryMock
			var repoPQ repository.ProductQualityRepositoryMock
			var repoTx repository.TxTransactionRepositoryMock
			repoTx.On("Update", ctx, tc.request).Return(tc.expectedTransactionRepoUpdate, tc.expectedTransactionRepoUpdateError)
			svc := NewTransactionService(&repoT, &repoPQ, &repoTx)
			result, err := svc.Update(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestTransactionService_TransferStock(t *testing.T) {
	testCases := []struct {
		name                                      string
		request                                   *request.TransferStockTransactionRequest
		expectedTransactionRepoTransferStock      *model.Transaction
		expectedSvc                               *response.TransactionResponse
		expectedSvcError                          error
		expectedTransactionRepoTransferStockError error
	}{
		{
			name: "Transfer stock with required fields",
			request: &request.TransferStockTransactionRequest{
				ProductQualityID:            1,
				ProductQualityIDTransferred: 2,
				Description:                 util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:                    23,
			},
			expectedTransactionRepoTransferStock: &model.Transaction{
				ID:                          1,
				Code:                        "WDWDARFSYH",
				ProductQualityID:            1,
				ProductQualityIDTransferred: util.ToPointerInt64(2),
				Description:                 util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:                    23,
				Type:                        "TRANSFER",
				UnitMassAcronym:             "kg",
			},
			expectedSvc: &response.TransactionResponse{
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
			expectedSvcError: nil,
			expectedTransactionRepoTransferStockError: nil,
		},
		{
			name: "Transaction doesnt exists with given ProductQualityID when transferring stock",
			request: &request.TransferStockTransactionRequest{
				ProductQualityID:            1,
				ProductQualityIDTransferred: 2,
				Description:                 util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:                    23,
			},
			expectedSvcError: errors.New(response.ErrorNotFound),
			expectedTransactionRepoTransferStockError: errors.New(response.ErrorNotFound),
		},
		{
			name: "Transaction doesnt exists with given ProductQualityIDTransferred when transferring stock",
			request: &request.TransferStockTransactionRequest{
				ProductQualityID:            1,
				ProductQualityIDTransferred: 2,
				Description:                 util.ToPointerString("Transfer stock dari product quality 1 ke product quality 2"),
				Quantity:                    23,
			},
			expectedSvcError: errors.New(response.ErrorNotFound),
			expectedTransactionRepoTransferStockError: errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoT repository.TransactionRepositoryMock
			var repoPQ repository.ProductQualityRepositoryMock
			var repoTx repository.TxTransactionRepositoryMock
			repoTx.On("TransferStock", ctx, tc.request).Return(tc.expectedTransactionRepoTransferStock, tc.expectedTransactionRepoTransferStockError)
			svc := NewTransactionService(&repoT, &repoPQ, &repoTx)
			result, err := svc.TransferStock(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestTransactionService_Delete(t *testing.T) {
	testCases := []struct {
		name                               string
		request                            string
		expectedSvcError                   error
		expectedTransactionRepoDeleteError error
	}{
		{
			name:                               "Transaction exists with given Code",
			request:                            "WDWDARFSYH",
			expectedSvcError:                   nil,
			expectedTransactionRepoDeleteError: nil,
		},
		{
			name:                               "Transaction doesnt exists with given Code when deleting data",
			request:                            "WDWDARFSYH",
			expectedSvcError:                   errors.New(response.ErrorNotFound),
			expectedTransactionRepoDeleteError: errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repoT repository.TransactionRepositoryMock
			var repoPQ repository.ProductQualityRepositoryMock
			var repoTx repository.TxTransactionRepositoryMock
			repoTx.On("Delete", ctx, tc.request).Return(tc.expectedTransactionRepoDeleteError)
			svc := NewTransactionService(&repoT, &repoPQ, &repoTx)
			err := svc.Delete(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}
		})
	}
}
