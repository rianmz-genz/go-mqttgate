package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Email    string `gorm:"type:varchar(100)" json:"email"`
	Password string `json:"-"`
	RoleID   uint   `json:"roleId"`
	OfficeID uint   `json:"officeId"`
	Office   Office
	Sessions []Session
	Role     Role
}
