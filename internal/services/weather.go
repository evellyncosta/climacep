package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
	// URL encode the city name to handle special characters
	encodedCity := url.QueryEscape(city)
	
	url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s&aqi=no", s.apiKey, encodedCity)
	
	// Print debug information
	fmt.Printf("WeatherAPI Request: URL=%s\n", url)
	fmt.Printf("WeatherAPI Key: %s\n", s.apiKey)
	fmt.Printf("Original City: %s, Encoded City: %s\n", city, encodedCity)

	resp, err := s.client.Get(url)
	if err != nil {
		fmt.Printf("WeatherAPI Error: %v\n", err)
		return nil, fmt.Errorf("%w: %v", ErrWeatherAPIRequest, err)
	}
	defer resp.Body.Close()

	// Print debug information
	fmt.Printf("WeatherAPI Response: Status=%d\n", resp.StatusCode)

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("WeatherAPI Error Response: %s\n", string(body))
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