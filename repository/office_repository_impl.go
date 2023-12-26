package repository

import (
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type OfficeRepositoryImpl struct {
}

func NewOfficeRepository() OfficeRepository {
	return &OfficeRepositoryImpl{}
}

func (repository *OfficeRepositoryImpl) FindOfficeByCode(ctx context.Context, db *gorm.DB, code string) domain.Office {
	office := domain.Office{}
	db.WithContext(ctx).Where("id = ?", code).Preload("").First(office)

	return office
}

func (repository *OfficeRepositoryImpl) Save(ctx context.Context, db *gorm.DB, office domain.Office) domain.Office {
	result := db.WithContext(ctx).Create(&office)

	if result.Error != nil {
		helper.PanicIfError(result.Error)
	}

	return office
}
