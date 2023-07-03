package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
)

type ProductController struct {
	ProductService service.ProductServiceContract
}

func NewProductController(productService service.ProductServiceContract, route fiber.Router) ProductController {
	controller := ProductController{
		ProductService: productService,
	}

	product := route.Group("/products")
	{
		product.Get("/", controller.FindAll)
		product.Get("/:code", controller.FindByCode)
		product.Post("/", controller.Create)
		product.Patch("/:code", controller.Update)
		product.Delete("/:code", controller.Delete)
	}

	return controller
}

func (controller *ProductController) FindAll(ctx *fiber.Ctx) error {
	currPage := ctx.QueryInt("page", 1)
	if currPage <= 0 {
		currPage = 1
	}
	limit := ctx.QueryInt("limit", 10)

	totalRecords, err := controller.ProductService.CountAll(ctx.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	pagination := util.CreatePagination(currPage, limit, totalRecords)
	offset := (currPage - 1) * limit
	products, err := controller.ProductService.FindAll(ctx.UserContext(), offset, limit)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", products).WithPagination(&pagination).Build()
}

func (controller *ProductController) FindByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")

	product, err := controller.ProductService.FindByCode(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", product).Build()
}

func (controller *ProductController) Create(ctx *fiber.Ctx) error {
	var productRequest request.CreateProductRequest
	if err := ctx.BodyParser(&productRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(productRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	product, err := controller.ProductService.Create(ctx.UserContext(), &productRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusCreated, "created", product).Build()
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
	product, err := controller.ProductService.Update(ctx.UserContext(), &productRequest)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "updated", product).Build()
}

func (controller *ProductController) Delete(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	err := controller.ProductService.Delete(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "deleted", nil).Build()
}
