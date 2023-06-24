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
	repository "inventory-management/backend/internal/repository/mock"
	"testing"
)

func TestUserService_FindAll(t *testing.T) {
	testCases := []struct {
		name                         string
		expectedUserRepoFindAll      []*model.User
		expectedUserRepoFindAllError error
		expectedSvc                  []*response.UserResponse
		expectedSvcError             error
	}{
		{
			name: "Number of users more than 1",
			expectedUserRepoFindAll: []*model.User{
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
			expectedSvc: []*response.UserResponse{
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
			expectedUserRepoFindAllError: nil,
			expectedSvcError:             nil,
		},
		{
			name:                         "Number of users is 0 or null",
			expectedUserRepoFindAll:      nil,
			expectedSvc:                  nil,
			expectedUserRepoFindAllError: nil,
			expectedSvcError:             nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			repo.On("FindAll", ctx).Return(tc.expectedUserRepoFindAll, tc.expectedUserRepoFindAllError)
			svc := NewUserService(&repo)
			result, err := svc.FindAll(ctx)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
			assert.Equal(t, len(tc.expectedSvc), len(result))
		})
	}
}

func TestUserService_FindByID(t *testing.T) {
	testCases := []struct {
		name                          string
		request                       int64
		expectedUserRepoFindByID      *model.User
		expectedUserRepoFindByIDError error
		expectedSvc                   *response.UserResponse
		expectedSvcError              error
	}{
		{
			name:    "User exists with given ID",
			request: 1,
			expectedUserRepoFindByID: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: "1234567",
			},
			expectedSvc: &response.UserResponse{
				ID:        1,
				Name:      "Widdy Arfiansyah",
				Username:  "wdyarfn",
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedUserRepoFindByIDError: nil,
			expectedSvcError:              nil,
		},
		{
			name:                          "User doesnt exists with given ID",
			request:                       1,
			expectedUserRepoFindByID:      nil,
			expectedSvc:                   nil,
			expectedUserRepoFindByIDError: errors.New(response.NotFound),
			expectedSvcError:              errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			repo.On("FindByID", ctx, tc.request).Return(tc.expectedUserRepoFindByID, tc.expectedUserRepoFindByIDError)
			svc := NewUserService(&repo)
			result, err := svc.FindByID(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)
		})
	}
}

