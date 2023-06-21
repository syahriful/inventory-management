package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
	"net/http"
)

type SupplierController struct {
	SupplierService service.SupplierServiceContract
}

func NewSupplierController(supplierService service.SupplierServiceContract) *SupplierController {
	return &SupplierController{
		SupplierService: supplierService,
	}
}

func (controller *SupplierController) FindAll(ctx *fiber.Ctx) error {
	suppliers, err := controller.SupplierService.FindAll(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, http.StatusOK, "OK", suppliers)
}

func (controller *SupplierController) FindByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	supplier, err := controller.SupplierService.FindByCode(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, http.StatusOK, "OK", supplier)
}

func (controller *SupplierController) Create(ctx *fiber.Ctx) error {
	var supplierRequest request.CreateSupplierRequest
	if err := ctx.BodyParser(&supplierRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(supplierRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	supplier, err := controller.SupplierService.Create(ctx.Context(), &supplierRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, http.StatusCreated, "created", supplier)
}

func (controller *SupplierController) Update(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	var supplierRequest request.UpdateSupplierRequest
	if err := ctx.BodyParser(&supplierRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(supplierRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	supplierRequest.Code = code
	supplier, err := controller.SupplierService.Update(ctx.Context(), &supplierRequest)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, http.StatusOK, "updated", supplier)
}

func (controller *SupplierController) Delete(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	err := controller.SupplierService.Delete(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, http.StatusOK, "deleted", nil)
}
