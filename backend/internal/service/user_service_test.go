package service

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"inventory-management/backend/internal/http/request"
	response "inventory-management/backend/internal/http/response"
	"inventory-management/backend/internal/model"
	repository "inventory-management/backend/internal/repository/mock"
	third_party "inventory-management/backend/internal/third_party/elasticsearch"
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
					CreatedAt: "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
				},
				{
					ID:        2,
					Name:      "Arfiansyah",
					Username:  "arfn",
					CreatedAt: "0001-01-01 07:00:00 +0700 +07",
					UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
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
			var es third_party.ElasticsearchMock
			repo.On("FindAll", ctx, 0, 10).Return(tc.expectedUserRepoFindAll, tc.expectedUserRepoFindAllError)
			svc := NewUserService(&repo, &es)
			result, err := svc.FindAll(ctx, 0, 10)
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
				CreatedAt: "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
			},
			expectedUserRepoFindByIDError: nil,
			expectedSvcError:              nil,
		},
		{
			name:                          "User doesnt exists with given ID",
			request:                       1,
			expectedUserRepoFindByID:      nil,
			expectedSvc:                   nil,
			expectedUserRepoFindByIDError: errors.New(response.ErrorNotFound),
			expectedSvcError:              errors.New(response.ErrorNotFound),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			var es third_party.ElasticsearchMock
			repo.On("FindByID", ctx, tc.request).Return(tc.expectedUserRepoFindByID, tc.expectedUserRepoFindByIDError)
			svc := NewUserService(&repo, &es)
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
			expectedUserRepoFindByUsernameError: errors.New(response.ErrorNotFound),
			expectedSvcError:                    errors.New(response.ErrorNotFound),
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
			expectedSvcError:                    errors.New(response.ErrorInvalidPassword),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			var es third_party.ElasticsearchMock
			repo.On("FindByUsername", ctx, tc.request.Username).Return(tc.expectedUserRepoFindByUsername, tc.expectedUserRepoFindByUsernameError)
			svc := NewUserService(&repo, &es)
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
		expected3rdParty                    *response.UserResponse
		expectedSvcError                    error
		expected3rdPartyError               error
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
				ID:        1,
				Name:      "Widdy Arfiansyah",
				Username:  "wdyarfn",
				CreatedAt: "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
			},
			expected3rdParty: &response.UserResponse{
				ID:        1,
				Name:      "Widdy Arfiansyah",
				Username:  "wdyarfn",
				CreatedAt: "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
			},
			expectedUserRepoCreateError:         nil,
			expectedSvcError:                    nil,
			expectedUserRepoFindByUsernameError: errors.New(response.ErrorNotFound),
			expected3rdPartyError:               nil,
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
			expected3rdParty: &response.UserResponse{
				ID:        1,
				Name:      "Widdy Arfiansyah",
				Username:  "wdyarfn",
				CreatedAt: "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
			},
			expectedUserRepoCreate:              nil,
			expectedSvc:                         nil,
			expectedUserRepoCreateError:         errors.New("getting an error"),
			expectedSvcError:                    errors.New(response.ErrorUsernameExists),
			expectedUserRepoFindByUsernameError: nil,
			expected3rdPartyError:               errors.New("getting an error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			var es third_party.ElasticsearchMock
			repo.On("FindByUsername", ctx, tc.requestRepo.Username).Return(tc.expectedUserRepoFindByUsername, tc.expectedUserRepoFindByUsernameError)
			repo.On("Create", ctx, mock.Anything).Return(tc.expectedUserRepoCreate, tc.expectedUserRepoCreateError)
			es.On("Create", ctx, "users", tc.expected3rdParty, tc.expected3rdParty.ID).Return(tc.expected3rdPartyError)
			svc := NewUserService(&repo, &es)
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
		expected3rdParty              *response.UserResponse
		expectedSvcError              error
		expected3rdPartyError         error
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
				CreatedAt: "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
			},
			expected3rdParty: &response.UserResponse{
				ID:        1,
				Name:      "Arfian",
				Username:  "wdyarfn",
				CreatedAt: "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
			},
			expectedUserRepoUpdateError:   nil,
			expectedSvcError:              nil,
			expectedUserRepoFindByIDError: nil,
			expected3rdPartyError:         nil,
		},
		{
			name: "User doesnt exists with given ID when updating data",
			request: &request.UpdateUserRequest{
				ID:       1,
				Name:     "Arfian",
				Password: "7654321",
			},
			expected3rdParty: &response.UserResponse{
				ID:        1,
				Name:      "Arfian",
				Username:  "wdyarfn",
				CreatedAt: "0001-01-01 07:00:00 +0700 +07",
				UpdatedAt: "0001-01-01 07:00:00 +0700 +07",
			},
			requestUserRepoFindByID:       1,
			expectedUserRepoFindByID:      nil,
			expectedUserRepoUpdate:        nil,
			expectedSvc:                   nil,
			expectedUserRepoUpdateError:   errors.New("getting an error"),
			expectedSvcError:              errors.New(response.ErrorNotFound),
			expectedUserRepoFindByIDError: errors.New(response.ErrorNotFound),
			expected3rdPartyError:         errors.New("getting an error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			var es third_party.ElasticsearchMock
			repo.On("FindByID", ctx, tc.requestUserRepoFindByID).Return(tc.expectedUserRepoFindByID, tc.expectedUserRepoFindByIDError)
			repo.On("Update", ctx, mock.Anything).Return(tc.expectedUserRepoUpdate, tc.expectedUserRepoUpdateError)
			es.On("Update", ctx, "users", tc.expected3rdParty, tc.expected3rdParty.ID).Return(tc.expected3rdPartyError)
			svc := NewUserService(&repo, &es)
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
		expected3rdPartyError         error
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
			expected3rdPartyError:         nil,
		},
		{
			name:                          "User doesnt exists with given ID when deleting data",
			request:                       1,
			expectedUserRepoFindByID:      nil,
			expectedUserRepoDeleteError:   errors.New("getting an error"),
			expectedSvcError:              errors.New(response.ErrorNotFound),
			expectedUserRepoFindByIDError: errors.New(response.ErrorNotFound),
			expected3rdPartyError:         errors.New("getting an error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			var repo repository.UserRepositoryMock
			var es third_party.ElasticsearchMock
			repo.On("FindByID", ctx, tc.request).Return(tc.expectedUserRepoFindByID, tc.expectedUserRepoFindByIDError)
			repo.On("Delete", ctx, tc.request).Return(tc.expectedUserRepoDeleteError)
			es.On("Delete", ctx, "users", tc.request).Return(tc.expected3rdPartyError)
			svc := NewUserService(&repo, &es)
			err := svc.Delete(ctx, tc.request)
			if tc.expectedSvcError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedSvcError.Error(), err.Error())
			}
		})
	}
}
