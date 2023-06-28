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

func (repository *ProductService) FindAll(ctx context.Context) ([]*response.ProductResponse, error) {
	products, err := repository.ProductRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var productResponses []*response.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, product.ToResponse())
	}

	return productResponses, nil
}

func (repository *ProductService) FindByCode(ctx context.Context, code string) (*response.ProductResponse, error) {
	product, err := repository.ProductRepository.FindByCodeWithAssociations(ctx, code)
	if err != nil {
		return nil, err
	}

	return product.ToResponseWithAssociations(), nil
}

func (repository *ProductService) Create(ctx context.Context, request *request.CreateProductRequest) (*response.ProductResponse, error) {
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

	product, err := repository.ProductRepository.Create(ctx, &productRequest)
	if err != nil {
		return nil, err
	}

	return product.ToResponse(), nil
}

func (repository *ProductService) Update(ctx context.Context, request *request.UpdateProductRequest) (*response.ProductResponse, error) {
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

	checkProduct, err := repository.ProductRepository.FindByCodeWithAssociations(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	checkProduct.Name = request.Name
	checkProduct.UnitMassAcronym = request.UnitMassAcronym
	checkProduct.UnitMassDescription = request.UnitMassDescription
	checkProduct.ProductQualities = productQualities

	product, err := repository.ProductRepository.Update(ctx, checkProduct)
	if err != nil {
		return nil, err
	}

	return product.ToResponse(), nil
}

func (repository *ProductService) Delete(ctx context.Context, code string) error {
	checkProduct, err := repository.ProductRepository.FindByCodeWithAssociations(ctx, code)
	if err != nil {
		return err
	}

	err = repository.ProductRepository.Delete(ctx, checkProduct.Code)
	if err != nil {
		return err
	}

	return nil
}
