package repomock

import (
	"errors"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MockMovieRepo struct {
	movies map[string]model.Movie
	ShouldError bool
}

func NewMockMovieRepo() *MockMovieRepo {
	return &MockMovieRepo{movies: make(map[string]model.Movie)}
}

func (mr *MockMovieRepo)AddMovie(newMovie model.Movie)error{
	  if mr.ShouldError{
		return errors.New("sdjsd")
	  }
      return nil
}
func(mr *MockMovieRepo)ViewAllMovies()([]model.Movie,error){
    if mr.ShouldError{
		return nil,errors.New("jsjd")
	}

	movies:=[]model.Movie{}
	return movies,nil
}