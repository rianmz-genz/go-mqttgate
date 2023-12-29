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

func (repository UserRepositoryImpl) GetUsersByOfficeId(ctx context.Context, db *gorm.DB, officeId uint) ([]domain.User, error) {
	var users []domain.User
	result := db.WithContext(ctx).Where("office_id = ?", officeId).Preload("Office").Preload("Role").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (repository UserRepositoryImpl) GetEmployeeByOfficeId(ctx context.Context, db *gorm.DB, officeId uint) ([]domain.User, error) {
	var users []domain.User
	result := db.WithContext(ctx).Where("office_id = ? and role_id = 1", officeId).Preload("Office").Preload("Role").Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
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


func (repository UserRepositoryImpl) SaveEmployee(ctx context.Context, db *gorm.DB, user domain.User) domain.User {
	user.RoleID = 2
	result := db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return user
}