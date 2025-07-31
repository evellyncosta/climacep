package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/evellyn/climacep/internal/models"
)

var (
	ErrWeatherAPIRequest = errors.New("error in weather API request")
	ErrCityNotFound      = errors.New("city not found")
)

// WeatherService handles interactions with the Weather API
type WeatherService struct {
	client *http.Client
	apiKey string
}

// NewWeatherService creates a new Weather service
func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		apiKey: apiKey,
	}
}

// GetWeatherByCity retrieves weather information by city name
func (s *WeatherService) GetWeatherByCity(city string) (*models.WeatherAPIResponse, error) {
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", s.apiKey, city)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrWeatherAPIRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode >= 400 {
		return nil, ErrCityNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrWeatherAPIRequest, err)
	}

	var weatherResponse models.WeatherAPIResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnmarshalResponse, err)
	}

	return &weatherResponse, nil
}