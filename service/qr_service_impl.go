package service

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"encoding/json"
	"errors"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/go-playground/validator/v10"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type QrServiceImpl struct {
	EnterActivityRepository repository.EnterActivityRepository
	OfficeRepository        repository.OfficeRepository
	UserRepository          repository.UserRepository
	SessionRepository       repository.SessionRepository
	DB                      *gorm.DB
	Validate                *validator.Validate
	MQTT                    mqtt.Client
}

func NewQrService(
	enterActivityRepository repository.EnterActivityRepository,
	officeRepository repository.OfficeRepository,
	sessionRepository repository.SessionRepository,
	userRepository repository.UserRepository,
	DB *gorm.DB, validate *validator.Validate,
	mqtt mqtt.Client,
) QrService {
	return &QrServiceImpl{
		EnterActivityRepository: enterActivityRepository,
		OfficeRepository:        officeRepository,
		SessionRepository:       sessionRepository,
		UserRepository:          userRepository,
		DB:                      DB,
		Validate:                validate,
		MQTT:                    mqtt,
	}
}

func (service QrServiceImpl) ScanQr(ctx context.Context, request web.ScanQrRequest, sessionId uint) web.ScanQrResponse {
	// validate request
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// get session
	session, err := service.SessionRepository.GetSessionById(ctx, service.DB, sessionId)
	helper.PanicIfError(err)
	// get user with office
	user := service.UserRepository.GetUserById(ctx, service.DB, session.UserID)
	// check user.office.code equal to request.code
	if user.Office.Code != request.Code {
		helper.PanicIfError(errors.New("you have not access"))
	}
	// if equal create enter_activity and send mqtt message

	// create enter_activity
	enterActivity := domain.EnterActivity{
		UserID:  user.ID,
		EnterAt: time.Now(),
	}

	service.EnterActivityRepository.Save(ctx, service.DB, enterActivity)

	// send mqtt message
	scanQrResponse := web.ScanQrResponse{
		Username:   user.Name,
		OfficeName: user.Office.Name,
	}
	responseJson, err := json.Marshal(scanQrResponse)
	helper.PanicIfError(err)
	service.MQTT.Publish("office/"+string(user.Office.Code), 0, false, responseJson)

	return scanQrResponse
}
