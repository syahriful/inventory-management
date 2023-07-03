package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
	"sync"
)

type UserController struct {
	UserService service.UserServiceContract
}

func NewUserController(userService service.UserServiceContract, route fiber.Router) UserController {
	controller := UserController{
		UserService: userService,
	}

	user := route.Group("/users")
	{
		user.Post("/search", controller.Search)
		user.Get("/", controller.FindAll)
		user.Get("/:id", controller.FindByID)
		user.Post("/", controller.Create)
		user.Patch("/:id", controller.Update)
		user.Delete("/:id", controller.Delete)
	}

	return controller
}

func (controller *UserController) Search(ctx *fiber.Ctx) error {
	var data bytes.Buffer
	var query map[string]interface{}
	err := ctx.BodyParser(&query)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := json.NewEncoder(&data).Encode(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	currPage := ctx.QueryInt("page", 1)
	if currPage <= 0 {
		currPage = 1
	}
	limit := ctx.QueryInt("limit", 10)

	var wg sync.WaitGroup
	var pagination response.Pagination
	totalRecords := make(chan int64)

	wg.Add(2)
	go func() {
		defer wg.Done()
		pagination = util.CreatePagination(currPage, limit, <-totalRecords)
	}()

	var searchResponse map[string]interface{}
	go func() {
		defer wg.Done()
		offset := (currPage - 1) * limit
		searchResponse, err = controller.UserService.Search(ctx.UserContext(), data, offset, limit, totalRecords)
	}()
	wg.Wait()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", searchResponse).WithPagination(&pagination).Build()
}

func (controller *UserController) FindAll(ctx *fiber.Ctx) error {
	currPage := ctx.QueryInt("page", 1)
	if currPage <= 0 {
		currPage = 1
	}
	limit := ctx.QueryInt("limit", 10)

	totalRecords, err := controller.UserService.CountAll(ctx.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	pagination := util.CreatePagination(currPage, limit, totalRecords)
	offset := (currPage - 1) * limit
	users, err := controller.UserService.FindAll(ctx.UserContext(), offset, limit)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", users).WithPagination(&pagination).Build()
}

func (controller *UserController) FindByID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := controller.UserService.FindByID(ctx.UserContext(), int64(id))
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", user).Build()
}

func (controller *UserController) Create(ctx *fiber.Ctx) error {
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

func (controller *UserController) Update(ctx *fiber.Ctx) error {
	var userRequest request.UpdateUserRequest
	if err := ctx.BodyParser(&userRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if errValidate := util.ValidateStruct(userRequest); errValidate != nil {
		return response.ReturnErrorValidation(ctx, errValidate)
	}

	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	userRequest.ID = int64(id)
	user, err := controller.UserService.Update(ctx.UserContext(), &userRequest)
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "updated", user).Build()
}

func (controller *UserController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = controller.UserService.Delete(ctx.UserContext(), int64(id))
	if err != nil {
		if err.Error() == response.ErrorNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "deleted", nil).Build()
}
