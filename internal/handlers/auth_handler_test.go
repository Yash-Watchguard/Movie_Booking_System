package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/servicemock"
)

func TestAuthSignup(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        string
		expectedStatusCode int
		expectedMessage    string
		mockAuthService    *servicemock.MockAuthService
		shouldSucceed      bool
	}{
		{
			name:               "invalid email",
			requestBody:        `{"name":"yash","email":"yayys","phone_number":"8372328123","password":"jskjds23"}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "Invalid Email",
			mockAuthService:    &servicemock.MockAuthService{},
			shouldSucceed:      false,
		},
		{
			name:               "invalid phone number",
			requestBody:        `{"name":"yash","email":"yashgoyal@gmail.com","phone_number":"8372328","password":"jskjds23"}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "Invalid phonenumber",
			mockAuthService:    &servicemock.MockAuthService{},
			shouldSucceed:      false,
		},
		{
			name:               "signup service error",
			// Use a valid phone number so the handler proceeds to call the service
			requestBody:        `{"name":"yash","email":"yash@example.com","phone_number":"8372328123","password":"jskjds23"}`,
			expectedStatusCode: http.StatusInternalServerError,
			expectedMessage:    "Signup Failed",
			mockAuthService: &servicemock.MockAuthService{
				ShouldError: true,
			},
			shouldSucceed: false,
		},
		{
			name:               "signup success",
			// Corrected the phone number to be valid
			requestBody:        `{"name":"yash","email":"yash@example.com","phone_number":"8372328123","password":"jskjds23"}`,
			expectedStatusCode: http.StatusCreated,
			expectedMessage:    "User Created Successfully",
			mockAuthService: &servicemock.MockAuthService{
				ShouldError: false,
			},
			shouldSucceed: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/v1/signup", strings.NewReader(tc.requestBody))
			w := httptest.NewRecorder()

			authHandler := NewAuthHandler(tc.mockAuthService)
			authHandler.SignUp(w, req)

			if w.Code != tc.expectedStatusCode {
				t.Fatalf("expected status %d, got %d", tc.expectedStatusCode, w.Code)
			}
			
			if tc.shouldSucceed {
				var got Successresponse
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal success response: %v", err)
				}
				if got.Message != tc.expectedMessage {
					t.Errorf("expected message %q, got %q", tc.expectedMessage, got.Message)
				}
			} else {
				var got ErrorResponse
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal error response: %v", err)
				}
				if got.Message != tc.expectedMessage {
					t.Errorf("expected message %q, got %q", tc.expectedMessage, got.Message)
				}
			}
		})
	}
}

func TestAuthLogin(t *testing.T) {
	tests := []struct {
		name               string
		requestBody        string
		expectedStatusCode int
		expectedMessage    string
		mockAuthService    *servicemock.MockAuthService
		shouldSucceed      bool
	}{
	
		{
			name:               "invalid email format",
			requestBody:        `{"email":"invalid-email","password":"password123"}`,
			expectedStatusCode: http.StatusBadRequest,
			expectedMessage:    "Invalid Email",
			mockAuthService:    &servicemock.MockAuthService{},
			shouldSucceed:      false,
		},
		{
			name:               "login service error",
			requestBody:        `{"email":"yash@example.com","password":"jskjds23"}`,
			expectedStatusCode: http.StatusInternalServerError,
			expectedMessage:    "login failed",
			mockAuthService: &servicemock.MockAuthService{
				ShouldError: true,
			},
			shouldSucceed: false,
		},
	
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(tc.requestBody))
			w := httptest.NewRecorder()

			authHandler := NewAuthHandler(tc.mockAuthService)
			authHandler.Login(w, req)

			if w.Code != tc.expectedStatusCode {
				t.Fatalf("expected status %d, got %d", tc.expectedStatusCode, w.Code)
			}

			// This is a simplified check that works for both success and error responses
			var got struct {
				Message string `json:"message"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}
			if got.Message != tc.expectedMessage {
				t.Errorf("expected message %q, got %q", tc.expectedMessage, got.Message)
			}
		})
	}
}
