package zipcode

import (
	"luizalabs-technical-test/pkg/constants/str"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestViaCepResponseToGetAddressByZipCodeResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    ViaCepResponse
		expected *GetAddressByZipCodeUnifiedResponse
		wantErr  bool
	}{
		{
			name: "Valid response from ViaCep",
			input: ViaCepResponse{
				Logradouro: "Main St",
				Bairro:     "Downtown",
				Localidade: "Cityville",
				Uf:         "ST",
			},
			expected: &GetAddressByZipCodeUnifiedResponse{
				Street:       "Main St",
				Neighborhood: "Downtown",
				City:         "Cityville",
				State:        "ST",
			},
			wantErr: false,
		},
		{
			name: "Empty response from ViaCep",
			input: ViaCepResponse{
				Logradouro: str.EmptyString,
				Bairro:     str.EmptyString,
				Localidade: str.EmptyString,
				Uf:         str.EmptyString,
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Partial valid response from ViaCep",
			input: ViaCepResponse{
				Logradouro: "Main St",
				Bairro:     str.EmptyString,
				Localidade: "Cityville",
				Uf:         "ST",
			},
			expected: &GetAddressByZipCodeUnifiedResponse{
				Street:       "Main St",
				Neighborhood: str.EmptyString,
				City:         "Cityville",
				State:        "ST",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.ToGetAddressByZipCodeResponse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Expected error status %v, got %v (error: %v)", tt.wantErr, !tt.wantErr, err)
				return
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestOpenCepResponseToGetAddressByZipCodeResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    OpenCepResponse
		expected *GetAddressByZipCodeUnifiedResponse
		wantErr  bool
	}{
		{
			name: "Valid response from OpenCep",
			input: OpenCepResponse{
				Logradouro: "Broadway",
				Bairro:     "Central Park",
				Localidade: "New York",
				Uf:         "NY",
			},
			expected: &GetAddressByZipCodeUnifiedResponse{
				Street:       "Broadway",
				Neighborhood: "Central Park",
				City:         "New York",
				State:        "NY",
			},
			wantErr: false,
		},
		{
			name: "Empty response from OpenCep",
			input: OpenCepResponse{
				Logradouro: str.EmptyString,
				Bairro:     str.EmptyString,
				Localidade: str.EmptyString,
				Uf:         str.EmptyString,
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Partial valid response from OpenCep",
			input: OpenCepResponse{
				Logradouro: "Broadway",
				Bairro:     str.EmptyString,
				Localidade: "New York",
				Uf:         "NY",
			},
			expected: &GetAddressByZipCodeUnifiedResponse{
				Street:       "Broadway",
				Neighborhood: str.EmptyString,
				City:         "New York",
				State:        "NY",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.ToGetAddressByZipCodeResponse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Expected error status %v, got %v (error: %v)", tt.wantErr, !tt.wantErr, err)
				return
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBrasilAPIResponseToGetAddressByZipCodeResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    BrasilAPIResponse
		expected *GetAddressByZipCodeUnifiedResponse
		wantErr  bool
	}{
		{
			name: "Valid response from Brasil API",
			input: BrasilAPIResponse{
				Street:       "Av. Paulista",
				Neighborhood: "Bela Vista",
				City:         "São Paulo",
				State:        "SP",
			},
			expected: &GetAddressByZipCodeUnifiedResponse{
				Street:       "Av. Paulista",
				Neighborhood: "Bela Vista",
				City:         "São Paulo",
				State:        "SP",
			},
			wantErr: false,
		},
		{
			name: "Empty response from Brasil API",
			input: BrasilAPIResponse{
				Street:       str.EmptyString,
				Neighborhood: str.EmptyString,
				City:         str.EmptyString,
				State:        str.EmptyString,
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Partial valid response from Brasil API",
			input: BrasilAPIResponse{
				Street:       "Av. Paulista",
				Neighborhood: str.EmptyString,
				City:         "São Paulo",
				State:        "SP",
			},
			expected: &GetAddressByZipCodeUnifiedResponse{
				Street:       "Av. Paulista",
				Neighborhood: str.EmptyString,
				City:         "São Paulo",
				State:        "SP",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.ToGetAddressByZipCodeResponse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Expected error status %v, got %v (error: %v)", tt.wantErr, !tt.wantErr, err)
				return
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestAPICepResponseToGetAddressByZipCodeResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    APICepResponse
		expected *GetAddressByZipCodeUnifiedResponse
		wantErr  bool
	}{
		{
			name: "Valid response from API Cep",
			input: APICepResponse{
				Address:  "Rua XV de Novembro",
				District: "Centro",
				City:     "Curitiba",
				State:    "PR",
			},
			expected: &GetAddressByZipCodeUnifiedResponse{
				Street:       "Rua XV de Novembro",
				Neighborhood: "Centro",
				City:         "Curitiba",
				State:        "PR",
			},
			wantErr: false,
		},
		{
			name: "Empty response from API Cep",
			input: APICepResponse{
				Address:  str.EmptyString,
				District: str.EmptyString,
				City:     str.EmptyString,
				State:    str.EmptyString,
			},
			expected: nil,
			wantErr:  true,
		},
		{
			name: "Partial valid response from API Cep",
			input: APICepResponse{
				Address:  "Rua XV de Novembro",
				District: str.EmptyString,
				City:     "Curitiba",
				State:    "PR",
			},
			expected: &GetAddressByZipCodeUnifiedResponse{
				Street:       "Rua XV de Novembro",
				Neighborhood: str.EmptyString,
				City:         "Curitiba",
				State:        "PR",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.ToGetAddressByZipCodeResponse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Expected error status %v, got %v (error: %v)", tt.wantErr, !tt.wantErr, err)
				return
			}
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestGetAddressByZipCodeUnifiedResponseToGetAddressByZipCodeResponse(t *testing.T) {
	tests := []struct {
		name     string
		input    GetAddressByZipCodeUnifiedResponse
		expected GetAddressByZipCodeResponse
	}{
		{
			name: "Convert unified response to wrapper",
			input: GetAddressByZipCodeUnifiedResponse{
				Street:       "Rua das Flores",
				Neighborhood: "Jardim",
				City:         "Rio de Janeiro",
				State:        "RJ",
			},
			expected: GetAddressByZipCodeResponse{
				GetAddressByZipCodeUnifiedResponse: GetAddressByZipCodeUnifiedResponse{
					Street:       "Rua das Flores",
					Neighborhood: "Jardim",
					City:         "Rio de Janeiro",
					State:        "RJ",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.ToGetAddressByZipCodeResponse()
			assert.Equal(t, tt.expected, result)
		})
	}
}
