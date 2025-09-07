package bookingservice

import (
	
	"errors"
	"testing"

	repomock "github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/repoaitorymock"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	
)

// TestBookTicket contains the complete table-driven tests

func TestXxx(t *testing.T) {

	const validShowId = "hdsdshd"
	const validuserId ="skdjshd"
	show:=&model.Show{
		ShowId: validShowId,
		AvailableSeat: 15,
	}
	test:=[]struct{
		name string
        showId string
		userId string
		numberOfTickets int
		MockTicketRepo *repomock.MockTicketRepo
		MockShowRepo   *repomock.MockShowRepo

		WantError  error
	}{
		{
			name:"show not present",
         showId: "hsdhs",
		 userId: "sjdjs",
		 numberOfTickets: 2,
		 MockTicketRepo: &repomock.MockTicketRepo{
			ShouldError: true,
		 },
		 MockShowRepo: &repomock.MockShowRepo{
			ShowIdError: true,
		 },
		 WantError: errors.New("no show available"),
		},
		{   name:"invalid seats",
			showId: validShowId,
			userId: validuserId,
			numberOfTickets: 50,
			MockTicketRepo: &repomock.MockTicketRepo{},
			MockShowRepo: &repomock.MockShowRepo{
				Shows: map[string]*model.Show{validShowId:show},
			},
			WantError: errors.New("required Number of seats are not available"),
		},
		{
			name: "update show",
			showId: validShowId,
			userId: validuserId,
			numberOfTickets: 2,
			MockTicketRepo: &repomock.MockTicketRepo{},
			MockShowRepo: &repomock.MockShowRepo{
				Shows: map[string]*model.Show{validShowId:show},
				UpdateShowerro: true,
			},
			WantError: errors.New("error in updating the show"),
		},
		{   name: "save ticket",
            showId: validShowId,
			userId: validuserId,
			numberOfTickets: 2,
			MockTicketRepo: &repomock.MockTicketRepo{
				ShouldError: true,
			},
			MockShowRepo: &repomock.MockShowRepo{
				Shows: map[string]*model.Show{validShowId:show},
			},
			WantError: errors.New("error in saving Tickets"), 
		},

	}
	for _,test:=range test{
		t.Run(test.name,func(t *testing.T) {bookingService:=NewBookingService(test.MockTicketRepo,test.MockShowRepo)

		_,err:=bookingService.BookTicket(test.showId,test.userId,test.numberOfTickets)

		if err==nil || test.WantError.Error()!=err.Error(){
			t.Errorf("jdshdj")
		}})
	}
}

