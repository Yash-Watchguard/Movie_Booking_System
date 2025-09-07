package servicemock

import (
	"context"
	"errors"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MockMovieService struct {
	Shoulderr bool
}

func (movieService *MockMovieService) AddMovie(ctx context.Context, movieName string, movieType string, movieDuration int) (string, error) {
	if movieService.Shoulderr{
		return "",errors.New("sjdjs")
	}
	return "",nil
}
func (movieService *MockMovieService) ViewAllMovies() ([]model.Movie, error) {
	if movieService.Shoulderr{
		return nil,errors.New("sds")
	}
	return nil,nil
}