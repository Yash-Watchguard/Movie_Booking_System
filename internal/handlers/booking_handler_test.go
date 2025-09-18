package handlers

import (
	"context"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/servicemock"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/models/contextkey"
)

type ErrorResponse struct {
	Message   string `json:"Message"`
	Status    string `json:"Status"`
	Errorcode int    `json:"Errorcode"`
}

type Successresponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func TestBookTicket(t *testing.T) {
	tests := []struct {
		status          bool
		name            string
		requestBody     string
		showId          string
		userId          string
		bookingService  *servicemock.MockBookingService
		expectedStatus  int
		expectedMessage string
	}{
		{
			name:            "missing show_id",
			requestBody:     `{"skdk":"jsdsjd"}`,
			showId:          "", 
			userId:          "user-1",
			bookingService:  &servicemock.MockBookingService{},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "Invalid request",
		},
		{
			name:            "invalid request body",
			requestBody:     `{"numberofseat": "notAnInt"}`, 
			showId:          "show-1",
			userId:          "user-1",
			bookingService:  &servicemock.MockBookingService{},
			expectedStatus:  http.StatusBadRequest,
			expectedMessage: "Invalid request body",
		},
		{
			name:        "invalid request body",
			requestBody: `{"numberofseat": 5}`,
			showId:      "show-1",
			userId:      "user-1",
			bookingService: &servicemock.MockBookingService{
				ShouldError: true,
			},
			expectedStatus:  http.StatusInternalServerError,
			expectedMessage: "Ticket booking failed",
		},
		{
			status:      true,
			name:        "invalid request body",
			requestBody: `{"numberofseat": 5}`,
			showId:      "show-1",
			userId:      "user-1",
			bookingService: &servicemock.MockBookingService{
				ShouldError: false,
			},
			expectedStatus:  http.StatusOK,
			expectedMessage: "Tickets Booked Successfully",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodPost, "/v1/booking/bookticket/"+tc.showId, strings.NewReader(tc.requestBody))
			req.SetPathValue("show_id", tc.showId)

			ctx := context.WithValue(context.Background(), contextkey.UserId, tc.userId)
			req = req.WithContext(ctx) 

			w := httptest.NewRecorder()

			bookingHandler := NewBookingHandler(tc.bookingService)
			
			bookingHandler.BookTicket(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Code)
			}

			if tc.status {
				var got Successresponse
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if got.Message != tc.expectedMessage {
					t.Errorf("expected message %q, got %q", tc.expectedMessage, got.Message)
				}
			} else {
				var got ErrorResponse
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}

				if got.Message != tc.expectedMessage {
					t.Errorf("expected message %q, got %q", tc.expectedMessage, got.Message)
				}
			}

		})
	}
}

func TestCancelTicket(t *testing.T) {
	tests := []struct {
		status          bool
		name            string
		ticketId        string
		userId          string
		bookingService  *servicemock.MockBookingService
		expectedStatus  int
		expectedMessage string
	}{
		{
			name:           "missing ticket_id",
			ticketId:       "",
			userId:         "user-1",
			bookingService: &servicemock.MockBookingService{},
			expectedStatus: http.StatusBadRequest,
			expectedMessage:"Invalid request",
		},
		
		{
			status:         true,
			name:           "cancel success",
			ticketId:       "ticket-1",
			userId:         "user-1",
			bookingService: &servicemock.MockBookingService{ShouldError: false},
			expectedStatus: http.StatusOK,
			expectedMessage: "ticket cancelled",
		},
		{
			status:         false,
			name:           "cancel fail",
			ticketId:       "ticket-1",
			userId:         "user-1",
			bookingService: &servicemock.MockBookingService{ShouldError: true},
			expectedStatus: http.StatusInternalServerError,
			expectedMessage: "Ticket Cancellation failed",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			
			req := httptest.NewRequest(http.MethodPost, "/v1/booking/cancel/"+tc.ticketId, nil)
			req.SetPathValue("ticket_id", tc.ticketId)

			ctx := context.WithValue(context.Background(), contextkey.UserId, tc.userId)
			req = req.WithContext(ctx)

			w := httptest.NewRecorder()

			bookingHandler := NewBookingHandler(tc.bookingService)

			
			bookingHandler.CancelTicket(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Code)
			}

			if tc.status {
				var got Successresponse
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}
				if got.Message != tc.expectedMessage {
					t.Errorf("expected message %q, got %q", tc.expectedMessage, got.Message)
				}
			} else {
				var got ErrorResponse
				if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil {
					t.Fatalf("failed to unmarshal response: %v", err)
				}
				if got.Message != tc.expectedMessage {
					t.Errorf("expected message %q, got %q", tc.expectedMessage, got.Message)
				}
			}
		})
	}
}
