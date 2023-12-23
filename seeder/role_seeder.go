package seeder

import (
	"adriandidimqttgate/app"
	"adriandidimqttgate/model/domain"
)

func RoleSeeder() {
	employee := domain.Role{
		Name: "Employee",
	}

	admin := domain.Role{
		Name: "Admin",
	}

	superAdmin := domain.Role{
		Name: "Super Admin",
	}

	roles := []domain.Role{employee, admin, superAdmin}

	if err := app.NewDBConnection().CreateInBatches(roles, 3).Error; err != nil {
		panic(err.Error())
	}
}
