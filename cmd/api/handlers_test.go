package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/evellyn/climacep/internal/models"
	"github.com/evellyn/climacep/internal/services"
)

// Mock ViaCEP service
type mockViaCEPService struct{}

func (m *mockViaCEPService) GetLocationByCEP(cep string) (*models.ViaCEPResponse, error) {
	if cep == "12345678" {
		return &models.ViaCEPResponse{
			CEP:        "12345-678",
			Localidade: "São Paulo",
			UF:         "SP",
		}, nil
	}
	return nil, services.ErrCEPNotFound
}

// Mock Weather service
type mockWeatherService struct{}

func (m *mockWeatherService) GetWeatherByCity(city string) (*models.WeatherAPIResponse, error) {
	if city == "São Paulo" {
		response := &models.WeatherAPIResponse{}
		response.Current.TempC = 25.0
		return response, nil
	}
	return nil, services.ErrCityNotFound
}

func TestHandleWeatherByCEP(t *testing.T) {
	// Create services with mocks
	viaCEPService := &mockViaCEPService{}
	weatherService := &mockWeatherService{}
	tempConverter := services.NewTemperatureConverter()

	// Create handlers with mock services
	handlers := NewHandlers(viaCEPService, weatherService, tempConverter)

	tests := []struct {
		name           string
		cep            string
		expectedStatus int
		checkBody      bool
	}{
		{"Valid CEP", "12345678", http.StatusOK, true},
		{"Invalid CEP format", "123", http.StatusUnprocessableEntity, false},
		{"CEP not found", "87654321", http.StatusNotFound, false},
		{"Missing CEP parameter", "", http.StatusBadRequest, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request with the CEP parameter
			req, err := http.NewRequest("GET", "/weather", nil)
			if err != nil {
				t.Fatal(err)
			}

			// Add query parameter if it's not the "Missing CEP parameter" test
			if tt.cep != "" {
				q := req.URL.Query()
				q.Add("cep", tt.cep)
				req.URL.RawQuery = q.Encode()
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Call the handler
			handlers.HandleWeatherByCEP(rr, req)

			// Check the status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Check the response body for the success case
			if tt.checkBody && tt.expectedStatus == http.StatusOK {
				var response models.WeatherResponse
				err := json.NewDecoder(rr.Body).Decode(&response)
				if err != nil {
					t.Errorf("Error decoding response body: %v", err)
				}

				// Expected values for São Paulo with tempC = 25.0
				expectedF := 77.0 // 25 * 1.8 + 32
				expectedK := 298.0 // 25 + 273

				// Check temperature values with some tolerance
				if !floatEquals(response.TempC, 25.0, 0.001) {
					t.Errorf("Expected TempC to be 25.0, got %f", response.TempC)
				}

				if !floatEquals(response.TempF, expectedF, 0.001) {
					t.Errorf("Expected TempF to be %f, got %f", expectedF, response.TempF)
				}

				if !floatEquals(response.TempK, expectedK, 0.001) {
					t.Errorf("Expected TempK to be %f, got %f", expectedK, response.TempK)
				}
			}
		})
	}
}

// floatEquals checks if two float values are equal within a certain tolerance
func floatEquals(a, b, tolerance float64) bool {
	if a < b {
		return b-a < tolerance
	}
	return a-b < tolerance
}