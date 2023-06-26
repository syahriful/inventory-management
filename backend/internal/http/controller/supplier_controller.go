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

func NewSupplierController(supplierService service.SupplierServiceContract, route fiber.Router) SupplierController {
	controller := SupplierController{
		SupplierService: supplierService,
	}

	supplier := route.Group("/suppliers")
	{
		supplier.Get("/", controller.FindAll)
		supplier.Get("/:code", controller.FindByCode)
		supplier.Post("/", controller.Create)
		supplier.Patch("/:code", controller.Update)
		supplier.Delete("/:code", controller.Delete)
	}

	return controller
}

func (controller *SupplierController) FindAll(ctx *fiber.Ctx) error {
	suppliers, err := controller.SupplierService.FindAll(ctx.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", suppliers)
}

func (controller *SupplierController) FindByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	supplier, err := controller.SupplierService.FindByCode(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", supplier)
}

func (controller *SupplierController) Create(ctx *fiber.Ctx) error {
	var supplierRequest request.CreateSupplierRequest
	if err := ctx.BodyParser(&supplierRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(supplierRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	supplier, err := controller.SupplierService.Create(ctx.UserContext(), &supplierRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusCreated, "created", supplier)
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
	supplier, err := controller.SupplierService.Update(ctx.UserContext(), &supplierRequest)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "updated", supplier)
}

func (controller *SupplierController) Delete(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	err := controller.SupplierService.Delete(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "deleted", nil)
}
