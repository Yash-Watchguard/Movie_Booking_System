package servicemock

import (
	"errors"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

type MockAuthService struct {
	ShouldError bool
}

func (authService *MockAuthService) SignUp(name string, userEmail string, phoneNumber string, Password string) (string, error) {
	if authService.ShouldError {
		return " ", errors.New("kdfjdfjkdkjf")
	}
	return "",nil
}
func(authService *MockAuthService)Login(name string,email string,password string)(*model.User,string,error){
	if authService.ShouldError{
		return nil,"",errors.New("sjdjs")
	}
	return nil,"",nil
}