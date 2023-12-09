package model

import (
	"adriandidimqttgate/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func OpenConnection() {
	if DB != nil {
        return
    }
	
	dsn := "root:@tcp(127.0.0.1:3306)/qrgate?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	helper.PanicIfError(err)

    if err := db.AutoMigrate(&HousingArea{}, &User{}); err != nil {
        panic(err)
    }

	// Set up foreign key
	DB = db
}

