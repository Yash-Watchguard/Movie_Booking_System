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
		return errors.New("mock error")
	  }
	  mr.movies[newMovie.MovieId] = newMovie
      return nil
}
func(mr *MockMovieRepo)ViewAllMovies()([]model.Movie,error){
    if mr.ShouldError{
		return nil,errors.New("mock error")
	}

	movies:=[]model.Movie{}
	for _, movie := range mr.movies {
		movies = append(movies, movie)
	}
	return movies,nil
}
