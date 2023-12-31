package controller

import (
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
	"net/http"
)

type AuthController struct {
	UserService service.UserServiceContract
}

func NewAuthController(userService service.UserServiceContract, route fiber.Router) AuthController {
	controller := AuthController{
		UserService: userService,
	}

	route.Post("/login", controller.Login)
	route.Post("/register", controller.Register)

	return controller
}

func (controller *AuthController) Login(ctx *fiber.Ctx) error {
	var loginUserRequest request.LoginUserRequest
	if err := ctx.BodyParser(&loginUserRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(loginUserRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	userResponse, err := controller.UserService.VerifyLogin(ctx.UserContext(), &loginUserRequest)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		if err.Error() == response.ErrorInvalidPassword {
			return fiber.NewError(fiber.StatusBadRequest, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, http.StatusOK, "OK", userResponse).Build()
}

func (controller *AuthController) Register(ctx *fiber.Ctx) error {
	var userRequest request.CreateUserRequest
	if err := ctx.BodyParser(&userRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(userRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	user, err := controller.UserService.Create(ctx.UserContext(), &userRequest)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusCreated, "created", user).Build()
}
