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
	return nil,nil
}
func(showRepo *MockShowRepo)GetShowByMovieId(movieId string)([]model.Show,error){
	if showRepo.ShouldError{
		return nil,errors.New("sdjs")
	}
	return nil,nil
}

func(showRepo *MockShowRepo)UpdateShow(updatedSeat int,showId string)error{
	if showRepo.UpdateShowerro{
		return errors.New("dhfhd")
	}
	return nil
}
func(showRepo *MockShowRepo)IsConflict(showStartTime,showEndTime time.Time)(bool,error){
	if showRepo.ShouldError{
		return true,nil
	}
	return false,errors.New("sjd")
}
func(showRepo *MockShowRepo)GetShowByShowId(showId string)(*model.Show,error){
	if showRepo.ShowIdError{
		return nil,errors.New("sdhs")
	}
	return showRepo.Shows[showId],nil
}