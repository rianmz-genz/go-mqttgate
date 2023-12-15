package seeder

import (
	"adriandidimqttgate/app"
	"adriandidimqttgate/model/domain"
)

func SessionSeeder() {

	req := domain.Session{
		UserID: 1,
	}

	if err := app.NewDBConnection().Create(&req).Error; err != nil {
		panic(err.Error())
	}
}