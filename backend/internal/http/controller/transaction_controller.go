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
		transaction.Get("/:code/supplier", controller.FindAllBySupplierCode)
		transaction.Get("/:code/customer", controller.FindAllByCustomerCode)
		transaction.Post("/transfer", controller.TransferStock)
	}

	return controller
}

func (controller *TransactionController) FindAll(ctx *fiber.Ctx) error {
	currPage := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	totalRecords, err := controller.TransactionService.CountAll(ctx.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	pagination := util.CreatePagination(currPage, limit, totalRecords)
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	transactions, err := controller.TransactionService.FindAll(ctx.UserContext(), offset, pagination.Limit)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", transactions).WithPagination(&pagination).Build()
}

func (controller *TransactionController) FindAllBySupplierCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	transactions, err := controller.TransactionService.FindAllBySupplierCode(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", transactions).Build()
}

func (controller *TransactionController) FindAllByCustomerCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	transactions, err := controller.TransactionService.FindAllByCustomerCode(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", transactions).Build()
}

func (controller *TransactionController) FindByCode(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	transaction, err := controller.TransactionService.FindByCode(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", transaction).Build()
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
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusCreated, "created", transaction).Build()
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

	return response.ReturnJSON(ctx, http.StatusOK, "updated", transaction).Build()
}

func (controller *TransactionController) Delete(ctx *fiber.Ctx) error {
	code := ctx.Params("code")
	err := controller.TransactionService.Delete(ctx.UserContext(), code)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(http.StatusNotFound, err.Error())
		}
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "deleted", nil).Build()
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

	return response.ReturnJSON(ctx, http.StatusCreated, "transferred", transaction).Build()
}
