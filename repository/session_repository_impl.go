package repository

import (
	"adriandidimqttgate/model/domain"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SessionRepositoryImpl struct {
}

func NewSessionRepository() SessionRepository {
	return &SessionRepositoryImpl{}
}

func (repository *SessionRepositoryImpl) GetSessionById(ctx context.Context, db *gorm.DB, sessionId uint) (domain.Session, error) {
	var session = domain.Session{}

	result := db.WithContext(ctx).Where("id = ?", sessionId).Preload("User").First(&session)

	if result.Error != nil {
		return session, result.Error
	}

	return session, nil
}

func (repository *SessionRepositoryImpl) Save(ctx context.Context, db *gorm.DB, userId uint) (uint, error) {
	session := domain.Session{UserID: userId}
	result := db.WithContext(ctx).Create(&session)

	if result.Error != nil {
		return 0, result.Error
	}

	return session.ID, nil
}

func (repository *SessionRepositoryImpl) DeleteSessionById(ctx context.Context, db *gorm.DB, sessionId uint) (uint, error) {
	session := domain.Session{}
	result := db.WithContext(ctx).Where("id = ?", sessionId).Delete(&session)

	if result.Error != nil {
		return 0, result.Error
	}
	return session.ID, nil
}
