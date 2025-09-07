package showrepo

import (
	"time"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type ShowRepoInterface interface {
	CreateShow(show *model.Show)(error)
	GetAllShow()([]model.Show,error)
	GetShowByMovieId(movieId string)([]model.Show,error)
	UpdateShow(updateSeat int,showId string)error
	IsConflict(startTime ,endTime time.Time)(bool,error)
	GetShowByShowId(showId string)(*model.Show,error)

}