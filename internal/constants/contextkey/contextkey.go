package contextkey

import role "github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/roles"

type ContextKey role.Role

const (
	UserId   ContextKey = "user_id"
	UserRole ContextKey = "user_role"
)