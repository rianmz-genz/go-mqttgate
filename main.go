package main

import (
	"adriandidimqttgate/app"
	"adriandidimqttgate/controller"
	"adriandidimqttgate/model"
	"fmt"

	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Hello!")

	token := app.NewMqttClient().Publish("testtopic/esp8266", 0, false, "Hello from GO!")
	token.Wait()
	r := gin.Default()
	model.OpenConnection()
	// seeder.HousingAreaSeeder()
	api := r.Group("/api")
	auth := api.Group("/auth")

	auth.POST("/login", controller.Login)
	auth.POST("/register", controller.Register)

	r.Run(":6666")
}