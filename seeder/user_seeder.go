package seeder

import (
	"adriandidimqttgate/app"
	"adriandidimqttgate/helper"
	"adriandidimqttgate/model/domain"
)

func UserSeeder() {
	email1 := "user1@gmail.com"
	email2 := "ucup@gmail.com"
	email3 := "admin@gmail.com"
	password := "password"
	passwordEncrypted := helper.HashPassword(password)

	user1 := domain.User{
		Name:     "Coba",
		Email:    email1,
		Password: string(passwordEncrypted),
		OfficeID: 1,
		RoleID:   1,
	}

	user2 := domain.User{
		Name:     "Ucup",
		Email:    email2,
		Password: string(passwordEncrypted),
		OfficeID: 1,
		RoleID:   1,
	}

	admin1 := domain.User{
		Name:     "Admin",
		Email:    email3,
		Password: string(passwordEncrypted),
		OfficeID: 1,
		RoleID:   1,
	}

	users := []domain.User{user1, user2, admin1}

	if err := app.NewDBConnection().CreateInBatches(users, 3).Error; err != nil {
		panic(err.Error())
	}
}
