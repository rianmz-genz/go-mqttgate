package service

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewAuthService(userRepository repository.UserRepository, DB *gorm.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Register(ctx context.Context, request web.RegisterRequest) web.RegisterResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user := domain.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: helper.HashPassword(request.Password),
		OfficeID: request.OfficeID,
	}

	user = service.UserRepository.Save(ctx, service.DB, user)

	return web.RegisterResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		OfficeID: user.OfficeID,
	}
}
