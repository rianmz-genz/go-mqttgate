package domain

import (
	"time"

	"gorm.io/gorm"
)

type EnterActivity struct {
	gorm.Model
	UserID  uint `json:"userId"`
	User    User
	EnterAt time.Time
}
