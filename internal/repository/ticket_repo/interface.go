package ticketrepo

import model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"

type TicketRepoInterface interface {
	SaveTickets(tickets []model.Ticket)error
	GetTicketById(ticketId string)(*model.Ticket,error)
	CancleTicket(ticktId string)(error)
}