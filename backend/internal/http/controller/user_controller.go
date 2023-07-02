package controller

import (
	"bytes"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gofiber/fiber/v2"
	"inventory-management/backend/internal/http/request"
	"inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/service"
	"inventory-management/backend/util"
	"strconv"
)

type UserController struct {
	UserService   service.UserServiceContract
	Elasticsearch *elasticsearch.Client
}

func NewUserController(userService service.UserServiceContract, elasticsearch *elasticsearch.Client, route fiber.Router) UserController {
	controller := UserController{
		UserService:   userService,
		Elasticsearch: elasticsearch,
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
	limit := ctx.QueryInt("limit", 10)

	count, err := controller.Elasticsearch.Count(
		controller.Elasticsearch.Count.WithContext(ctx.UserContext()),
		controller.Elasticsearch.Count.WithIndex("users"),
		controller.Elasticsearch.Count.WithBody(&data),
		controller.Elasticsearch.Count.WithPretty(),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer count.Body.Close()

	var totalHits map[string]interface{}
	if err := json.NewDecoder(count.Body).Decode(&totalHits); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var totalHitsInt int64
	if totalHits["count"] != nil {
		totalHitsInt = int64(totalHits["count"].(float64))
	}
	pagination := util.CreatePagination(currPage, limit, totalHitsInt)
	offset := (currPage - 1) * limit

	if err := json.NewEncoder(&data).Encode(query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	res, err := controller.Elasticsearch.Search(
		controller.Elasticsearch.Search.WithContext(ctx.UserContext()),
		controller.Elasticsearch.Search.WithIndex("users"),
		controller.Elasticsearch.Search.WithBody(&data),
		controller.Elasticsearch.Search.WithTrackTotalHits(true),
		controller.Elasticsearch.Search.WithPretty(),
		controller.Elasticsearch.Search.WithSize(limit),
		controller.Elasticsearch.Search.WithFrom(offset),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer res.Body.Close()

	if res.IsError() {
		return fiber.NewError(fiber.StatusBadRequest, res.String())
	}

	var searchResponse map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&searchResponse); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return response.ReturnJSON(ctx, fiber.StatusOK, "OK", searchResponse).WithPagination(&pagination).Build()
}

func (controller *UserController) FindAll(ctx *fiber.Ctx) error {
	currPage := ctx.QueryInt("page", 1)
	limit := ctx.QueryInt("limit", 10)

	totalRecords, err := controller.UserService.CountAll(ctx.UserContext())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	pagination := util.CreatePagination(currPage, limit, totalRecords)
	offset := (pagination.CurrentPage - 1) * pagination.Limit
	users, err := controller.UserService.FindAll(ctx.UserContext(), offset, pagination.Limit)
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

	// Insert to Elasticsearch
	data, err := json.Marshal(user)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	reqEs := esapi.IndexRequest{
		Index:      "users",
		DocumentID: strconv.FormatInt(user.ID, 10),
		Body:       bytes.NewReader(data),
	}

	res, err := reqEs.Do(ctx.UserContext(), controller.Elasticsearch)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer res.Body.Close()

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

	// Delete from Elasticsearch
	idString := strconv.FormatInt(int64(id), 10)
	res, err := controller.Elasticsearch.Delete("users", idString,
		controller.Elasticsearch.Delete.WithContext(ctx.UserContext()),
	)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer res.Body.Close()

	return response.ReturnJSON(ctx, fiber.StatusOK, "deleted", nil).Build()
}
