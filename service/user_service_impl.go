package service

import (
	"adriandidimqttgate/model/domain"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"context"

	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
}

func NewUserService(userRepository repository.UserRepository, DB *gorm.DB) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
	}
}

func (service *UserServiceImpl) GetUserById(ctx context.Context, userId uint) web.UserResponse {
	user := domain.User{}
	user = service.UserRepository.GetUserById(ctx, service.DB, userId)

	officeResponse := web.OfficeResponse{
		ID:      user.OfficeID,
		Name:    user.Office.Name,
		Code:    user.Office.Code,
		Address: user.Office.Address,
	}
	roleResponse := web.RoleResponse{
		ID:   user.RoleID,
		Name: user.Role.Name,
	}
	return web.UserResponse{
		Name:   user.Name,
		Email:  user.Email,
		Office: officeResponse,
		Role:   roleResponse,
	}
}
