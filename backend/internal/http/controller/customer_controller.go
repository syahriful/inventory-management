package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
	"net/http"
)

type CustomerController struct {
	CustomerService service.CustomerServiceContract
}

func NewCustomerController(customerService service.CustomerServiceContract, route fiber.Router) CustomerController {
	controller := CustomerController{
		CustomerService: customerService,
	}

	customer := route.Group("/customers")
	{
		customer.Get("/", controller.FindAll)
		customer.Get("/:code", controller.FindByCode)
		customer.Post("/", controller.Create)
		customer.Patch("/:code", controller.Update)
		customer.Delete("/:code", controller.Delete)
	}

	return controller
}

func (controller *CustomerController) FindAll(ctx *fiber.Ctx) error {
	totalRecords, err := controller.CustomerService.CountAll(ctx.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	pagination, offset := util.CreatePagination(ctx, totalRecords)
	customers, err := controller.CustomerService.FindAll(ctx.UserContext(), offset, pagination.Limit)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", customers).WithPagination(&pagination).Build()
}

func (controller *CustomerController) FindByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	customer, err := controller.CustomerService.FindByCode(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", customer).Build()
}

func (controller *CustomerController) Create(ctx *fiber.Ctx) error {
	var customerRequest request.CreateCustomerRequest
	if err := ctx.BodyParser(&customerRequest); err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(customerRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	customer, err := controller.CustomerService.Create(ctx.UserContext(), &customerRequest)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusCreated, "created", customer).Build()
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
	customer, err := controller.CustomerService.Update(ctx.UserContext(), &customerRequest)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "updated", customer).Build()
}

func (controller *CustomerController) Delete(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	err := controller.CustomerService.Delete(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "deleted", nil).Build()
}
