package service

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, sessionRepository repository.SessionRepository, DB *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service UserServiceImpl) GetUserById(ctx context.Context, userId uint) web.UserResponse {
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

func (service UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest, sessionId uint, userId uint) (web.UserUpdateResponse, error) {
	session, err := service.SessionRepository.GetSessionById(ctx, service.DB, sessionId)
	helper.PanicIfError(err)

	auth := service.UserRepository.GetUserById(ctx, service.DB, session.UserID)
	if auth.Role.Name != "Admin" && auth.ID != userId {
		return web.UserUpdateResponse{}, errors.New("unauthorized")
	}

	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	user := domain.User{
		Name:     request.Name,
		Email:    request.Email,
		OfficeID: request.OfficeID,
	}
	user.ID = userId

	return web.UserUpdateResponse{
		Name:     user.Name,
		Email:    user.Email,
		OfficeID: user.OfficeID,
	}, nil
}
