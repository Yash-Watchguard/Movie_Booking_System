package handlers

import (
	"encoding/json"
	"net/http"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/response"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/showservice"
	"github.com/Yash-Watchguard/MovieTicketBooking/utills"
)

type ShowHandler struct {
	showService showservice.ShowServiceInterface
}

func NewShowHandler(showService showservice.ShowServiceInterface) *ShowHandler {
	return &ShowHandler{showService: showService}
}

func (showHandler *ShowHandler) CreateShow(w http.ResponseWriter, r *http.Request) {
	// first store tye data of request body
	ctx := r.Context()
	type showData struct {
		MovieId       string `json:"movie_id"`
		StartTime     string `json:"start_time"`
		EndTime       string `json:"end_time"`
		TotalSeat     int    `json:"total_seats"`
		AvailableSeat int    `json:"available_seats"`
	}
	var newShow showData

	err := json.NewDecoder(r.Body).Decode(&newShow)
	if err != nil {
		response.ErrorResponse(w, "Invalid request Body", http.StatusBadRequest)
		return
	}

	// check the time is valid or not
	parsedStartTime, isValid := utills.ValidateTime(newShow.StartTime)
	if !isValid {
		response.ErrorResponse(w, "Invalid Time Format", http.StatusBadRequest)
		return
	}
	parsedEndTime, isValid := utills.ValidateTime(newShow.EndTime)
	if !isValid {
		response.ErrorResponse(w, "Invalid Time format", http.StatusBadRequest)
		return
	}

	show := model.Show{
		MovieId:       newShow.MovieId,
		StartTime:     parsedStartTime,
		EndTime:       parsedEndTime,
		TotalSeats:    newShow.TotalSeat,
		AvailableSeat: newShow.AvailableSeat,
	}

	showId, err := showHandler.showService.CreateShow(ctx, &show)

	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.SuccessResponse(w, map[string]any{"ShowId":showId}, "Show Added Successfully", http.StatusCreated)

}
func (showHandler *ShowHandler) GetAllShow(w http.ResponseWriter, r *http.Request) {
	var shows []model.Show
	var err error
		shows, err = showHandler.showService.GetAllShow()
		if err != nil {
			response.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}
		response.SuccessResponse(w, shows, "Shows Retrived successfully", http.StatusOK)
}

func (showHandler *ShowHandler) GetAllShowofMovie(w http.ResponseWriter, r *http.Request) {
	var shows []model.Show
	var err error
	movieId := r.PathValue("movie_id")
	if movieId == "" {
		response.ErrorResponse(w, "Invalid Time Format", http.StatusBadRequest)
		return
	}

	shows, err = showHandler.showService.GetShowsByMovieId(movieId)
	if err != nil {
		response.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response.SuccessResponse(w, shows, "Shows Retrived successfully", http.StatusOK)
}
