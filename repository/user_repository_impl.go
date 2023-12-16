package repository

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"
	"context"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) GetUserByEmail(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	result := db.Where("email = ?", user.Email).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (respository *UserRepositoryImpl) Save(ctx context.Context, db *gorm.DB, user domain.User) domain.User {
	result := db.Create(&user)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return user
}

func (respository *UserRepositoryImpl) GetUserById(ctx context.Context, db *gorm.DB, userId uint) domain.User {
	user := domain.User{}
	result := db.Where("id = ?", userId).Preload("Office").First(&user)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return user
}
