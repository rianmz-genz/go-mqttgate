package domain

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserID uint `json:"userId"`
	User   User
}
