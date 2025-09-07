package handlers

import (
	"encoding/json"
	"net/http"

	
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/response"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/movieservice"
)

type MovieHandler struct {
	movieService movieservice.MovieServiceInterface
}

func NewMovieHandler(movieService movieservice.MovieServiceInterface) *MovieHandler {
	return &MovieHandler{movieService: movieService}
}

func (movieHandler *MovieHandler) AddMovie(w http.ResponseWriter, r *http.Request) {
	ctx:=r.Context()
	type MovieData struct {
		Name      string `json:"name"`
		MovieType string `json:"movie_type"`
		Duration  int    `json:"duration"`
	}
    
	var NewMovie MovieData

	err:=json.NewDecoder(r.Body).Decode(&NewMovie)

	if err!=nil{
		response.ErrorResponse(w,"Invalid request body",http.StatusBadRequest,1000)
		return
	}

	if NewMovie.Duration<=0{
		response.ErrorResponse(w,"invalid movie duration",http.StatusBadRequest,1000)
	    return
	}

	movieId,err:=movieHandler.movieService.AddMovie(ctx,NewMovie.Name,NewMovie.MovieType,NewMovie.Duration)

	if err!=nil{
		response.ErrorResponse(w,err.Error(),http.StatusInternalServerError,1000)
		return
	}
    

	response.SuccessResponse(w,map[string]interface{}{"MovieId":movieId},"movie added successfully",http.StatusCreated)
}

func(movieHandler *MovieHandler)ViewAllMovies(w http.ResponseWriter,r *http.Request){
	
  movies,err:=movieHandler.movieService.ViewAllMovies()

  if err!=nil{
	response.ErrorResponse(w,err.Error(),http.StatusInternalServerError,1000)
	return
  }
  response.SuccessResponse(w,movies,"movies retrived successfully",http.StatusOK)

}

