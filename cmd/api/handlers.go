package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/evellyn/climacep/internal/services"
	"github.com/evellyn/climacep/internal/validator"
)

// Handlers holds all the dependencies for HTTP handlers
type Handlers struct {
	viaCEPService     services.ViaCEPServicer
	weatherService    services.WeatherServicer
	tempConverter     services.TemperatureConverterr
}

// NewHandlers creates a new handlers instance
func NewHandlers(viaCEPService services.ViaCEPServicer, weatherService services.WeatherServicer, tempConverter services.TemperatureConverterr) *Handlers {
	return &Handlers{
		viaCEPService:     viaCEPService,
		weatherService:    weatherService,
		tempConverter:     tempConverter,
	}
}

// HandleWeatherByCEP handles requests for weather by CEP
func (h *Handlers) HandleWeatherByCEP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get CEP from query parameters
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		http.Error(w, "CEP parameter is required", http.StatusBadRequest)
		return
	}

	// Validate and format CEP
	if !validator.ValidateCEP(cep) {
		writeJSONError(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	// Format CEP
	cep = validator.FormatCEP(cep)

	// Get location information by CEP
	location, err := h.viaCEPService.GetLocationByCEP(cep)
	if err != nil {
		if errors.Is(err, services.ErrCEPNotFound) {
			writeJSONError(w, "can not find zipcode", http.StatusNotFound)
			return
		}
		http.Error(w, "Error fetching location information", http.StatusInternalServerError)
		return
	}

	// Get weather information by city
	weather, err := h.weatherService.GetWeatherByCity(location.Localidade)
	if err != nil {
		if errors.Is(err, services.ErrCityNotFound) {
			writeJSONError(w, "City not found in weather API", http.StatusNotFound)
			return
		}
		http.Error(w, "Error fetching weather information", http.StatusInternalServerError)
		return
	}

	// Convert temperature
	temps := h.tempConverter.ConvertTemperature(weather.Current.TempC)

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(temps)
}

// writeJSONError writes a JSON error response
func writeJSONError(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}