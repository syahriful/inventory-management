package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"inventory-management/backend/internal/http/presenter/request"
	"inventory-management/backend/internal/http/presenter/response"
	"inventory-management/backend/internal/model"
	"inventory-management/backend/internal/repository"
	"testing"
)

func TestUserService_FindAll(t *testing.T) {
	testCases := []struct {
		name              string
		expectedUserRepo  []*model.User
		expectedUserSvc   []*response.UserResponse
		expectedRepoError error
		expectedSvcError  error
	}{
		{
			name: "Number of users more than 1",
			expectedUserRepo: []*model.User{
				{
					ID:       1,
					Name:     "Widdy Arfiansyah",
					Username: "wdyarfn",
					Password: "1234567",
				},
				{
					ID:       2,
					Name:     "Arfiansyah",
					Username: "arfn",
					Password: "1234567",
				},
			},
			expectedUserSvc: []*response.UserResponse{
				{
					ID:        1,
					Name:      "Widdy Arfiansyah",
					Username:  "wdyarfn",
					CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
				},
				{
					ID:        2,
					Name:      "Arfiansyah",
					Username:  "arfn",
					CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
					UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
				},
			},
			expectedRepoError: nil,
			expectedSvcError:  nil,
		},
		{
			name:              "Number of users is 0 or null",
			expectedUserRepo:  nil,
			expectedUserSvc:   nil,
			expectedRepoError: nil,
			expectedSvcError:  nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var repo repository.UserRepositoryMock
			repo.On("FindAll").Return(tc.expectedUserRepo, tc.expectedRepoError)
			svc := NewUserService(&repo)
			ctx := context.Background()
			result, err := svc.FindAll(ctx)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedUserSvc, result)
			assert.Equal(t, len(tc.expectedUserSvc), len(result))
		})
	}
}

func TestUserService_FindByID(t *testing.T) {
	testCases := []struct {
		name              string
		requestID         int64
		expectedUserRepo  *model.User
		expectedUserSvc   *response.UserResponse
		expectedRepoError error
		expectedSvcError  error
	}{
		{
			name:      "User exists with given ID",
			requestID: 1,
			expectedUserRepo: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: "1234567",
			},
			expectedUserSvc: &response.UserResponse{
				ID:        1,
				Name:      "Widdy Arfiansyah",
				Username:  "wdyarfn",
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedRepoError: nil,
			expectedSvcError:  nil,
		},
		{
			name:              "User doesnt exists with given ID",
			requestID:         1,
			expectedUserRepo:  nil,
			expectedUserSvc:   nil,
			expectedRepoError: errors.New(response.NotFound),
			expectedSvcError:  errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var repo repository.UserRepositoryMock
			repo.On("FindByID", tc.requestID).Return(tc.expectedUserRepo, tc.expectedRepoError)
			svc := NewUserService(&repo)
			ctx := context.Background()
			result, err := svc.FindByID(ctx, tc.requestID)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedUserSvc, result)
		})
	}
}

func TestUserService_VerifyLogin(t *testing.T) {
	password, _ := bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
	testCases := []struct {
		name              string
		request           *request.LoginUserRequest
		expectedUserRepo  *model.User
		expectedUserSvc   *response.UserLoginResponse
		expectedRepoError error
		expectedSvcError  error
	}{
		{
			name: "Verify login with required fields",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
				Password: "1234567",
			},
			expectedUserRepo: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedUserSvc: &response.UserLoginResponse{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc3NjUxMDQsInVzZXJuYW1lIjoiYXJmaWFuIn0.xchUDFf3xBJlfQWUzq4FCBjAHcYSAEcAiruZuuGAJBk",
			},
			expectedRepoError: nil,
			expectedSvcError:  nil,
		},
		{
			name: "User doesnt exists with given username",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
				Password: "1234567",
			},
			expectedUserRepo:  nil,
			expectedUserSvc:   nil,
			expectedRepoError: errors.New(response.NotFound),
			expectedSvcError:  errors.New(response.NotFound),
		},
		{
			name: "Password doesnt match with given username",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
				Password: "12345678910",
			},
			expectedUserRepo: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedUserSvc:   nil,
			expectedRepoError: nil,
			expectedSvcError:  errors.New(response.InvalidPassword),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var repo repository.UserRepositoryMock
			repo.On("FindByUsername", tc.request.Username).Return(tc.expectedUserRepo, tc.expectedRepoError)
			svc := NewUserService(&repo)
			ctx := context.Background()
			result, err := svc.VerifyLogin(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			if tc.expectedUserSvc != nil {
				assert.Greater(t, len(tc.expectedUserSvc.Token), 100)
				assert.Greater(t, len(result.Token), 100)
			} else {
				assert.Nil(t, result)
			}
		})
	}
}

func TestUserService_Create(t *testing.T) {
	password, _ := bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
	testCases := []struct {
		name                            string
		request                         *request.CreateUserRequest
		requestRepo                     *model.User
		expectedUserRepo                *model.User
		expectedUserSvc                 *response.UserResponse
		expectedUserCreateRepo          *model.User
		expectedRepoError               error
		expectedSvcError                error
		expectedRepoFindByUsernameError error
	}{
		{
			name: "Create user with required fields",
			request: &request.CreateUserRequest{
				Name:     "Widdy Arfian",
				Username: "wdyarfn",
				Password: "1234567",
			},
			requestRepo: &model.User{
				Name:     "Widdy Arfian",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedUserRepo: nil,
			expectedUserCreateRepo: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedUserSvc: &response.UserResponse{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
			},
			expectedRepoError:               nil,
			expectedSvcError:                nil,
			expectedRepoFindByUsernameError: errors.New(response.NotFound),
		},
		{
			name: "Create user with given the exists username",
			request: &request.CreateUserRequest{
				Name:     "Widdy Arfian",
				Username: "wdyarfn",
				Password: "1234567",
			},
			requestRepo: &model.User{
				Name:     "Widdy Arfian",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedUserRepo: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedUserCreateRepo:          nil,
			expectedUserSvc:                 nil,
			expectedRepoError:               errors.New("Whoops"),
			expectedSvcError:                errors.New(response.UsernameExists),
			expectedRepoFindByUsernameError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var repo repository.UserRepositoryMock
			repo.On("FindByUsername", tc.requestRepo.Username).Return(tc.expectedUserRepo, tc.expectedRepoFindByUsernameError)
			repo.On("Create", mock.Anything).Return(tc.expectedUserCreateRepo, tc.expectedRepoError)
			svc := NewUserService(&repo)
			ctx := context.Background()
			result, err := svc.Create(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			if err == nil {
				assert.Equal(t, tc.expectedUserSvc.Name, result.Name)
			}
		})
	}
}
