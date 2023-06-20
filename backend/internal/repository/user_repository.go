package repository

import (
	"context"
	"gorm.io/gorm"
	"inventory-management/backend/internal/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryContract {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) FindAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	err := repository.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (repository *UserRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	err := repository.DB.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := repository.DB.WithContext(ctx).Select("username", "password").Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repository *UserRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	err := repository.DB.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	err := repository.DB.WithContext(ctx).Select("name", "password").Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repository *UserRepository) Delete(ctx context.Context, id int64) error {
	var user model.User
	err := repository.DB.WithContext(ctx).Delete(&user, id).Error
	if err != nil {
		return err
	}

	return nil
}
