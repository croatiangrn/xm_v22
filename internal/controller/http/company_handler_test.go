package http

import (
	"bytes"
	"encoding/json"
	customErrors "github.com/croatiangrn/xm_v22/internal/pkg/errors"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/croatiangrn/xm_v22/internal/controller/http/dto"
	"github.com/croatiangrn/xm_v22/internal/usecase/company/mocks"
)

func TestCompanyCreate(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// Create a mock use case
	mockUC := new(mocks.UseCase)
	var buf bytes.Buffer
	zeroLog := zerolog.New(&buf).With().Timestamp().Logger()
	
	h := NewCompanyHandler(mockUC, zeroLog)

	// Create a Gin router and register the handler
	router := gin.Default()
	router.POST("/companies", h.CompanyCreate)

	mockUUID, _ := uuid.Parse("00000000-0000-0000-0000-000000000000")

	// Test cases
	tests := []struct {
		name           string
		requestBody    interface{} // Use interface{} to allow invalid JSON
		mockResponse   *dto.CompanyResponse
		mockError      error
		expectedStatus int
		expectedBody   interface{} // Use interface{} to handle both *company.Company and gin.H
	}{
		{
			name: "Success - Company created",
			requestBody: dto.CreateCompanyRequest{
				Name:              "Test Company",
				Description:       "Test description",
				AmountOfEmployees: 10,
				Registered:        false,
				Type:              "Corporations",
			},
			mockResponse: &dto.CompanyResponse{
				ID:                mockUUID.String(),
				Name:              "Test Company",
				Description:       "Test description",
				AmountOfEmployees: 10,
				Registered:        false,
				Type:              "Corporations",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			expectedBody: &dto.CompanyResponse{
				Name:              "Test Company",
				Description:       "Test description",
				AmountOfEmployees: 10,
				Registered:        false,
				Type:              "Corporations",
			},
		},
		{
			name: "Amount of employees cannot be negative",
			requestBody: dto.CreateCompanyRequest{
				Name:              "Test Company",
				Description:       "Test description",
				AmountOfEmployees: -1,
				Registered:        false,
				Type:              "Corporations",
			},
			mockResponse:   nil,
			mockError:      customErrors.NewBadRequestError("amount_of_employees", "amount of employees cannot be negative"),
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "validation error on field amount_of_employees: amount of employees cannot be negative",
			},
		},
		{
			name:           "Error - Invalid JSON",
			requestBody:    "{description",
			mockResponse:   nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   map[string]interface{}{"error": "invalid character 'd' looking for beginning of object key string"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the use case response (only for successful cases)
			if tt.mockResponse != nil || tt.mockError != nil {
				mockUC.On("CreateCompany", mock.Anything, tt.requestBody).Return(tt.mockResponse, tt.mockError)
			}

			// Create the request body
			var requestBodyBytes []byte
			switch body := tt.requestBody.(type) {
			case string:
				// Use the string directly as the request body (for invalid JSON)
				requestBodyBytes = []byte(body)
			default:
				// Marshal the request body into JSON
				var err error
				requestBodyBytes, err = json.Marshal(body)
				assert.NoError(t, err)
			}

			// Create the HTTP request
			req, err := http.NewRequest(http.MethodPost, "/companies", bytes.NewBuffer(requestBodyBytes))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			// Create a response recorder
			rr := httptest.NewRecorder()

			// Perform the request
			router.ServeHTTP(rr, req)

			// Assert the status code
			assert.Equal(t, tt.expectedStatus, rr.Code)

			// Assert the response body
			if tt.expectedBody != nil {
				var responseBody interface{}
				err := json.Unmarshal(rr.Body.Bytes(), &responseBody)
				assert.NoError(t, err)

				// Handle different types of expectedBody
				switch expected := tt.expectedBody.(type) {
				case *dto.CompanyResponse:
					var actual dto.CompanyResponse
					err := json.Unmarshal(rr.Body.Bytes(), &actual)
					assert.NoError(t, err)
					assert.Equal(t, expected.Name, actual.Name)
					assert.Equal(t, expected.Description, actual.Description)
					assert.Equal(t, expected.AmountOfEmployees, actual.AmountOfEmployees)
					assert.Equal(t, expected.Registered, actual.Registered)
					assert.Equal(t, expected.Type, actual.Type)
				case map[string]interface{}:
					assert.Equal(t, expected, responseBody)
				default:
					log.Printf("Type: %T", expected)
					t.Fatalf("Unexpected type for expectedBody: %T", expected)
				}
			}

			// Assert that the mock was called as expected
			mockUC.AssertExpectations(t)
		})
	}
}
