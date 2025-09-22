package bookingservice

import (
	"context"
	"errors"
	"time"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/contextkey"
	showrepo "github.com/Yash-Watchguard/MovieTicketBooking/internal/repository/show_repo"
	ticketrepo "github.com/Yash-Watchguard/MovieTicketBooking/internal/repository/ticket_repo"
	"github.com/Yash-Watchguard/MovieTicketBooking/utills"
)

type BookingService struct {
	ticketRepo ticketrepo.TicketRepoInterface
	showRepo   showrepo.ShowRepoInterface
}

func NewBookingService(ticketrepo ticketrepo.TicketRepoInterface, showRepo showrepo.ShowRepoInterface) *BookingService {
	return &BookingService{ticketRepo: ticketrepo, showRepo: showRepo}
}

func (bookingService *BookingService) BookTicket(showId string, userId string, numberOfTickets int) ([]model.Ticket, error) {

	showDetails, err := bookingService.showRepo.GetShowByShowId(showId)
	if err != nil {
		return nil, errors.New("no show available")
	}
	if showDetails.AvailableSeat < numberOfTickets {
		return nil, errors.New("required Number of seats are not available")
	}

	var tickets []model.Ticket
	for i := 0; i < numberOfTickets; i++ {

		ticket := model.Ticket{
			TicketId:    utills.GenerateUuid(),
			ShowId:      showId,
			UserId:      userId,
			BookingTime: time.Now(),
		}
		showDetails.AvailableSeat = showDetails.AvailableSeat - 1
		tickets = append(tickets, ticket)
	}

	err = bookingService.showRepo.UpdateShow(showDetails.AvailableSeat, showId)
	if err != nil {
		return nil, errors.New("error in updating the show")
	}

	err = bookingService.ticketRepo.SaveTickets(tickets)
	if err != nil {
		return nil, errors.New("error in saving Tickets")
	}
	return tickets, nil
}

func (bookingService *BookingService) CancelTicket(ctx context.Context, ticketId string) error {
	// get the ticket by ticketId
	ticket, err := bookingService.ticketRepo.GetTicketById(ticketId)

	if err != nil {
		return errors.New("ticket is not available")

	}

	userId := ctx.Value(contextkey.UserId).(string)

	if userId != ticket.UserId {
		return errors.New("you are not allowed to cancel this ticket")
	}

	err = bookingService.ticketRepo.CancleTicket(ticketId)
	if err != nil {
		return errors.New("enable to cancel ticket")
	}

	show, _ := bookingService.showRepo.GetShowByShowId(ticket.ShowId)
	showId := show.ShowId
	err = bookingService.showRepo.UpdateShow(show.AvailableSeat+1, showId)
	if err != nil {
		return err
	}
	return nil
}
