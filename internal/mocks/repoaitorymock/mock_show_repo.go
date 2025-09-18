package repomock

import (
	"errors"
	"time"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MockShowRepo struct {
	Shows map[string]*model.Show
	ShouldError bool
	ShowIdError bool
	UpdateShowerro bool
	HasShows bool
	Conflict bool
	IsConflictError bool
}

func NewMokeShowRepo()*MockShowRepo{
	return &MockShowRepo{Shows: make(map[string]*model.Show)}
}

func(mokeshowRepo *MockShowRepo)CreateShow(show *model.Show)(error){
	if mokeshowRepo.ShouldError{
		return errors.New("jsdsjd")
	}
	return nil
}
func(showRepo *MockShowRepo)GetAllShow()([]model.Show,error){
	if showRepo.ShouldError{
		return nil,errors.New("sjds")
	}
	if showRepo.HasShows {
		return []model.Show{{ShowId: "show1", MovieId: "movie1"}}, nil
	}
	return []model.Show{}, nil
}
func(showRepo *MockShowRepo)GetShowByMovieId(movieId string)([]model.Show,error){
	if showRepo.ShouldError{
		return nil,errors.New("sdjs")
	}
	if showRepo.HasShows {
		return []model.Show{{ShowId: "show1", MovieId: movieId}}, nil
	}
	return []model.Show{}, nil
}

func(showRepo *MockShowRepo)UpdateShow(updatedSeat int,showId string)error{
	if showRepo.UpdateShowerro{
		return errors.New("dhfhd")
	}
	return nil
}
func(showRepo *MockShowRepo)IsConflict(showStartTime,showEndTime time.Time)(bool,error){
	if showRepo.IsConflictError {
		return false, errors.New("repo error on conflict check")
	}
	if showRepo.Conflict {
		return true, nil
	}
	return false, nil
}
func(showRepo *MockShowRepo)GetShowByShowId(showId string)(*model.Show,error){
	if showRepo.ShowIdError{
		return nil,errors.New("sdhs")
	}
	return showRepo.Shows[showId],nil
}