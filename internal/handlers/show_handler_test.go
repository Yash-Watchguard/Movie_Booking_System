package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/servicemock"
)

func TestShowHandler_CreateShow(t *testing.T) {
	validShow := map[string]interface{}{
		"movie_id":        "movie123",
		"start_time":      time.Now().Format(time.RFC3339),
		"end_time":        time.Now().Add(time.Hour).Format(time.RFC3339),
		"total_seats":     100,
		"available_seats": 100,
	}

	tests := []struct {
		name           string
		body           map[string]interface{}
		shouldErr      bool
		expectedStatus int
	}{
		{
			name:           "valid request",
			body:           validShow,
			shouldErr:      false,
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "invalid json",
			body:           nil,
			shouldErr:      false,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid start time",
			body: map[string]interface{}{
				"movie_id":        "movie123",
				"start_time":      "invalid-time",
				"end_time":        time.Now().Add(time.Hour).Format(time.RFC3339),
				"total_seats":     100,
				"available_seats": 100,
			},
			shouldErr:      false,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "invalid end time",
			body: map[string]interface{}{
				"movie_id":        "movie123",
				"start_time":      time.Now().Format(time.RFC3339),
				"end_time":        "invalid-time",
				"total_seats":     100,
				"available_seats": 100,
			},
			shouldErr:      false,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "service error",
			body:           validShow,
			shouldErr:      true,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockShowService := &servicemock.MockShowService{Shoulderr: tt.shouldErr}
			handler := NewShowHandler(mockShowService)

			var reqBody *bytes.Buffer
			if tt.body != nil {
				b, _ := json.Marshal(tt.body)
				reqBody = bytes.NewBuffer(b)
			} else {
				reqBody = bytes.NewBuffer([]byte("invalid json"))
			}

			req := httptest.NewRequest(http.MethodPost, "/shows", reqBody)
			w := httptest.NewRecorder()

			handler.CreateShow(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
		})
	}
}


func TestGetAllShow(t *testing.T) {
	tests := []struct {
		name           string
		shouldErr      bool
		expectedStatus int
	}{
		{
			name:           "success",
			shouldErr:      false,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "service error",
			shouldErr:      true,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockShowService := &servicemock.MockShowService{Shoulderr: tt.shouldErr}
			handler := NewShowHandler(mockShowService)

			req := httptest.NewRequest(http.MethodGet, "/shows", nil)
			w := httptest.NewRecorder()

			handler.GetAllShow(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
		})
	}
}

func TestGetAllShowofMovie(t *testing.T) {
	tests := []struct {
		name           string
		movieId        string
		shouldErr      bool
		expectedStatus int
	}{
		{
			name:           "success",
			movieId:        "movie1",
			shouldErr:      false,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "missing movie id",
			movieId:        "",
			shouldErr:      false,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "service error",
			movieId:        "movie1",
			shouldErr:      true,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockShowService := &servicemock.MockShowService{Shoulderr: tt.shouldErr}
			handler := NewShowHandler(mockShowService)

			req := httptest.NewRequest(http.MethodGet, "/shows/movie/"+tt.movieId, nil)
			req = req.WithContext(context.WithValue(req.Context(), "movie_id", tt.movieId))
			w := httptest.NewRecorder()

			handler.GetAllShowofMovie(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
		})
	}
}
