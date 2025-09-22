package model

type Movie struct {
	MovieId   string    `json:"movie_id"`
	MovieName string    `json:"movie_name"`
	MovieType string    `json:"movie_type"`
	Duration  int `json:"duration"`
}

