package seeder

import (
	"adriandidimqttgate/app"
	"adriandidimqttgate/model/domain"
)

func OfficeSeeder() {
	 req  := domain.Office{
		Name: "Coba",
		Code: "PERUM1",
		Address: "BLABLABLABLABLABAL",
	}

	if err := app.NewDBConnection().Create(&req).Error; err != nil {
        panic(err.Error())
	}
}