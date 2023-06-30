package service

import (
	"context"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
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

func (service *SupplierService) FindAll(ctx context.Context, offset int, limit int) ([]*response.SupplierResponse, error) {
	suppliers, err := service.SupplierRepository.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var supplierResponses []*response.SupplierResponse
	for _, supplier := range suppliers {
		supplierResponses = append(supplierResponses, supplier.ToResponse())
	}

	return supplierResponses, nil
}

func (service *SupplierService) CountAll(ctx context.Context) (int64, error) {
	count, err := service.SupplierRepository.CountAll(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (service *SupplierService) FindByCode(ctx context.Context, code string) (*response.SupplierResponse, error) {
	supplier, err := service.SupplierRepository.FindByCodeWithAssociations(ctx, code)
	if err != nil {
		return nil, err
	}

	return supplier.ToResponseWithAssociations(), nil
}

func (service *SupplierService) Create(ctx context.Context, request *request.CreateSupplierRequest) (*response.SupplierResponse, error) {
	var supplierRequest model.Supplier
	supplierRequest.Name = request.Name
	supplierRequest.Address = request.Address
	supplierRequest.Phone = request.Phone

	supplier, err := service.SupplierRepository.Create(ctx, &supplierRequest)
	if err != nil {
		return nil, err
	}

	return supplier.ToResponse(), nil
}

func (service *SupplierService) Update(ctx context.Context, request *request.UpdateSupplierRequest) (*response.SupplierResponse, error) {
	checkSupplier, err := service.SupplierRepository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	checkSupplier.Name = request.Name
	checkSupplier.Address = request.Address
	checkSupplier.Phone = request.Phone

	supplier, err := service.SupplierRepository.Update(ctx, checkSupplier)
	if err != nil {
		return nil, err
	}

	return supplier.ToResponse(), nil
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
