package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/service"
)

type ProductQualityController struct {
	ProductQualityService service.ProductQualityServiceContract
}

func NewProductQualityController(productQualityService service.ProductQualityServiceContract) *ProductQualityController {
	return &ProductQualityController{
		ProductQualityService: productQualityService,
	}
}

func (controller *ProductQualityController) FindAll(ctx *fiber.Ctx) error {
	productQualities, err := controller.ProductQualityService.FindAll(ctx.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", productQualities)
}

func (controller *ProductQualityController) FindAllByProductCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	productQualities, err := controller.ProductQualityService.FindAllByProductCode(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", productQualities)
}

func (controller *ProductQualityController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = controller.ProductQualityService.Delete(ctx.UserContext(), int64(id))
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "deleted", nil)
}
