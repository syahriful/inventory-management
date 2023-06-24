package service

import (
	"context"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
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
		productResponses = append(productResponses, &response.ProductResponse{
			ID:                  product.ID,
			Code:                product.Code,
			Name:                product.Name,
			UnitMassAcronym:     product.UnitMassAcronym,
			UnitMassDescription: product.UnitMassDescription,
			CreatedAt:           product.CreatedAt.String(),
			UpdatedAt:           product.UpdatedAt.String(),
		})
	}

	return productResponses, nil
}

func (repository *ProductService) FindByCode(ctx context.Context, code string) (*response.ProductResponse, error) {
	product, err := repository.ProductRepository.FindByCode(ctx, code)
	if err != nil {
		return nil, err
	}

	var productQualities []*response.ProductQualityResponse
	for _, productQuality := range product.ProductQualities {
		productQualities = append(productQualities, &response.ProductQualityResponse{
			ID:          productQuality.ID,
			ProductCode: productQuality.ProductCode,
			Quality:     productQuality.Quality,
			Price:       productQuality.Price,
			Quantity:    productQuality.Quantity,
			Type:        productQuality.Type,
		})
	}

	return &response.ProductResponse{
		ID:                  product.ID,
		Code:                product.Code,
		Name:                product.Name,
		UnitMassAcronym:     product.UnitMassAcronym,
		UnitMassDescription: product.UnitMassDescription,
		CreatedAt:           product.CreatedAt.String(),
		UpdatedAt:           product.UpdatedAt.String(),
		ProductQualities:    productQualities,
	}, nil
}

func (repository *ProductService) Create(ctx context.Context, request *request.CreateProductRequest) (*response.ProductResponse, error) {
	var productQualities []*model.ProductQuality
	for _, productQuality := range request.ProductQualities {
		productQualities = append(productQualities, &model.ProductQuality{
			Quality:  productQuality.Quality,
			Price:    productQuality.Price,
			Quantity: productQuality.Quantity,
			Type:     productQuality.Type,
		})
	}

	var productModel model.Product
	productModel.Name = request.Name
	productModel.UnitMassAcronym = request.UnitMassAcronym
	productModel.UnitMassDescription = request.UnitMassDescription
	productModel.ProductQualities = productQualities
	product, err := repository.ProductRepository.Create(ctx, &productModel)
	if err != nil {
		return nil, err
	}

	return &response.ProductResponse{
		ID:                  product.ID,
		Code:                product.Code,
		Name:                product.Name,
		UnitMassAcronym:     product.UnitMassAcronym,
		UnitMassDescription: product.UnitMassDescription,
		CreatedAt:           product.CreatedAt.String(),
		UpdatedAt:           product.UpdatedAt.String(),
	}, nil
}

func (repository *ProductService) Update(ctx context.Context, request *request.UpdateProductRequest) (*response.ProductResponse, error) {
	checkProduct, err := repository.ProductRepository.FindByCode(ctx, request.Code)
	if err != nil {
		return nil, err
	}

	var productQualities []*model.ProductQuality
	for _, productQuality := range request.ProductQualities {
		productQualities = append(productQualities, &model.ProductQuality{
			ID:       productQuality.ID,
			Quality:  productQuality.Quality,
			Price:    productQuality.Price,
			Quantity: productQuality.Quantity,
			Type:     productQuality.Type,
		})
	}

	checkProduct.Name = request.Name
	checkProduct.UnitMassAcronym = request.UnitMassAcronym
	checkProduct.UnitMassDescription = request.UnitMassDescription
	checkProduct.ProductQualities = productQualities
	product, err := repository.ProductRepository.Update(ctx, checkProduct)
	if err != nil {
		return nil, err
	}

	return &response.ProductResponse{
		ID:                  product.ID,
		Code:                product.Code,
		Name:                product.Name,
		UnitMassAcronym:     product.UnitMassAcronym,
		UnitMassDescription: product.UnitMassDescription,
		CreatedAt:           product.CreatedAt.String(),
		UpdatedAt:           product.UpdatedAt.String(),
	}, nil
}

func (repository *ProductService) Delete(ctx context.Context, code string) error {
	checkProduct, err := repository.ProductRepository.FindByCode(ctx, code)
	if err != nil {
		return err
	}

	err = repository.ProductRepository.Delete(ctx, checkProduct.Code)
	if err != nil {
		return err
	}

	return nil
}
