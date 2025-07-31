package services

import (
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	converter := NewTemperatureConverter()

	tests := []struct {
		name      string
		tempC     float64
		expectedF float64
		expectedK float64
	}{
		{"Zero Celsius", 0.0, 32.0, 273.0},
		{"Room temperature", 25.0, 77.0, 298.0},
		{"Boiling point", 100.0, 212.0, 373.0},
		{"Negative temperature", -10.0, 14.0, 263.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := converter.ConvertTemperature(tt.tempC)

			// Check Celsius value (should be the same as input)
			if result.TempC != tt.tempC {
				t.Errorf("Expected Celsius temperature to be %f, got %f", tt.tempC, result.TempC)
			}

			// Check Fahrenheit conversion with small tolerance for floating point errors
			if !floatEquals(result.TempF, tt.expectedF, 0.001) {
				t.Errorf("Fahrenheit conversion: got %f, expected %f", result.TempF, tt.expectedF)
			}

			// Check Kelvin conversion with small tolerance for floating point errors
			if !floatEquals(result.TempK, tt.expectedK, 0.001) {
				t.Errorf("Kelvin conversion: got %f, expected %f", result.TempK, tt.expectedK)
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