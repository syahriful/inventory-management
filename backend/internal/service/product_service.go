package service

import (
	"context"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/internal/repository"
)

type ProductService struct {
	ProductRepository repository.ProductRepositoryContract
}

func NewProductService(productRepository repository.ProductRepositoryContract) ProductServiceContract {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (service *ProductService) FindAll(ctx context.Context, offset int, limit int) ([]*response.ProductResponse, error) {
	products, err := service.ProductRepository.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var productResponses []*response.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, product.ToResponse())
	}

	return productResponses, nil
}

func (service *ProductService) CountAll(ctx context.Context) (int64, error) {
	count, err := service.ProductRepository.CountAll(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (service *ProductService) FindByCode(ctx context.Context, code string) (*response.ProductResponse, error) {
	product, err := service.ProductRepository.FindByCodeWithAssociations(ctx, code)
	if err != nil {
		return nil, err
	}

	return product.ToResponseWithAssociations(), nil
}

func (service *ProductService) Create(ctx context.Context, request *request.CreateProductRequest) (*response.ProductResponse, error) {
	var productQualities []*model.ProductQuality
	for _, productQuality := range request.ProductQualities {
		var productQualityRequest model.ProductQuality
		productQualityRequest.Quality = productQuality.Quality
		productQualityRequest.Price = productQuality.Price
		productQualityRequest.Quantity = productQuality.Quantity
		productQualityRequest.Type = productQuality.Type

		productQualities = append(productQualities, &productQualityRequest)
	}

	var productRequest model.Product
	productRequest.Name = request.Name
	productRequest.UnitMassAcronym = request.UnitMassAcronym
	productRequest.UnitMassDescription = request.UnitMassDescription
	productRequest.ProductQualities = productQualities

	product, err := service.ProductRepository.Create(ctx, &productRequest)
	if err != nil {
		return nil, err
	}

	return product.ToResponse(), nil
}

func (service *ProductService) Update(ctx context.Context, request *request.UpdateProductRequest) (*response.ProductResponse, error) {
	var productQualities []*model.ProductQuality
	for _, productQuality := range request.ProductQualities {
		var productQualityRequest model.ProductQuality
		productQualityRequest.ID = productQuality.ID
		productQualityRequest.Quality = productQuality.Quality
		productQualityRequest.Price = productQuality.Price
		productQualityRequest.Quantity = productQuality.Quantity
		productQualityRequest.Type = productQuality.Type

		productQualities = append(productQualities, &productQualityRequest)
	}

	checkProduct, err := service.ProductRepository.FindByCodeWithAssociations(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	checkProduct.Name = request.Name
	checkProduct.UnitMassAcronym = request.UnitMassAcronym
	checkProduct.UnitMassDescription = request.UnitMassDescription
	checkProduct.ProductQualities = productQualities

	product, err := service.ProductRepository.Update(ctx, checkProduct)
	if err != nil {
		return nil, err
	}

	return product.ToResponse(), nil
}

func (service *ProductService) Delete(ctx context.Context, code string) error {
	checkProduct, err := service.ProductRepository.FindByCodeWithAssociations(ctx, code)
	if err != nil {
		return err
	}

	err = service.ProductRepository.Delete(ctx, checkProduct.Code)
	if err != nil {
		return err
	}

	return nil
}
