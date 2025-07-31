package services

import "github.com/evellyn/climacep/internal/models"

// ViaCEPServicer is an interface for ViaCEP service
type ViaCEPServicer interface {
	GetLocationByCEP(cep string) (*models.ViaCEPResponse, error)
}

// WeatherServicer is an interface for Weather service
type WeatherServicer interface {
	GetWeatherByCity(city string) (*models.WeatherAPIResponse, error)
}

// TemperatureConverterr is an interface for temperature converter
type TemperatureConverterr interface {
	ConvertTemperature(tempC float64) models.WeatherResponse
}