package movieservice

import (
	"context"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MovieServiceInterface interface {
	AddMovie(ctx context.Context,movieName string,movieType string,movieDuration int)(string,error)
	ViewAllMovies()([]model.Movie,error)
}
