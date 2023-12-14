package repository

import (
	"adriandidimqttgate/model/domain"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SessionReposiotyImpl struct {
}

func NewSessionRepository() SessionReposioty {
	return &SessionReposiotyImpl{}
}

func (repository *SessionReposiotyImpl) GetSessionById(ctx context.Context, db *gorm.DB, sessionId uint) (domain.Session, error) {
	var session = domain.Session{}

	result := db.Where("id = ?", sessionId).Preload("User").First(&session)

	if result.Error != nil {
		return session, result.Error
	}

	return session, nil
}

func (repository *SessionReposiotyImpl) Save(ctx context.Context, db *gorm.DB, userId uint) (uint, error) {
	session := domain.Session{UserID: userId}
	result := db.Create(&session)

	if result.Error != nil {
		return 0, result.Error
	}

	return session.ID, nil
}

func (repository *SessionReposiotyImpl) DeleteSessionById(ctx context.Context, db *gorm.DB, sessionId uint) (uint, error) {
	session := domain.Session{}
	result := db.Where("id = ?", sessionId).Delete(&session)

	if result.Error != nil {
		return 0, result.Error
	}
	return session.ID, nil
}