func TestUserService_VerifyLogin(t *testing.T) {
	password, _ := bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
	testCases := []struct {
		name                                string
		request                             *request.LoginUserRequest
		expectedUserRepoFindByUsername      *model.User
		expectedUserRepoFindByUsernameError error
		expectedSvc                         *response.UserLoginResponse
		expectedSvcError                    error
	}{
		{
			name: "Verify login with required fields",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
				Password: "1234567",
			},
			expectedUserRepoFindByUsername: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedSvc: &response.UserLoginResponse{
				Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODc3NjUxMDQsInVzZXJuYW1lIjoiYXJmaWFuIn0.xchUDFf3xBJlfQWUzq4FCBjAHcYSAEcAiruZuuGAJBk",
			},
			expectedUserRepoFindByUsernameError: nil,
			expectedSvcError:                    nil,
		},
		{
			name: "User doesnt exists with given username",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
				Password: "1234567",
			},
			expectedUserRepoFindByUsername:      nil,
			expectedSvc:                         nil,
			expectedUserRepoFindByUsernameError: errors.New(response.NotFound),
			expectedSvcError:                    errors.New(response.NotFound),
		},
		{
			name: "Password doesnt match with given password input",
			request: &request.LoginUserRequest{
				Username: "wdyarfn",
				Password: "12345678910",
			},
			expectedUserRepoFindByUsername: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedSvc:                         nil,
			expectedUserRepoFindByUsernameError: nil,
			expectedSvcError:                    errors.New(response.InvalidPassword),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			repo.On("FindByUsername", ctx, tc.request.Username).Return(tc.expectedUserRepoFindByUsername, tc.expectedUserRepoFindByUsernameError)
			svc := NewUserService(&repo)
			result, err := svc.VerifyLogin(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			if tc.expectedSvc != nil {
				assert.Greater(t, len(tc.expectedSvc.Token), 100)
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
		name                                string
		request                             *request.CreateUserRequest
		requestRepo                         *model.User
		expectedUserRepoFindByUsername      *model.User
		expectedUserRepoFindByUsernameError error
		expectedUserRepoCreate              *model.User
		expectedUserRepoCreateError         error
		expectedSvc                         *response.UserResponse
		expectedSvcError                    error
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
			expectedUserRepoFindByUsername: nil,
			expectedUserRepoCreate: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedSvc: &response.UserResponse{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
			},
			expectedUserRepoCreateError:         nil,
			expectedSvcError:                    nil,
			expectedUserRepoFindByUsernameError: errors.New(response.NotFound),
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
			expectedUserRepoFindByUsername: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Username: "wdyarfn",
				Password: string(password),
			},
			expectedUserRepoCreate:              nil,
			expectedSvc:                         nil,
			expectedUserRepoCreateError:         errors.New("getting an error"),
			expectedSvcError:                    errors.New(response.UsernameExists),
			expectedUserRepoFindByUsernameError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			repo.On("FindByUsername", ctx, tc.requestRepo.Username).Return(tc.expectedUserRepoFindByUsername, tc.expectedUserRepoFindByUsernameError)
			repo.On("Create", ctx, mock.Anything).Return(tc.expectedUserRepoCreate, tc.expectedUserRepoCreateError)
			svc := NewUserService(&repo)
			result, err := svc.Create(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			if err == nil {
				assert.Equal(t, tc.expectedSvc.Name, result.Name)
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
	passwordBeforeUpdated, _ := bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
	passwordAfterUpdated, _ := bcrypt.GenerateFromPassword([]byte("7654321"), bcrypt.DefaultCost)
	testCases := []struct {
		name                          string
		request                       *request.UpdateUserRequest
		requestUserRepoFindByID       int64
		expectedUserRepoFindByID      *model.User
		expectedUserRepoFindByIDError error
		expectedUserRepoUpdate        *model.User
		expectedUserRepoUpdateError   error
		expectedSvc                   *response.UserResponse
		expectedSvcError              error
	}{
		{
			name: "Update user with required fields",
			request: &request.UpdateUserRequest{
				ID:       1,
				Name:     "Arfian",
				Password: "7654321",
			},
			requestUserRepoFindByID: 1,
			expectedUserRepoFindByID: &model.User{
				ID:       1,
				Name:     "Widdy Arfiansyah",
				Password: string(passwordBeforeUpdated),
			},
			expectedUserRepoUpdate: &model.User{
				ID:       1,
				Name:     "Arfian",
				Username: "wdyarfn",
				Password: string(passwordAfterUpdated),
			},
			expectedSvc: &response.UserResponse{
				ID:        1,
				Name:      "Arfian",
				Username:  "wdyarfn",
				CreatedAt: "0001-01-01 00:00:00 +0000 UTC",
				UpdatedAt: "0001-01-01 00:00:00 +0000 UTC",
			},
			expectedUserRepoUpdateError:   nil,
			expectedSvcError:              nil,
			expectedUserRepoFindByIDError: nil,
		},
		{
			name: "User doesnt exists with given ID when updating data",
			request: &request.UpdateUserRequest{
				ID:       1,
				Name:     "Arfian",
				Password: "7654321",
			},
			requestUserRepoFindByID:       1,
			expectedUserRepoFindByID:      nil,
			expectedUserRepoUpdate:        nil,
			expectedSvc:                   nil,
			expectedUserRepoUpdateError:   errors.New("getting an error"),
			expectedSvcError:              errors.New(response.NotFound),
			expectedUserRepoFindByIDError: errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			repo.On("FindByID", ctx, tc.requestUserRepoFindByID).Return(tc.expectedUserRepoFindByID, tc.expectedUserRepoFindByIDError)
			repo.On("Update", ctx, mock.Anything).Return(tc.expectedUserRepoUpdate, tc.expectedUserRepoUpdateError)
			svc := NewUserService(&repo)
			result, err := svc.Update(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}

			assert.Equal(t, tc.expectedSvc, result)

			// really changed
			if err == nil {
				assert.Equal(t, tc.expectedUserRepoFindByID.Name, result.Name)
			}
		})
	}
}

func TestUserService_Delete(t *testing.T) {
	testCases := []struct {
		name                          string
		request                       int64
		expectedUserRepoFindByID      *model.User
		expectedUserRepoFindByIDError error
		expectedUserRepoDeleteError   error
		expectedSvcError              error
	}{
		{
			name:    "User exists with given ID",
			request: 1,
			expectedUserRepoFindByID: &model.User{
				ID:   1,
				Name: "Widdy Arfiansyah",
			},
			expectedUserRepoDeleteError:   nil,
			expectedSvcError:              nil,
			expectedUserRepoFindByIDError: nil,
		},
		{
			name:                          "User doesnt exists with given ID when deleting data",
			request:                       1,
			expectedUserRepoFindByID:      nil,
			expectedUserRepoDeleteError:   errors.New("getting an error"),
			expectedSvcError:              errors.New(response.NotFound),
			expectedUserRepoFindByIDError: errors.New(response.NotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			repo.On("FindByID", ctx, tc.request).Return(tc.expectedUserRepoFindByID, tc.expectedUserRepoFindByIDError)
			repo.On("Delete", ctx, tc.request).Return(tc.expectedUserRepoDeleteError)
			svc := NewUserService(&repo)
			err := svc.Delete(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}
		})
	}
}
