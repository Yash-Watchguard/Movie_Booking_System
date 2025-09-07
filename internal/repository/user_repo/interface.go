package userrepo

import (
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	
)

type UserRepoInterface interface {
	SaveUser(userId,name,email,phoneNumber,password string)error
	GetUserById(userId string) (*model.User,error)
	GetUserByEmail(email string)(*model.User,error)
}