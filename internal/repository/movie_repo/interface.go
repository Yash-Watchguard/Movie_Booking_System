package movierepo

import (
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MovieRepoInterface interface {
	AddMovie(newMovie model.Movie)(error)
	ViewAllMovies()([]model.Movie,error)
}
