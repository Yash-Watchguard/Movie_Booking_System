package showservice

import (
	"context"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type ShowServiceInterface interface {
	CreateShow(ctx context.Context,newShow *model.Show)(string,error)
	GetAllShow()([]model.Show,error)
	GetShowsByMovieId(movieId string)([]model.Show,error)
}