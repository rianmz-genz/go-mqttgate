package app

import (
	"adriandidimqttgate/config"
	"adriandidimqttgate/helper"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDBConnection() *gorm.DB {
	dsn := config.GetEnv("DB_USERNAME") + ":" + config.GetEnv("DB_PASSWORD") + "@tcp(" + config.GetEnv("DB_HOSTNAME") + ":3306)/" + config.GetEnv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	sqlSB, err := db.DB()
	helper.PanicIfError(err)

	sqlSB.SetMaxIdleConns(5)
	sqlSB.SetMaxOpenConns(20)
	sqlSB.SetConnMaxLifetime(60 * time.Minute)
	sqlSB.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
