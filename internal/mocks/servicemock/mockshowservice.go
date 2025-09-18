package servicemock

import (
	"context"
	"errors"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MockShowService struct {
	Shoulderr bool
}

func (showService *MockShowService) CreateShow(ctx context.Context, newShow *model.Show) (string, error) {
	if showService.Shoulderr {
		return "", errors.New("create show error")
	}
	return "show123", nil
}

func (showService *MockShowService) GetAllShow() ([]model.Show, error) {
	if showService.Shoulderr {
		return nil, errors.New("get all show error")
	}
	return []model.Show{{MovieId: "movie1"}}, nil
}

func (showService *MockShowService) GetShowByMovieId(movieId string) ([]model.Show, error) {
	if showService.Shoulderr {
		return nil, errors.New("get show by movie id error")
	}
	return []model.Show{{MovieId: movieId}}, nil
}
