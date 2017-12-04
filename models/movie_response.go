package models

import (
	"time"
)

type MovieResponse struct {
	Movie
	TheaterDetail *MovieTheater `json:"theater_detail"`
}

type MovieScheduleResponse struct {
	*Movie
	TheaterDetail *MovieTheater `json:"theater_detail"`
	PlayDate      time.Time     `json:"play_date"`
}
