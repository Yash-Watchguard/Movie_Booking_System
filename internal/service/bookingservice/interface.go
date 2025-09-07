package bookingservice

import (
	"context"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type BookingServiceInterface interface {
	BookTicket(showId string, userId string, numberOfTickets int) ([]model.Ticket, error)
	CancelTicket(ctx context.Context, ticketId string) error
}