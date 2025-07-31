package validator

import (
	"testing"
)

func TestValidateCEP(t *testing.T) {
	tests := []struct {
		name     string
		cep      string
		expected bool
	}{
		{"Valid CEP", "12345678", true},
		{"Valid formatted CEP", "12345-678", true},
		{"Invalid CEP (too short)", "1234567", false},
		{"Invalid CEP (too long)", "123456789", false},
		{"Invalid CEP (non-digits)", "abcdefgh", false},
		{"Invalid CEP (mixed)", "1234abcd", false},
		{"Empty CEP", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateCEP(tt.cep)
			if result != tt.expected {
				t.Errorf("ValidateCEP(%s) = %v, expected %v", tt.cep, result, tt.expected)
			}
		})
	}
}

func TestFormatCEP(t *testing.T) {
	tests := []struct {
		name     string
		cep      string
		expected string
	}{
		{"Already formatted", "12345678", "12345678"},
		{"With hyphen", "12345-678", "12345678"},
		{"With dots", "12.345.678", "12345678"},
		{"Mixed format", "12.345-678", "12345678"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FormatCEP(tt.cep)
			if result != tt.expected {
				t.Errorf("FormatCEP(%s) = %s, expected %s", tt.cep, result, tt.expected)
			}
		})
	}
}