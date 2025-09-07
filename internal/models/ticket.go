package model

import "time"

type Ticket struct {
	TicketId     string `json:"ticket_id"`
	ShowId       string `json:"show_id"`
	UserId string `json:"user_id"`
	BookingTime time.Time `json:"booking_time"`
}