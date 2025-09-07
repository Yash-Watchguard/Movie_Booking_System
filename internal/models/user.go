package model

import role "github.com/Yash-Watchguard/MovieTicketBooking/internal/models/roles"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password string `json:"password"`
	Role     role.Role `json:"role"`
}