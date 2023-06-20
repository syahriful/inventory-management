package service

import (
	"context"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/internal/repository"
)

type SupplierService struct {
	SupplierRepository repository.SupplierRepositoryContract
}

func NewSupplierService(supplierRepository repository.SupplierRepositoryContract) SupplierServiceContract {
	return &SupplierService{
		SupplierRepository: supplierRepository,
	}
}

func (service *SupplierService) FindAll(ctx context.Context) ([]*response.SupplierResponse, error) {
	suppliers, err := service.SupplierRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var supplierResponses []*response.SupplierResponse
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, &response.SupplierResponse{
			ID:        supplier.ID,
			Code:      supplier.Code,
			Name:      supplier.Name,
			Address:   supplier.Address,
			Phone:     supplier.Phone,
			CreatedAt: supplier.CreatedAt.String(),
			UpdatedAt: supplier.UpdatedAt.String(),
		})
	}

	return supplierResponses, nil
}

func (service *SupplierService) FindByCode(ctx context.Context, code string) (*response.SupplierResponse, error) {
	supplier, err := service.SupplierRepository.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	var transactionResponses []*response.TransactionResponse
	for _, transaction := range supplier.Transactions {
		var productQualityResponse response.ProductQualityResponse
		productQualityResponse.ID = transaction.ProductQuality.ID
		productQualityResponse.Quality = transaction.ProductQuality.Quality
		productQualityResponse.Price = transaction.ProductQuality.Price

		transactionResponses = append(transactionResponses, &response.TransactionResponse{
			ID:               transaction.ID,
			Code:             transaction.Code,
			ProductQuality:   &productQualityResponse,
			ProductQualityID: transaction.ProductQualityID,
			SupplierCode:     transaction.SupplierCode,
			Description:      transaction.Description,
			Quantity:         transaction.Quantity,
			Type:             transaction.Type,
			CreatedAt:        transaction.CreatedAt.String(),
		})
	}

	return &response.SupplierResponse{
		ID:           supplier.ID,
		Code:         supplier.Code,
		Name:         supplier.Name,
		Address:      supplier.Address,
		Phone:        supplier.Phone,
		CreatedAt:    supplier.CreatedAt.String(),
		UpdatedAt:    supplier.UpdatedAt.String(),
		Transactions: transactionResponses,
	}, nil
}

func (service *SupplierService) Create(ctx context.Context, request *request.CreateSupplierRequest) (*response.SupplierResponse, error) {
	supplier, err := service.SupplierRepository.Create(ctx, &model.Supplier{
		Name:    request.Name,
		Address: request.Address,
		Phone:   request.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &response.SupplierResponse{
		ID:        supplier.ID,
		Code:      supplier.Code,
		Name:      supplier.Name,
		Address:   supplier.Address,
		Phone:     supplier.Phone,
		CreatedAt: supplier.CreatedAt.String(),
		UpdatedAt: supplier.UpdatedAt.String(),
	}, nil
}

func (service *SupplierService) Update(ctx context.Context, request *request.UpdateSupplierRequest) (*response.SupplierResponse, error) {
	checkSupplier, err := service.SupplierRepository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	supplier, err := service.SupplierRepository.Update(ctx, &model.Supplier{
		Code:    checkSupplier.Code,
		Name:    request.Name,
		Address: request.Address,
		Phone:   request.Phone,
	})
	if err != nil {
		return nil, err
	}

	return &response.SupplierResponse{
		ID:        supplier.ID,
		Code:      supplier.Code,
		Name:      supplier.Name,
		Address:   supplier.Address,
		Phone:     supplier.Phone,
		CreatedAt: supplier.CreatedAt.String(),
		UpdatedAt: supplier.UpdatedAt.String(),
	}, nil
}

func (service *SupplierService) Delete(ctx context.Context, code string) error {
	checkSupplier, err := service.SupplierRepository.FindByCode(ctx, code)
	if err != nil {
		return err
	}

	err = service.SupplierRepository.Delete(ctx, checkSupplier.Code)
	if err != nil {
		return err
	}

	return nil
}
