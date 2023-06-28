package service

import (
	"context"
	"errors"
	"inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/internal/repository"
)

type UserService struct {
	UserRepository repository.UserRepositoryContract
}

func NewUserService(userRepository repository.UserRepositoryContract) UserServiceContract {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (service *UserService) FindAll(ctx context.Context) ([]*response.UserResponse, error) {
	users, err := service.UserRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var userResponses []*response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	return userResponses, nil
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

	return user.ToResponse(), nil
}

func (service *UserService) Delete(ctx context.Context, id int64) error {
	_, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return service.UserRepository.Delete(ctx, id)
}
