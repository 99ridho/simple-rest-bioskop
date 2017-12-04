package models

type MovieResponse struct {
	Movie
	TheaterDetail *MovieTheater `json:"theater_detail"`
}
