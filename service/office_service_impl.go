package service

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"context"
	"encoding/json"
	"errors"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gorm.io/gorm"
)

type OfficeServiceImpl struct {
	UserRepository          repository.UserRepository
	SessionRepository       repository.SessionRepository
	EnterActivityRepository repository.EnterActivityRepository
	DB                      *gorm.DB
	MQTT                    mqtt.Client
}

func NewOfficeService(
	userrepository repository.UserRepository,
	sessionRepository repository.SessionRepository,
	enterActivityRepository repository.EnterActivityRepository,
	DB *gorm.DB, mqtt mqtt.Client,
) OfficeService {

	return &OfficeServiceImpl{
		UserRepository:          userrepository,
		SessionRepository:       sessionRepository,
		EnterActivityRepository: enterActivityRepository,
		DB:                      DB,
		MQTT:                    mqtt,
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
