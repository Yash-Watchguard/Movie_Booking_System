package showservice

import (
	"context"
	"errors"
	"time"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/models/contextkey"
	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/models/roles"
	showrepo "github.com/Yash-Watchguard/MovieTicketBooking/internal/repository/show_repo"
	"github.com/Yash-Watchguard/MovieTicketBooking/utills"
)

type ShowService struct {
	showRepo showrepo.ShowRepoInterface
}

func NewShowService(showRepo showrepo.ShowRepoInterface) *ShowService {
	return &ShowService{showRepo: showRepo}
}

func (showService *ShowService) CreateShow(ctx context.Context, show *model.Show) (string, error) {
	userRole := ctx.Value(contextkey.UserRole).(role.Role)

	if userRole != role.Admin {
		return "", errors.New("unauthoe=rized for adding the show")
	}

	currentTime := time.Now()
	if !show.EndTime.After(show.StartTime) || !show.StartTime.After(currentTime) {
		return "", errors.New("enter valid date start time cannot be in the past or endtime should be greater the start time")
	}

	if ok, err := showService.showRepo.IsConflict(show.StartTime, show.EndTime); ok || err != nil {
		return "", errors.New("show timings conflicting with other show")
	}

	showId := utills.GenerateUuid()
	show.ShowId = showId
	
	if err := showService.showRepo.CreateShow(show); err != nil {
		return "", err
	}

	return showId, nil
}

func (showService *ShowService) GetAllShow() ([]model.Show, error) {
	shows, err := showService.showRepo.GetAllShow()
	if err != nil {
		return nil, err
	}
	if len(shows) == 0 {
		return nil, errors.New("no show available")
	}
	return shows, nil

}

func (showService *ShowService) GetShowByMovieId(movieId string) ([]model.Show, error) {
	shows, err := showService.showRepo.GetShowByMovieId(movieId)

	if err != nil {
		return nil, err
	}
	if len(shows) == 0 {
		return nil, errors.New("no show available for this movie")
	}

	return shows, nil
}

