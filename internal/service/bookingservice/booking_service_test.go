
package bookingservice

import (
	"context"
	"errors"
	"testing"
	"time"

	repomock "github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/repoaitorymock"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/contextkey"
)

// unit test for booking service
func TestBookTicket(t *testing.T) {
	const validShowId = "hdsdshd"
	const validUserId = "skdjshd"

	show := &model.Show{
		ShowId:        validShowId,
		AvailableSeat: 15,
	}

	tests := []struct {
		name             string
		showId           string
		userId           string
		numberOfTickets  int
		mockTicketRepo   *repomock.MockTicketRepo
		mockShowRepo     *repomock.MockShowRepo
		wantError        error
		expectSuccess    bool
	}{
		{
			name:            "show not present",
			showId:          "invalid",
			userId:          validUserId,
			numberOfTickets: 2,
			mockTicketRepo:  &repomock.MockTicketRepo{},
			mockShowRepo:    &repomock.MockShowRepo{ShowIdError: true},
			wantError:       errors.New("no show available"),
		},
		{
			name:            "invalid seats",
			showId:          validShowId,
			userId:          validUserId,
			numberOfTickets: 50,
			mockTicketRepo:  &repomock.MockTicketRepo{},
			mockShowRepo:    &repomock.MockShowRepo{Shows: map[string]*model.Show{validShowId: show}},
			wantError:       errors.New("required Number of seats are not available"),
		},
		{
			name:            "update show fails",
			showId:          validShowId,
			userId:          validUserId,
			numberOfTickets: 2,
			mockTicketRepo:  &repomock.MockTicketRepo{},
			mockShowRepo:    &repomock.MockShowRepo{Shows: map[string]*model.Show{validShowId: show}, UpdateShowerro: true},
			wantError:       errors.New("error in updating the show"),
		},
		{
			name:            "save ticket fails",
			showId:          validShowId,
			userId:          validUserId,
			numberOfTickets: 2,
			mockTicketRepo:  &repomock.MockTicketRepo{ShouldError: true},
			mockShowRepo:    &repomock.MockShowRepo{Shows: map[string]*model.Show{validShowId: show}},
			wantError:       errors.New("error in saving Tickets"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bookingService := NewBookingService(tt.mockTicketRepo, tt.mockShowRepo)

			tickets, err := bookingService.BookTicket(tt.showId, tt.userId, tt.numberOfTickets)

			if tt.wantError != nil {
				if err == nil || err.Error() != tt.wantError.Error() {
					t.Errorf("expected error %v, got %v", tt.wantError, err)
				}
			}

			if tt.expectSuccess {
				if len(tickets) != tt.numberOfTickets {
					t.Errorf("expected %d tickets, got %d", tt.numberOfTickets, len(tickets))
				}
				for _, ticket := range tickets {
					if ticket.ShowId != tt.showId || ticket.UserId != tt.userId {
						t.Errorf("ticket fields mismatch: %+v", ticket)
					}
				}
			}
		})
	}
}



func TestCancelTicket(t *testing.T) {
	const validTicketId = "ticket-123"
	const validUserId = "user-123"
	const validShowId = "show-123"

	ticket := &model.Ticket{
		TicketId:    validTicketId,
		ShowId:      validShowId,
		UserId:      validUserId,
		BookingTime: time.Now(),
	}
	show := &model.Show{
		ShowId:        validShowId,
		AvailableSeat: 10,
	}

	tests := []struct {
		name           string
		ctxUserId      string
		mockTicketRepo *repomock.MockTicketRepo
		mockShowRepo   *repomock.MockShowRepo
		wantError      error
	}{
		{
			name:           "ticket not found",
			ctxUserId:      validUserId,
			mockTicketRepo: &repomock.MockTicketRepo{ShouldError: true},
			mockShowRepo:   &repomock.MockShowRepo{},
			wantError:      errors.New("ticket is not available"),
		},
		{
			name:           "user not allowed to cancel",
			ctxUserId:      "other-user",
			mockTicketRepo: &repomock.MockTicketRepo{Tickets: map[string]*model.Ticket{validTicketId: ticket}},
			mockShowRepo:   &repomock.MockShowRepo{Shows: map[string]*model.Show{validShowId: show}},
			wantError:      errors.New("you are not allowed to cancel this ticket"),
		},
		{
			name:           "cancel ticket fails",
			ctxUserId:      validUserId,
			mockTicketRepo: &repomock.MockTicketRepo{Tickets: map[string]*model.Ticket{validTicketId: ticket}, CancelError: true},
			mockShowRepo:   &repomock.MockShowRepo{Shows: map[string]*model.Show{validShowId: show}},
			wantError:      errors.New("enable to cancel ticket"),
		},
		{
			name:           "update show fails",
			ctxUserId:      validUserId,
			mockTicketRepo: &repomock.MockTicketRepo{Tickets: map[string]*model.Ticket{validTicketId: ticket}},
			mockShowRepo:   &repomock.MockShowRepo{Shows: map[string]*model.Show{validShowId: show}, UpdateShowerro: true},
			wantError:      errors.New("dhfhd"),
		},
		{
			name:           "successful cancel",
			ctxUserId:      validUserId,
			mockTicketRepo: &repomock.MockTicketRepo{Tickets: map[string]*model.Ticket{validTicketId: ticket}},
			mockShowRepo:   &repomock.MockShowRepo{Shows: map[string]*model.Show{validShowId: show}},
			wantError:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.WithValue(context.Background(), contextkey.UserId, tt.ctxUserId)
			bookingService := NewBookingService(tt.mockTicketRepo, tt.mockShowRepo)

			err := bookingService.CancelTicket(ctx, validTicketId)

			if tt.wantError == nil {
				if err != nil {
					t.Errorf("expected success, got error: %v", err)
				}
			} else {
				if err == nil || err.Error() != tt.wantError.Error() {
					t.Errorf("expected error %v, got %v", tt.wantError, err)
				}
			}
		})
	}
}
