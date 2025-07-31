package services

import "github.com/evellyn/climacep/internal/models"

// TemperatureConverter converts temperatures between different scales
type TemperatureConverter struct{}

// NewTemperatureConverter creates a new temperature converter
func NewTemperatureConverter() *TemperatureConverter {
	return &TemperatureConverter{}
}

// ConvertTemperature converts temperature from Celsius to Fahrenheit and Kelvin
func (tc *TemperatureConverter) ConvertTemperature(tempC float64) models.WeatherResponse {
	// Convert to Fahrenheit: F = C * 1.8 + 32
	tempF := tempC*1.8 + 32

	// Convert to Kelvin: K = C + 273
	tempK := tempC + 273.0

	return models.WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}
}