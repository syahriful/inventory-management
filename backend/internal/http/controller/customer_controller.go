package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
	"net/http"
)

type CustomerController struct {
	CustomerService service.CustomerServiceContract
}

func NewCustomerController(customerService service.CustomerServiceContract) *CustomerController {
	return &CustomerController{
		CustomerService: customerService,
	}
}

func (controller *CustomerController) FindAll(ctx *fiber.Ctx) error {
	customers, err := controller.CustomerService.FindAll(ctx.Context())
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", customers)
}

func (controller *CustomerController) FindByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	customer, err := controller.CustomerService.FindByCode(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", customer)
}

func (controller *CustomerController) Create(ctx *fiber.Ctx) error {
	var customerRequest request.CreateCustomerRequest
	if err := ctx.BodyParser(&customerRequest); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(customerRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	customer, err := controller.CustomerService.Create(ctx.Context(), &customerRequest)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusCreated, "created", customer)
}

func (controller *CustomerController) Update(ctx *fiber.Ctx) error {
	var customerRequest request.UpdateCustomerRequest
	if err := ctx.BodyParser(&customerRequest); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(customerRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	code := ctx.Params("code")
	customerRequest.Code = code
	customer, err := controller.CustomerService.Update(ctx.Context(), &customerRequest)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "updated", customer)
}

func (controller *CustomerController) Delete(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	err := controller.CustomerService.Delete(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.NotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "deleted", nil)
}
