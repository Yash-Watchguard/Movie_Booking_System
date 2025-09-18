package repomock

import (
	"errors"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MockTicketRepo struct {
	Tickets       map[string]*model.Ticket
	ShouldError   bool // for SaveTickets
	CancelError   bool // for CancleTicket
	GetTicketError bool // for GetTicketById
}

func NewMockTicketRepo() *MockTicketRepo {
	return &MockTicketRepo{Tickets: make(map[string]*model.Ticket)}
}

// SaveTickets simulates saving tickets
func (ticketRepo *MockTicketRepo) SaveTickets(tickets []model.Ticket) error {
	if ticketRepo.ShouldError {
		return errors.New("error in saving Tickets")
	}
	for _, t := range tickets {
		copy := t // avoid referencing loop variable
		ticketRepo.Tickets[t.TicketId] = &copy
	}
	return nil
}

// GetTicketById simulates fetching ticket
func (ticketRepo *MockTicketRepo) GetTicketById(ticketId string) (*model.Ticket, error) {
	if ticketRepo.GetTicketError {
		return nil, errors.New("ticket is not available")
	}
	t, ok := ticketRepo.Tickets[ticketId]
	if !ok {
		return nil, errors.New("ticket is not available")
	}
	return t, nil
}

// CancleTicket simulates canceling a ticket
func (ticketRepo *MockTicketRepo) CancleTicket(ticketId string) error {
	if ticketRepo.CancelError {
		return errors.New("enable to cancel ticket")
	}
	if _, ok := ticketRepo.Tickets[ticketId]; !ok {
		return errors.New("ticket is not available")
	}
	delete(ticketRepo.Tickets, ticketId)
	return nil
}
