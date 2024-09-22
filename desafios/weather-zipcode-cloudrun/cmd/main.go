package main

import (
	"fmt"

	"cloudrun/configs"
	"cloudrun/internal/infra/webserver"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	newWebserver := webserver.NewWebServer(configs.WebServerPort, configs.WeatherApiKey)
	newWebserver.AddHandler("/weather/", webserver.WeatherZipCodeHandler)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	newWebserver.Start()

}
