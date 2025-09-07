package model

import "time"

type Show struct {
	ShowId     string `json:"show_id"`
	MovieId    string `json:"movie_id"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	TotalSeats int `json:"total_seats"`
	AvailableSeat int `json:"available_seat"`
}