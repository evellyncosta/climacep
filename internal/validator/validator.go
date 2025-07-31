package validator

import (
	"regexp"
	"strings"
)

// ValidateCEP checks if a CEP is valid
func ValidateCEP(cep string) bool {
	// Remove non-digit characters
	cep = strings.ReplaceAll(cep, "-", "")
	cep = strings.ReplaceAll(cep, ".", "")
	
	// Check if it has exactly 8 digits
	if len(cep) != 8 {
		return false
	}
	
	// Check if it contains only digits
	match, _ := regexp.MatchString(`^\d{8}$`, cep)
	return match
}

// FormatCEP formats a CEP to the standard format
func FormatCEP(cep string) string {
	// Remove non-digit characters
	cep = strings.ReplaceAll(cep, "-", "")
	cep = strings.ReplaceAll(cep, ".", "")
	
	return cep
}