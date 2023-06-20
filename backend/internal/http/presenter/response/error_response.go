package response

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

var (
	NotFound = "record not found"
)

func ReturnErrorValidation(c *fiber.Ctx, err []*ErrorResponse) error {
	return c.Status(http.StatusBadRequest).JSON(ErrorValidationResponse{
		Code:   http.StatusBadRequest,
		Status: "There are inaccuracies in the validation process",
		Error:  err,
	})
}

func ReturnError(c *fiber.Ctx, code int, err error) error {
	return c.Status(code).JSON(ApiResponse{
		Code:   code,
		Status: err.Error(),
		Data:   nil,
	})
}
