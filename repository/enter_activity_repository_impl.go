package repository

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type EnterActivityRepositoryImpl struct {
}

func NewEnterActivityRepository() EnterActivityRepository {
	return &EnterActivityRepositoryImpl{}
}

func (repository EnterActivityRepositoryImpl) Save(ctx context.Context, db *gorm.DB, enterActivity domain.EnterActivity) uint {
	result := db.WithContext(ctx).Create(&enterActivity)

	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return enterActivity.ID
}

func (repository EnterActivityRepositoryImpl) GetByUserIds(ctx context.Context, db *gorm.DB, userIds ...uint) []domain.EnterActivity {
	var enterActivities []domain.EnterActivity
	result := db.WithContext(ctx).Where("user_id IN ?", userIds).Preload("User").Find(&enterActivities)
	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return enterActivities
}
