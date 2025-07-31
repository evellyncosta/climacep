package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/evellyn/climacep/config"
	"github.com/evellyn/climacep/internal/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Create services
	viaCEPService := services.NewViaCEPService()
	weatherService := services.NewWeatherService(cfg.WeatherAPIKey)
	tempConverter := services.NewTemperatureConverter()

	// Create handlers
	handlers := NewHandlers(viaCEPService, weatherService, tempConverter)

	// Set up routes
	http.HandleFunc("/weather", handlers.HandleWeatherByCEP)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Printf("Starting server on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}