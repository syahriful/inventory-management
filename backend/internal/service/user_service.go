package service

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/internal/repository"
	"inventory-management/backend/util"
	"time"
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
		userResponses = append(userResponses, &response.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Username:  user.Username,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		})
	}

	return userResponses, nil
}

func (service *UserService) FindByID(ctx context.Context, id int64) (*response.UserResponse, error) {
	user, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (service *UserService) VerifyLogin(ctx context.Context, request *request.LoginUserRequest) (*response.UserLoginResponse, error) {
	user, err := service.UserRepository.FindByUsername(ctx, request.Username)
	if err != nil {
		return nil, err
	}

	err = util.VerifyPassword(user.Password, request.Password)
	if err != nil {
		return nil, err
	}

	myClaims := jwt.Claims(jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &response.UserLoginResponse{
		Token: token,
	}, nil
}

func (service *UserService) Create(ctx context.Context, request *request.CreateUserRequest) (*response.UserResponse, error) {
	passwordHashed, err := util.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	_, err = service.UserRepository.FindByUsername(ctx, request.Username)
	if err == nil {
		return nil, errors.New(response.UsernameExists)
	}

	user, err := service.UserRepository.Create(ctx, &model.User{
		Name:     request.Name,
		Username: request.Username,
		Password: passwordHashed,
	})
	if err != nil {
		return nil, err
	}

	return &response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (service *UserService) Update(ctx context.Context, request *request.UpdateUserRequest) (*response.UserResponse, error) {
	checkUser, err := service.UserRepository.FindByID(ctx, request.ID)
	if err != nil {
		return nil, err
	}

	newPassword := checkUser.Password
	if request.Password != "" {
		passwordHashed, err := util.HashPassword(request.Password)
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

	return &response.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}, nil
}

func (service *UserService) Delete(ctx context.Context, id int64) error {
	_, err := service.UserRepository.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return service.UserRepository.Delete(ctx, id)
}
