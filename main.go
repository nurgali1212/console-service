package main

import (
	"console-service/handler"
	"console-service/service"
)

func main() {

	services := service.NewService()
	handlers := handler.NewHandler(services)

	router := handlers.SetupRouter()

	router.Run(":8081")
}
