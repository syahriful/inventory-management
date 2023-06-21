package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
)

type ProductController struct {
	ProductService service.ProductServiceContract
}

func NewProductController(productService service.ProductServiceContract) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (controller *ProductController) FindAll(ctx *fiber.Ctx) error {
	products, err := controller.ProductService.FindAll(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusOK, "OK", products)
}

func (controller *ProductController) FindByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")

	product, err := controller.ProductService.FindByCode(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusOK, "OK", product)
}

func (controller *ProductController) Create(ctx *fiber.Ctx) error {
	var productRequest request.CreateProductRequest
	if err := ctx.BodyParser(&productRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(productRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	product, err := controller.ProductService.Create(ctx.Context(), &productRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusCreated, "created", product)
}

func (controller *ProductController) Update(ctx *fiber.Ctx) error {
	var productRequest request.UpdateProductRequest
	if err := ctx.BodyParser(&productRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(productRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	code := ctx.Params("code")
	productRequest.Code = code
	product, err := controller.ProductService.Update(ctx.Context(), &productRequest)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusOK, "updated", product)
}

func (controller *ProductController) Delete(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	err := controller.ProductService.Delete(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusOK, "deleted", nil)
}
