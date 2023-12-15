package seeder

import (
	"adriandidimqttgate/app"
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"
)

func UserSeeder() {
	email1 := "user1@gmail.com"
	password := "password"
	passwordEncrypted := helper.HashPassword(password)

	req := domain.User{
		Name:     "Coba",
		Email:    email1,
		Password: string(passwordEncrypted),
		OfficeID: 1,
	}

	if err := app.NewDBConnection().Create(&req).Error; err != nil {
		panic(err.Error())
	}
}
