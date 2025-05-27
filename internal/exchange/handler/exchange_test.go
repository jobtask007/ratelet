package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"ratelet/internal/exchange/handler/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestRates_GetRates(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		url            string
		mockSetup      func(s *mocks.MockService)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Success",
			url:  "/exchange?from=WBTC&to=USDT&amount=1.0",
			mockSetup: func(s *mocks.MockService) {
				s.On("Exchange", mock.Anything, mock.Anything, mock.Anything).
					Return("57613.353535", nil).Once()
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"from": "WBTC", "to": "USDT", "amount": "57613.353535",
			},
		},
		{
			name: "Fails on missing 'from' param",
			url:  "/exchange?to=USDT&amount=1.0",
			mockSetup: func(s *mocks.MockService) {
				s.On("Exchange", mock.Anything, mock.Anything, mock.Anything).
					Return(mock.Anything, nil).Maybe()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Fails on missing 'to' param",
			url:  "/exchange?from=WBTC&amount=1.0",
			mockSetup: func(s *mocks.MockService) {
				s.On("Exchange", mock.Anything, mock.Anything, mock.Anything).
					Return(mock.Anything, nil).Maybe()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Fails on missing 'amount' param",
			url:  "/exchange?from=WBTC&to=USDT",
			mockSetup: func(s *mocks.MockService) {
				s.On("Exchange", mock.Anything, mock.Anything, mock.Anything).
					Return(mock.Anything, nil).Maybe()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Fails on parse 'amount' param",
			url:  "/exchange?from=WBTC&to=USDT&amount=1.0x",
			mockSetup: func(s *mocks.MockService) {
				s.On("Exchange", mock.Anything, mock.Anything, mock.Anything).
					Return(mock.Anything, nil).Maybe()
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
		{
			name: "Fails on service error",
			url:  "/exchange?from=WBTC&to=USDT&amount=1.0",
			mockSetup: func(s *mocks.MockService) {
				s.On("Exchange", mock.Anything, mock.Anything, mock.Anything).
					Return("", errors.New("unexpected error")).Maybe()
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
			handler.Exchange(c)

			// then
			assert.Equal(t, tt.expectedStatus, rr.Code)

			var actual map[string]interface{}

			if rr.Body.Len() > 0 {
				err := json.Unmarshal([]byte(rr.Body.String()), &actual)
				require.NoError(t, err)
			}

			assert.Equal(t, tt.expectedBody, actual)

			mockSvc.AssertExpectations(t)
		})
	}
}
