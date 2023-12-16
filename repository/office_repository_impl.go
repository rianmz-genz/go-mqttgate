package repository

import (
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
	db.Where("id = ?", code).Preload("").First(office)

	return office
}
