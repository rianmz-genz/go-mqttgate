package repository

import (
	"adriandidimqttgate/model/domain"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type EnterActivityRepository interface {
	Save(ctx context.Context, db *gorm.DB, enterActivity domain.EnterActivity) uint
	GetByUserIds(ctx context.Context, db *gorm.DB, userIds ...uint) []domain.EnterActivity
}
