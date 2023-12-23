package service

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/web"
	"adriandidimqttgate/repository"
	"context"

	"gorm.io/gorm"
)

type SessionServiceImpl struct {
	SessionRepository repository.SessionRepository
	DB                *gorm.DB
}

func NewSessionService(sessionRepository repository.SessionRepository, DB *gorm.DB) SessionService {
	return &SessionServiceImpl{
		SessionRepository: sessionRepository,
		DB:                DB,
	}
}

func (service *SessionServiceImpl) GetSessionById(ctx context.Context, sessionId uint) web.SessionResponse {
	session, err := service.SessionRepository.GetSessionById(ctx, service.DB, sessionId)
	helper.PanicIfError(err)

	return web.SessionResponse{
		SessionID: session.ID,
		Name:      session.User.Name,
		Email:     session.User.Email,
		UserId:    session.User.ID,
	}
}
