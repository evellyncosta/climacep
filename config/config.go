package config

import (
	"os"
)

// Config holds all configuration for the application
type Config struct {
	Port        string
	WeatherAPIKey string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	weatherAPIKey := os.Getenv("WEATHER_API_KEY")
	if weatherAPIKey == "" {
		weatherAPIKey = "YOUR_WEATHER_API_KEY" // Default for testing
	}

	return Config{
		Port:        port,
		WeatherAPIKey: weatherAPIKey,
	}
}