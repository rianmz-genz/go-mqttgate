package domain

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string `gorm:"type:varchar(200)" json:"name"`
	Users []User
}
