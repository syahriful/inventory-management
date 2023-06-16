package main

import (
	"fmt"
	"inventory-management/backend/cmd/config"
	"inventory-management/backend/internal/http"
	"log"
	"os"
)

func main() {
	configuration := config.New()

	// Init Log File
	file, err := os.OpenFile("./logs/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("There is something wrong with the log file", err)
	}
	defer file.Close()

	// Init Routing
	app, err := http.NewInitializedRoutes(configuration, file)
	if err != nil {
		log.Fatalln("There is something wrong with the server", err)
	}

	// Start Server
	port := fmt.Sprintf(":%s", configuration.Get("APP_PORT"))
	err = app.Listen(port)
	if err != nil {
		log.Fatalln("Cannot connect to server", err)
	}
}
