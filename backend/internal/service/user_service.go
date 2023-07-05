package service

import (
	"bytes"
	"context"
	"errors"
	"inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/internal/repository"
	third_party "inventory-management/backend/internal/third_party/elasticsearch"
)

type UserService struct {
	UserRepository     repository.UserRepositoryContract
	Elasticsearch      third_party.ElasticsearchContract
	IndexElasticsearch string
}

func NewUserService(userRepository repository.UserRepositoryContract, elasticsearch third_party.ElasticsearchContract) UserServiceContract {
	return &UserService{
		UserRepository:     userRepository,
		Elasticsearch:      elasticsearch,
		IndexElasticsearch: "users",
	}
}

func (service *UserService) Search(ctx context.Context, data bytes.Buffer, offset int, limit int, totalRecord chan<- int64) (map[string]interface{}, error) {
	totalRecords, err := service.Elasticsearch.CountAll(ctx, service.IndexElasticsearch, data)
	if err != nil {
		return nil, err
	}

	totalRecord <- totalRecords
	searchResponse, err := service.Elasticsearch.Search(ctx, service.IndexElasticsearch, data, offset, limit)
	if err != nil {
		return nil, err
	}

	return searchResponse, nil
}

func (service *UserService) FindAll(ctx context.Context, offset int, limit int) ([]*response.UserResponse, error) {
	users, err := service.UserRepository.FindAll(ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	var userResponses []*response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	return userResponses, nil
}

func (service *UserService) CountAll(ctx context.Context) (int64, error) {
	count, err := service.UserRepository.CountAll(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (service *UserService) FindByID(ctx context.Context, id int64) (*response.UserResponse, error) {
	user, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

func (service *UserService) VerifyLogin(ctx context.Context, request *request.LoginUserRequest) (*response.UserLoginResponse, error) {
	user, err := service.UserRepository.FindByUsername(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	err = user.VerifyPassword(request.Password)
	if err != nil {
		return nil, err
	}

	token, err := user.GenerateTokenJWT()
	if err != nil {
		return nil, err
	}

	return &response.UserLoginResponse{
		Token: token,
	}, nil
}

func (service *UserService) Create(ctx context.Context, request *request.CreateUserRequest) (*response.UserResponse, error) {
	_, err := service.UserRepository.FindByUsername(ctx, request.Username)
	if err == nil {
		return nil, errors.New(response.ErrorUsernameExists)
	}

	var userRequest model.User
	userRequest.Name = request.Name
	userRequest.Username = request.Username
	userRequest.Password = request.Password

	user, err := service.UserRepository.Create(ctx, &userRequest)
	if err != nil {
		return nil, err
	}

	// Insert to Elasticsearch
	err = service.Elasticsearch.Create(ctx, service.IndexElasticsearch, user.ToResponse(), user.ID)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

func (service *UserService) Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error) {
	checkUser, err := service.UserRepository.FindByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	newPassword := checkUser.Password
	if request.Password != "" {
		passwordHashed, err := checkUser.HashPassword(request.Password)
		if err != nil {
			return nil, err
		}
		newPassword = passwordHashed
	}

	checkUser.ID = request.ID
	checkUser.Name = request.Name
	checkUser.Password = newPassword
	user, err := service.UserRepository.Update(ctx, checkUser)
	if err != nil {
		return nil, err
	}

	// Update to Elasticsearch
	err = service.Elasticsearch.Update(ctx, service.IndexElasticsearch, user.ToResponse(), user.ID)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

func (service *UserService) Delete(ctx context.Context, id int64) error {
	_, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	err = service.UserRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	// Delete from Elasticsearch
	err = service.Elasticsearch.Delete(ctx, service.IndexElasticsearch, id)
	if err != nil {
		return err
	}

	return nil
}
