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
	result := db.Create(&enterActivity)

	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return enterActivity.ID
}