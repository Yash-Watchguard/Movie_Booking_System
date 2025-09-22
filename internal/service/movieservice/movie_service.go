package movieservice

import (
	"context"
	"errors"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/contextkey"
	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/roles"
	movierepo "github.com/Yash-Watchguard/MovieTicketBooking/internal/repository/movie_repo"
	"github.com/Yash-Watchguard/MovieTicketBooking/utills"
)

type MovieService struct {
	movieRepo movierepo.MovieRepoInterface
}

func NewMovieService(movieRepo movierepo.MovieRepoInterface) *MovieService {
	return &MovieService{movieRepo: movieRepo}
}

func (movieService *MovieService) AddMovie(ctx context.Context, movieName string, movieType string, movieDuration int) (string, error) {
	userRole := ctx.Value(contextkey.UserRole).(role.Role)

	if userRole != role.Admin {
		return "", errors.New("unauthorized for adding a movie")
	}

	if movieDuration <= 0 {
		return "", errors.New("invalid movie duration")
	}

	newMovie := model.Movie{
		MovieId:   utills.GenerateUuid(),
		MovieName: movieName,
		MovieType: movieType,
		Duration:  movieDuration,
	}

	err := movieService.movieRepo.AddMovie(newMovie)
	if err != nil {
		return "", errors.New("failed to add movie")
	}

	return newMovie.MovieId, nil
}

func (movieService *MovieService) ViewAllMovies() ([]model.Movie, error) {
	movies, err := movieService.movieRepo.ViewAllMovies()

	if err != nil {
		return nil, errors.New("internal server err")
	}

	if len(movies) == 0 {
		return nil, errors.New("no movies available")
	}

	return movies, nil
}
