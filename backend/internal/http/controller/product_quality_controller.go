package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/service"
)

type ProductQualityController struct {
	ProductQualityService service.ProductQualityServiceContract
}

func NewProductQualityController(productQualityService service.ProductQualityServiceContract, route fiber.Router) ProductQualityController {
	controller := ProductQualityController{
		ProductQualityService: productQualityService,
	}

	productQuality := route.Group("/product-qualities")
	{
		productQuality.Get("/", controller.FindAll)
		productQuality.Get("/:id", controller.FindByID)
		productQuality.Get("/:code/product", controller.FindAllByProductCode)
		productQuality.Delete("/:id", controller.Delete)
	}

	return controller
}

func (controller *ProductQualityController) FindAll(ctx *fiber.Ctx) error {
	productQualities, err := controller.ProductQualityService.FindAll(ctx.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", productQualities).Build()
}

func (controller *ProductQualityController) FindAllByProductCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	productQualities, err := controller.ProductQualityService.FindAllByProductCode(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", productQualities).Build()
}

func (controller *ProductQualityController) FindByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	productQuality, err := controller.ProductQualityService.FindByID(ctx.UserContext(), int64(id))
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", productQuality).Build()
}

func (controller *ProductQualityController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = controller.ProductQualityService.Delete(ctx.UserContext(), int64(id))
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "deleted", nil).Build()
}
