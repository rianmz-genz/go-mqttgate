package domain

import "gorm.io/gorm"

type Office struct {
	gorm.Model
	Name    string `gorm:"type:varchar(100)" json:"name"`
	Code    string `gorm:"type:varchar(6);index:;unique" json:"code"`
	Address string `gorm:"type:varchar(300)" json:"address"`
	Users   []User
}
