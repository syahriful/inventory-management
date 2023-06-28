package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
	"net/http"
)

type TransactionController struct {
	TransactionService service.TransactionServiceContract
}

func NewTransactionController(transactionService service.TransactionServiceContract, route fiber.Router) TransactionController {
	controller := TransactionController{
		TransactionService: transactionService,
	}

	transaction := route.Group("/transactions")
	{
		transaction.Get("/", controller.FindAll)
		transaction.Get("/:code", controller.FindByCode)
		transaction.Post("/", controller.Create)
		transaction.Delete("/:code", controller.Delete)
		transaction.Patch("/:code", controller.Update)
		transaction.Get("/:code/supplier", controller.FindAllSupplierCode)
		transaction.Get("/:code/customer", controller.FindAllCustomerCode)
		transaction.Post("/transfer", controller.TransferStock)
	}

	return controller
}

func (controller *TransactionController) FindAll(ctx *fiber.Ctx) error {
	transactions, err := controller.TransactionService.FindAll(ctx.Context())
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", transactions)
}

func (controller *TransactionController) FindAllSupplierCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	transactions, err := controller.TransactionService.FindAllBySupplierCode(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", transactions)
}

func (controller *TransactionController) FindAllCustomerCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	transactions, err := controller.TransactionService.FindAllByCustomerCode(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", transactions)
}

func (controller *TransactionController) FindByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	transaction, err := controller.TransactionService.FindByCode(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", transaction)
}

func (controller *TransactionController) Create(ctx *fiber.Ctx) error {
	var transactionRequest request.CreateTransactionRequest
	err := ctx.BodyParser(&transactionRequest)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	errValidation := util.ValidateStruct(transactionRequest)
	if errValidation != nil {
		return response.ReturnErrorValidation(ctx, errValidation)
	}

	transaction, err := controller.TransactionService.Create(ctx.UserContext(), &transactionRequest)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusCreated, "created", transaction)
}

func (controller *TransactionController) Update(ctx *fiber.Ctx) error {
	var transactionRequest request.UpdateTransactionRequest
	err := ctx.BodyParser(&transactionRequest)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	errValidation := util.ValidateStruct(transactionRequest)
	if errValidation != nil {
		return response.ReturnErrorValidation(ctx, errValidation)
	}

	code := ctx.Params("code")
	transactionRequest.Code = code
	transaction, err := controller.TransactionService.Update(ctx.UserContext(), &transactionRequest)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		if err.Error() == response.ErrorUpdateTransactionTypeTransfer {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "updated", transaction)
}

func (controller *TransactionController) Delete(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	err := controller.TransactionService.Delete(ctx.Context(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "deleted", nil)
}

func (controller *TransactionController) TransferStock(ctx *fiber.Ctx) error {
	var transferStockRequest request.TransferStockTransactionRequest
	err := ctx.BodyParser(&transferStockRequest)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	errValidation := util.ValidateStruct(transferStockRequest)
	if errValidation != nil {
		return response.ReturnErrorValidation(ctx, errValidation)
	}

	if transferStockRequest.ProductQualityIDTransferred == transferStockRequest.ProductQualityID {
		return fiber.NewError(http.StatusBadRequest, "product quality id transferred and received cannot be the same")
	}

	transaction, err := controller.TransactionService.TransferStock(ctx.UserContext(), &transferStockRequest)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		if err.Error() == response.ErrorTransferStockDifferentProduct || err.Error() == response.ErrorStockNotEnough {
			return fiber.NewError(http.StatusBadRequest, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusCreated, "transferred", transaction)
}
