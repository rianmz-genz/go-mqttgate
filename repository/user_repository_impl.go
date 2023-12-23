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

func (repository UserRepositoryImpl) GetUserByEmail(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error) {
	result := db.WithContext(ctx).Where("email = ?", user.Email).First(&user)

	if result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (repository UserRepositoryImpl) Save(ctx context.Context, db *gorm.DB, user domain.User) domain.User {
	result := db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return user
}

func (repository UserRepositoryImpl) GetUserById(ctx context.Context, db *gorm.DB, userId uint) domain.User {
	user := domain.User{}
	result := db.WithContext(ctx).Where("id = ?", userId).Preload("Office").Preload("Role").First(&user)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return user
}

func (repository UserRepositoryImpl) Update(ctx context.Context, db *gorm.DB, user domain.User) domain.User {
	result := db.WithContext(ctx).Save(&user)

	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return user
}

func (repository UserRepositoryImpl) Delete(ctx context.Context, db *gorm.DB, userId uint) {
	result := db.WithContext(ctx).Delete(&domain.User{}, userId)

	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}
}
