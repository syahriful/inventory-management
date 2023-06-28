package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var (
	ErrorNotFound                      = "record not found"
	ErrorInvalidPassword               = "invalid password"
	ErrorUsernameExists                = "username already exist"
	ErrorValidation                    = "there are inaccuracies in the validation process"
	ErrorTransferStockDifferentProduct = "transfer stock must be the same product"
	ErrorStockNotEnough                = "stock is not enough"
	ErrorUpdateTransactionTypeTransfer = "transaction type cannot be changed in transfer process"
)

type ApiResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResponse struct {
	FailedField string `json:"failed_field,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Value       string `json:"value,omitempty"`
}

type ErrorValidationResponse struct {
	Code   int              `json:"code"`
	Status string           `json:"status"`
	Error  []*ErrorResponse `json:"error"`
}

func ReturnErrorValidation(c *fiber.Ctx, err []*ErrorResponse) error {
	return c.Status(http.StatusBadRequest).JSON(ErrorValidationResponse{
		Code:   http.StatusBadRequest,
		Status: ErrorValidation,
		Error:  err,
	})
}

func ReturnJSON(c *fiber.Ctx, code int, status string, data interface{}) error {
	return c.Status(code).JSON(ApiResponse{
		Code:   code,
		Status: status,
		Data:   data,
	})
}
