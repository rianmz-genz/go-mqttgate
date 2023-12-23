package repository

import (
	"adriandidimqttgate/model/domain"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(ctx context.Context, db *gorm.DB, user domain.User) (domain.User, error)
	Save(ctx context.Context, db *gorm.DB, user domain.User) domain.User
	GetUserById(ctx context.Context, db *gorm.DB, userId uint) domain.User
	Update(ctx context.Context, db *gorm.DB, user domain.User) domain.User
}
