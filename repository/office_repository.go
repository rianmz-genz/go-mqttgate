package repository

import (
	"adriandidimqttgate/model/domain"

	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type OfficeRepository interface {
	FindOfficeByCode(ctx context.Context, db *gorm.DB, code string) domain.Office
	Save(ctx context.Context, db *gorm.DB, office domain.Office) domain.Office
}
