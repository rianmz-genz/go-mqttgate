package repository

import (
	"adriandidimqttgate/model/domain"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type SessionRepository interface {
	GetSessionById(ctx context.Context, db *gorm.DB, sessionId uint) (domain.Session, error)
	Save(ctx context.Context, db *gorm.DB, userId uint) (uint, error)
	DeleteSessionById(ctx context.Context, db *gorm.DB, sessionId uint) (uint, error)
}
