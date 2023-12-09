package seeder

import "adriandidimqttgate/model"

func HousingAreaSeeder() {
	 req  := model.HousingArea{
		Name: "Coba",
		Code: "PERUM1",
		Address: "BLABLABLABLABLABAL",
	}
	if err := model.DB.Create(&req).Error; err != nil {
        panic(err.Error())
	}
}