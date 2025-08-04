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
	ErrCEPNotFound      = errors.New("can not find zipcode")
	ErrInvalidCEP       = errors.New("invalid zipcode")
	ErrAPIRequest       = errors.New("error in API request")
	ErrUnmarshalResponse = errors.New("error unmarshalling response")
)

// ViaCEPService handles interactions with the ViaCEP API
type ViaCEPService struct {
	client *http.Client
}

// NewViaCEPService creates a new ViaCEP service
func NewViaCEPService() *ViaCEPService {
	return &ViaCEPService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetLocationByCEP retrieves location information by CEP
func (s *ViaCEPService) GetLocationByCEP(cep string) (*models.ViaCEPResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrAPIRequest, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound || resp.StatusCode >= 400 {
		return nil, ErrCEPNotFound
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrAPIRequest, err)
	}

	var cepResponse models.ViaCEPResponse
	if err := json.Unmarshal(body, &cepResponse); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrUnmarshalResponse, err)
	}

	// Check if the response contains an error
	if cepResponse.Erro {
		return nil, ErrCEPNotFound
	}
	
	// Print debug information
	fmt.Printf("ViaCEP Response for %s: City=%s\n", cep, cepResponse.Localidade)

	return &cepResponse, nil
}