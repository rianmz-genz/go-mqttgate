package main

import (
	"adriandidimqttgate/app"
	"fmt"
)


func main() {
	fmt.Println("Hello!")

	token := app.NewMqttClient().Publish("testtopic/esp8266", 0, false, "Hello from GO!")
	token.Wait()
}