package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"ratelet/internal/rate/domain"
	"ratelet/internal/rate/handler/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRates_GetRates(t *testing.T) {
	gin.SetMode(gin.TestMode)

	sampleRates := domain.Rates{
		Rates: []domain.Rate{
			{From: "USD", To: "PLN", Rate: 3.85},
			{From: "PLN", To: "USD", Rate: 0.26},
		},
	}

	tests := []struct {
		name           string
		url            string
		mockSetup      func(s *mocks.MockService)
		expectedStatus int
		expectedBody   []map[string]interface{}
	}{
		{
			name: "Success",
			url:  "/rates?currencies=USD,PLN",
			mockSetup: func(s *mocks.MockService) {
				s.On("GetRates", mock.Anything).Return(sampleRates, nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedBody: []map[string]interface{}{
				{"from": "USD", "to": "PLN", "rate": 3.85},
				{"from": "PLN", "to": "USD", "rate": 0.26},
			},
		},
		{
			name: "Fails on missing currencies param",
			url:  "/rates",
			mockSetup: func(s *mocks.MockService) {
				s.On("GetRates", mock.Anything).
					Return(mock.Anything, mock.Anything).Maybe()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Fails on invalid currencies param",
			url:  "/rates?currencies=USD",
			mockSetup: func(s *mocks.MockService) {
				s.On("GetRates", mock.Anything).
					Return(mock.Anything, mock.Anything).Maybe()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Fails on service error",
			url:  "/rates?currencies=USD,PLN",
			mockSetup: func(s *mocks.MockService) {
				s.On("GetRates", mock.Anything).
					Return(domain.Rates{}, errors.New("unexpected error")).Once()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// given
			mockSvc := new(mocks.MockService)
			handler := New(mockSvc)

			if tt.mockSetup != nil {
				tt.mockSetup(mockSvc)
			}

			rr := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rr)

			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			c.Request = req

			// when
			handler.GetRates(c)

			// then
			assert.Equal(t, tt.expectedStatus, rr.Code)

			var actual []map[string]interface{}

			if rr.Body.Len() > 0 {
				err := json.Unmarshal([]byte(rr.Body.String()), &actual)
				require.NoError(t, err)
			}

			assert.Equal(t, tt.expectedBody, actual)

			mockSvc.AssertExpectations(t)
		})
	}
}
