package service

import (
	"context"
	response "inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/repository"
	"sync"
)

type ProductQualityService struct {
	ProductQualityRepository repository.ProductQualityRepositoryContract
	ProductRepository        repository.ProductRepositoryContract
}

func NewProductQualityService(productQualityRepository repository.ProductQualityRepositoryContract, productRepository repository.ProductRepositoryContract) ProductQualityServiceContract {
	return &ProductQualityService{
		ProductQualityRepository: productQualityRepository,
		ProductRepository:        productRepository,
	}
}

func (service *ProductQualityService) FindAll(ctx context.Context) ([]*response.ProductQualityResponse, error) {
	productQualities, err := service.ProductQualityRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var productQualityResponses []*response.ProductQualityResponse
	for _, productQuality := range productQualities {
		productQualityResponses = append(productQualityResponses, productQuality.ToResponse())
	}

	return productQualityResponses, nil
}

func (service *ProductQualityService) FindAllByProductCode(ctx context.Context, productCode string) (*response.ProductQualityWithOwnProductResponse, error) {
	var productResponse *response.ProductResponse
	var productQualityResponses []*response.ProductQualityResponse
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var err error

	wg.Add(2)
	go func() {
		defer wg.Done()
		product, errorProduct := service.ProductRepository.FindByCode(ctx, productCode)
		if errorProduct != nil {
			mutex.Lock()
			err = errorProduct
			mutex.Unlock()
			return
		}

		productResponse = product.ToResponse()
	}()

	go func() {
		defer wg.Done()
		productQualities, errorProduct := service.ProductQualityRepository.FindAllByProductCode(ctx, productCode)
		if errorProduct != nil {
			mutex.Lock()
			err = errorProduct
			mutex.Unlock()
			return
		}

		for _, productQuality := range productQualities {
			productQualityResponses = append(productQualityResponses, productQuality.ToResponse())
		}
	}()
	wg.Wait()
	if err != nil {
		return nil, err
	}

	return &response.ProductQualityWithOwnProductResponse{
		Product:          productResponse,
		ProductQualities: productQualityResponses,
	}, nil
}

func (service *ProductQualityService) FindByID(ctx context.Context, id int64) (*response.ProductQualityResponse, error) {
	productQuality, err := service.ProductQualityRepository.FindByIDWithAssociations(ctx, id)
	if err != nil {
		return nil, err
	}

	return productQuality.ToResponseWithAssociations(), nil
}

func (service *ProductQualityService) Delete(ctx context.Context, id int64) error {
	checkProductQuality, err := service.ProductQualityRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = service.ProductQualityRepository.Delete(ctx, checkProductQuality.ID)
	if err != nil {
		return err
	}

	return nil
}
