package routers

import (
	"net/http"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/handlers"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/middleware"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/authservice"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/bookingservice"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/movieservice"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/showservice"
)

func SetUpRouter(authService authservice.AuthServiceInterface,movieService movieservice.MovieServiceInterface, showService showservice.ShowServiceInterface,bookingService bookingservice.BookingServiceInterface)*http.ServeMux{
   r:=http.NewServeMux()

   authHandler:=handlers.NewAuthHandler(authService)
   movieHandler:=handlers.NewMovieHandler(movieService)
   showHandler:=handlers.NewShowHandler(showService)
   bookingHandler:=handlers.NewBookingHandler(bookingService)

   r.Handle("/v1/signup",(http.HandlerFunc(authHandler.SignUp)))
   r.Handle("/v1/login",http.HandlerFunc(authHandler.Login))

   r.Handle("/v1/movies/addmovie",middleware.AuthMiddleware(http.HandlerFunc(movieHandler.AddMovie)))
   r.Handle("/v1/movies/viewmovie/",http.HandlerFunc(movieHandler.ViewAllMovies))

   r.Handle("/v1/shows/addshow",middleware.AuthMiddleware(http.HandlerFunc(showHandler.CreateShow)))
   r.Handle("/v1/shows/viewshows/{movie_id}",http.HandlerFunc(showHandler.GetAllShow))

   r.Handle("/v1/booking/bookticket/{show_id}",middleware.AuthMiddleware(http.HandlerFunc(bookingHandler.BookTicket)))
   r.Handle("/v1/booking/cancelticket/{ticket_id}",middleware.AuthMiddleware(http.HandlerFunc(bookingHandler.CancelTicket)))

   return r
}