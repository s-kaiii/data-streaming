package main

import (
	"fmt"
	"net/http"

	"real-time-streaming/internal/api"
	"real-time-streaming/internal/utils/config"
	"real-time-streaming/internal/utils/log"
)

func main() {
	// Load configurations
	config := config.LoadConfig()

	// Set up HTTP router
	router := api.NewRouter()

	// Start HTTP server
	log.Info(fmt.Sprintf("Starting server on %s", config.ServerAddr))
	if err := http.ListenAndServe(config.ServerAddr, router); err != nil {
		//TODO: change to log
		fmt.Println("Server failed:", err)
	}
}
