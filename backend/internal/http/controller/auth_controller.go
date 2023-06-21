package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
	"net/http"
)

type AuthController struct {
	UserService service.UserServiceContract
}

func NewAuthController(userService service.UserServiceContract) *AuthController {
	return &AuthController{
		UserService: userService,
	}
}

func (controller *AuthController) Login(ctx *fiber.Ctx) error {
	var loginUserRequest request.LoginUserRequest
	if err := ctx.BodyParser(&loginUserRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(loginUserRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	userResponse, err := controller.UserService.VerifyLogin(ctx.Context(), &loginUserRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return response.ReturnSuccess(ctx, http.StatusOK, "OK", userResponse)
}

func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	var userRequest request.CreateUserRequest
	if err := ctx.BodyParser(&userRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(userRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	user, err := controller.UserService.Create(ctx.Context(), &userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnSuccess(ctx, fiber.StatusCreated, "created", user)
}
