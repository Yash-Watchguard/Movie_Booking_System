package servicemock

import (
	"context"
	"errors"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MockBookingService struct {
	ShouldError bool
}

func NewBookingService() *MockBookingService {
	return &MockBookingService{}
}


func (bookingService *MockBookingService) BookTicket(showId string, userId string, numberOfTickets int) ([]model.Ticket, error) {
	if bookingService.ShouldError{
		return nil,errors.New("sjjds")
	}
	return nil,nil
}
func(bookingService *MockBookingService)CancelTicket(ctx context.Context,ticketId string)(error){
	if bookingService.ShouldError{
		return errors.New("sjdjsd")
	}
	return nil
}