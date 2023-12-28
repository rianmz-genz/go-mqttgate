package service

import (
	"adriandidimqttgate/exception"
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"context"
	"encoding/json"
	"errors"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type OfficeServiceImpl struct {
	OfficeRepository        repository.OfficeRepository
	UserRepository          repository.UserRepository
	SessionRepository       repository.SessionRepository
	EnterActivityRepository repository.EnterActivityRepository
	DB                      *gorm.DB
	MQTT                    mqtt.Client
	Validate                *validator.Validate
}

func NewOfficeService(
	officeRepository repository.OfficeRepository,
	userrepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
	enterActivityRepository repository.EnterActivityRepository,
	DB *gorm.DB,
	mqtt mqtt.Client,
	validate *validator.Validate,
) OfficeService {

	return &OfficeServiceImpl{
		OfficeRepository:        officeRepository,
		UserRepository:          userrepository,
		SessionRepository:       sessionRepository,
		EnterActivityRepository: enterActivityRepository,
		DB:                      DB,
		MQTT:                    mqtt,
		Validate:                validate,
		
	}
}

func (service OfficeServiceImpl) GetEntryActivities(ctx context.Context, sessionId uint, officeId uint) ([]web.EntryActivityResponse, error) {
	session, err := service.SessionRepository.GetSessionById(ctx, service.DB, sessionId)
	helper.PanicIfError(err)

	auth := service.UserRepository.GetUserById(ctx, service.DB, session.UserID)
	isSuperAdmin := auth.Role.Name == "Super Admin"
	isRoleAdmin := auth.Role.Name == "Admin"
	isAuthOfficeIdEqualToRequestOfficeId := officeId == auth.OfficeID
	isAdminAuthorized := isRoleAdmin && isAuthOfficeIdEqualToRequestOfficeId
	if !isSuperAdmin && !isAdminAuthorized {
		return nil, errors.New("unauthorized")
	}

	officeEmployees, err := service.UserRepository.GetUsersByOfficeId(ctx, service.DB, officeId)
	helper.PanicIfError(err)

	var userIds []uint
	for _, item := range officeEmployees {
		userIds = append(userIds, item.ID)
	}

	enterActivities := service.EnterActivityRepository.GetByUserIds(ctx, service.DB, userIds...)

	var enterActivitiesResponse []web.EntryActivityResponse
	for _, item := range enterActivities {
		enterActivitiesResponse = append(enterActivitiesResponse, web.EntryActivityResponse{
			Name:    item.User.Name,
			EntryAt: item.CreatedAt,
		})
	}

	return enterActivitiesResponse, nil
}

func (service OfficeServiceImpl) CloseGate(ctx context.Context, sessionId uint, officeId uint) (web.CloseGateResponse, error) {
	session, err := service.SessionRepository.GetSessionById(ctx, service.DB, sessionId)
	helper.PanicIfError(err)

	auth := service.UserRepository.GetUserById(ctx, service.DB, session.UserID)

	if officeId != auth.OfficeID {
		return web.CloseGateResponse{}, errors.New("unauthorized")
	}

	closeResponse := web.MqttGateResponse{
		Type: "Close",
		Name: auth.Name,
	}
	jsonMqttResponse, err := json.Marshal(closeResponse)
	helper.PanicIfError(err)

	service.MQTT.Publish("office/"+string(auth.Office.Code), 0, false, jsonMqttResponse)

	closeGaateResponse := web.CloseGateResponse{
		Name:    auth.Name,
		CloseAt: time.Now(),
	}

	return closeGaateResponse, nil
}

func (service OfficeServiceImpl) Add(ctx context.Context, request web.AddOfficeRequest, sessionId uint) web.AddOfficeResponse {
	session, err := service.SessionRepository.GetSessionById(ctx, service.DB, sessionId)
	helper.PanicIfError(err)

	auth := service.UserRepository.GetUserById(ctx, service.DB, session.UserID)

	if auth.Role.Name != "Super Admin" {
		panic(exception.NewForbiddenError("Forbidden: You are not Super Admin"))
	}

	errValidation := service.Validate.Struct(request)
	helper.PanicIfError(errValidation)

	office := domain.Office{
		Name:    request.Name,
		Code:    request.Code,
		Address: request.Address,
	}

	office = service.OfficeRepository.Save(ctx, service.DB, office)

	return web.AddOfficeResponse{
		ID:      office.ID,
		Name:    office.Name,
		Code:    office.Code,
		Address: office.Address,
	}
}


func (service OfficeServiceImpl) GetAllOffice(ctx context.Context, sessionId uint) []web.OfficeResponse {
	session, err := service.SessionRepository.GetSessionById(ctx, service.DB, sessionId)
	helper.PanicIfError(err)

	auth := service.UserRepository.GetUserById(ctx, service.DB, session.UserID)

	if auth.Role.Name == "Employee" {
		panic(exception.NewForbiddenError("Forbidden: You are not Super Admin or Admin"))
	}
	offices := service.OfficeRepository.FindAll(ctx, service.DB) // Retrieve all offices
	   // Map offices to web.AddOfficeResponse objects
    var responses []web.OfficeResponse
    for _, office := range offices {
        response := web.OfficeResponse{
           ID: office.ID,
		   Name: office.Name,
		   Code: office.Code,
		   Address: office.Address,
        }
        responses = append(responses, response)
    }

    return responses
}
