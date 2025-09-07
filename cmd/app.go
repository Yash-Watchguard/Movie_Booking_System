package main

import (
	"database/sql"
	"log"
	"net/http"

	"fmt"

	movierepo "github.com/Yash-Watchguard/MovieTicketBooking/internal/repository/movie_repo"
	showrepo "github.com/Yash-Watchguard/MovieTicketBooking/internal/repository/show_repo"
	ticketrepo "github.com/Yash-Watchguard/MovieTicketBooking/internal/repository/ticket_repo"
	userrepo "github.com/Yash-Watchguard/MovieTicketBooking/internal/repository/user_repo"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/routers"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/authservice"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/bookingservice"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/movieservice"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/showservice"
)

func RunApp(db *sql.DB){

	UserRepo:=userrepo.NewUserRepo(db)
	MovieRepo:=movierepo.NewMovieRepo(db)
	showRepo:=showrepo.NewShowRepo(db)
	ticketRepo:=ticketrepo.NewTicketRepo(db)

	AuthService:=authservice.NewAuthService(UserRepo)
	MovieService:=movieservice.NewMovieService(MovieRepo)
	showService:=showservice.NewShowService(showRepo)
	bookingservice:=bookingservice.NewBookingService(ticketRepo,showRepo)


	router:=routers.SetUpRouter(AuthService,MovieService,showService,bookingservice)
    fmt.Println("Server stating on the Port 8080...")
	err:=http.ListenAndServe(":8080",router)

	if err!=nil{
		log.Fatal(err)
	}
     
	
}
