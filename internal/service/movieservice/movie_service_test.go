package movieservice

import (
	"context"
	"errors"

	"testing"

	repomock "github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/repoaitorymock"
	
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/models/contextkey"
	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/models/roles"
)



func TestAddMovie(t *testing.T) {
      tests:=[]struct{
        ctx context.Context
		movieName string
		movieType string
		movieDuration int
		MovieAddError bool
		ExpectedError error
	  }{
		{
          ctx:context.WithValue(context.Background(),contextkey.UserRole,role.Customer),
		  movieName: "sbdb",
		  movieType: "sjbsd",
		  movieDuration: 123,
		  MovieAddError: false,
		  ExpectedError: errors.New("unauthorized for adding a movie"),
		},
		{
          ctx:context.WithValue(context.Background(),contextkey.UserRole,role.Admin),
		  movieName: "sbdb",
		  movieType: "sjbsd",
		  movieDuration: 0,
		  MovieAddError: false,
		  ExpectedError: errors.New("invalid movie duration"),
		},
		{
		  ctx:context.WithValue(context.Background(),contextkey.UserRole,role.Admin),
		  movieName: "sbdb",
		  movieType: "sjbsd",
		  movieDuration: 5,
		  MovieAddError: true,
		  ExpectedError: errors.New("failed to add movie"),
		},
	  }

	  for _,test:=range tests{
		mockMovieRepo:=repomock.NewMockMovieRepo()

		if test.MovieAddError{
			mockMovieRepo.ShouldError=true
		}
		movieService:=NewMovieService(mockMovieRepo)

		_,err:=movieService.AddMovie(test.ctx,test.movieName,test.movieType,test.movieDuration)

		if err==nil ||err.Error()!=test.ExpectedError.Error(){
			t.Errorf("skdjshdj")
		}
	  }
}

func TestShowAllMovie(t *testing.T){
	tests:=[]struct{
         wanterr error
		 ShouldError bool
	}{
		{
           wanterr: errors.New("internal server err"),
		   ShouldError: true,
		},
		{
			wanterr: errors.New("no movies available"),
			ShouldError: false,
		},
		
	}

	for _,test:=range tests{
		mokeMovieRepo:=repomock.NewMockMovieRepo()

		if test.ShouldError{
			mokeMovieRepo.ShouldError=true
		}
	    movieService:=NewMovieService(mokeMovieRepo)
        _,err:=movieService.ViewAllMovies()
		if err.Error()!=test.wanterr.Error(){
           t.Errorf("want error %v and we got this error %v",test.wanterr.Error(),err.Error())
		}
	
}
}
