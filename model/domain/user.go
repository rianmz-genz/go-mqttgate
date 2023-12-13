package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Password string `json:"-"`
	OfficeID int64  `json:"officeId"`
	Office   Office
	Sessions []Session
}
