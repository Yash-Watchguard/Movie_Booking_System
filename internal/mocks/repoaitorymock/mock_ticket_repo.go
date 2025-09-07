package repomock

import (
	"errors"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MockTicketRepo struct {
	Tickets map[string]*model.Ticket
	ShouldError bool
	CancleError bool
}

func NewMockTicketRepo()*MockTicketRepo{
	return &MockTicketRepo{Tickets: make(map[string]*model.Ticket)}
}

func(ticketRepo *MockTicketRepo)SaveTickets(tickets []model.Ticket)error{
	if ticketRepo.ShouldError{
		return errors.New("sjdj")
	}
	return nil
}
func(ticketRepo *MockTicketRepo)GetTicketById(ticketId string)(*model.Ticket,error){
	if ticketRepo.ShouldError{
		return nil,errors.New("jdj")
	}
	return nil,nil
}

func(ticketRepo *MockTicketRepo)CancleTicket(ticketId string)(error){
	if ticketRepo.CancleError{
		return errors.New("sjdjs")
	}
	return nil
}