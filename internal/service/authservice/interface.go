package authservice

import model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"

type AuthServiceInterface interface {
	Login(userId string, email string, password string) (*model.User, string, error)
	SignUp(name string, userEmail string, phoneNumber string, Password string) (string, error)
}