package showservice

import (
	"context"
	"testing"
	"time"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	repomock "github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/repoaitorymock"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/contextkey"
	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/roles"
)


// unit test for all showservice functions
func TestCreateShow(t *testing.T) {
	now := time.Now()
	validShow := &model.Show{
		MovieId:    "movie1",
		StartTime:  now.Add(time.Hour),
		EndTime:    now.Add(2 * time.Hour),
		TotalSeats: 100,
	}

	tests := []struct {
		name           string
		ctx            context.Context
		show           *model.Show
		mockSetup      func(*repomock.MockShowRepo)
		expectedErrMsg string
	}{
		{
			name:           "unauthorized user",
			ctx:            context.WithValue(context.Background(), contextkey.UserRole, role.Customer),
			show:           validShow,
			mockSetup:      func(m *repomock.MockShowRepo) {},
			expectedErrMsg: "unauthoe=rized for adding the show",
		},
		{
			name:           "start time in past",
			ctx:            context.WithValue(context.Background(), contextkey.UserRole, role.Admin),
			show:           &model.Show{StartTime: now.Add(-time.Hour), EndTime: now.Add(time.Hour)},
			mockSetup:      func(m *repomock.MockShowRepo) {},
			expectedErrMsg: "enter valid date start time cannot be in the past or endtime should be greater the start time",
		},
		{
			name:           "end time before start time",
			ctx:            context.WithValue(context.Background(), contextkey.UserRole, role.Admin),
			show:           &model.Show{StartTime: now.Add(time.Hour), EndTime: now},
			mockSetup:      func(m *repomock.MockShowRepo) {},
			expectedErrMsg: "enter valid date start time cannot be in the past or endtime should be greater the start time",
		},
		{
			name:  "conflicting show times",
			ctx:   context.WithValue(context.Background(), contextkey.UserRole, role.Admin),
			show:  validShow,
			mockSetup: func(m *repomock.MockShowRepo) {
				m.Conflict = true
			},
			expectedErrMsg: "show timings conflicting with other show",
		},
		{
			name:  "conflict check repo error",
			ctx:   context.WithValue(context.Background(), contextkey.UserRole, role.Admin),
			show:  validShow,
			mockSetup: func(m *repomock.MockShowRepo) {
				m.IsConflictError = true
			},
			expectedErrMsg: "show timings conflicting with other show",
		},
		{
			name:  "repo create show error",
			ctx:   context.WithValue(context.Background(), contextkey.UserRole, role.Admin),
			show:  validShow,
			mockSetup: func(m *repomock.MockShowRepo) {
				m.ShouldError = true
			},
			expectedErrMsg: "jsdsjd",
		},
		{
			name:  "successful create show",
			ctx:   context.WithValue(context.Background(), contextkey.UserRole, role.Admin),
			show:  validShow,
			mockSetup: func(m *repomock.MockShowRepo) {
				m.ShouldError = false
			},
			expectedErrMsg: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repomock.NewMokeShowRepo()
			tt.mockSetup(mockRepo)
			service := NewShowService(mockRepo)

			showId, err := service.CreateShow(tt.ctx, tt.show)
			if tt.expectedErrMsg != "" {
				if err == nil || err.Error() != tt.expectedErrMsg {
					t.Errorf("expected error %q, got %v", tt.expectedErrMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if showId == "" {
					t.Errorf("expected showId to be set, got empty string")
				}
			}
		})
	}
}

func TestGetAllShow(t *testing.T) {
	tests := []struct {
		name           string
		mockSetup      func(*repomock.MockShowRepo)
		expectedErrMsg string
		expectedCount  int
	}{
		{
			name: "repo error",
			mockSetup: func(m *repomock.MockShowRepo) {
				m.ShouldError = true
			},
			expectedErrMsg: "sjds",
			expectedCount:  0,
		},
		{
			name: "no shows available",
			mockSetup: func(m *repomock.MockShowRepo) {
				m.ShouldError = false
				m.HasShows = false
			},
			expectedErrMsg: "no show available",
			expectedCount:  0,
		},
		{
			name: "shows available",
			mockSetup: func(m *repomock.MockShowRepo) {
				m.ShouldError = false
				m.HasShows = true
			},
			expectedErrMsg: "",
			expectedCount:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repomock.NewMokeShowRepo()
			tt.mockSetup(mockRepo)
			service := NewShowService(mockRepo)

			shows, err := service.GetAllShow()
			if tt.expectedErrMsg != "" {
				if err == nil || err.Error() != tt.expectedErrMsg {
					t.Errorf("expected error %q, got %v", tt.expectedErrMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if len(shows) != tt.expectedCount {
					t.Errorf("expected %d shows, got %d", tt.expectedCount, len(shows))
				}
			}
		})
	}
}

func TestGetShowByMovieId(t *testing.T) {
	tests := []struct {
		name           string
		movieId        string
		mockSetup      func(*repomock.MockShowRepo)
		expectedErrMsg string
		expectedCount  int
	}{
		{
			name:    "repo error",
			movieId: "movie1",
			mockSetup: func(m *repomock.MockShowRepo) {
				m.ShouldError = true
			},
			expectedErrMsg: "sdjs",
			expectedCount:  0,
		},
		{
			name:    "no shows for movie",
			movieId: "movie1",
			mockSetup: func(m *repomock.MockShowRepo) {
				m.ShouldError = false
				m.HasShows = false
			},
			expectedErrMsg: "no show available for this movie",
			expectedCount:  0,
		},
		{
			name:    "shows available for movie",
			movieId: "movie1",
			mockSetup: func(m *repomock.MockShowRepo) {
				m.ShouldError = false
				m.HasShows = true
			},
			expectedErrMsg: "",
			expectedCount:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := repomock.NewMokeShowRepo()
			tt.mockSetup(mockRepo)
			service := NewShowService(mockRepo)

			shows, err := service.GetShowByMovieId(tt.movieId)
			if tt.expectedErrMsg != "" {
				if err == nil || err.Error() != tt.expectedErrMsg {
					t.Errorf("expected error %q, got %v", tt.expectedErrMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if len(shows) != tt.expectedCount {
					t.Errorf("expected %d shows, got %d", tt.expectedCount, len(shows))
				}
			}
		})
	}
}
