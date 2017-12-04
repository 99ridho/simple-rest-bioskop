package models

import (
	"time"
)

type MovieSchedule struct {
	ID        int       `db:"id" json:"id"`
	MovieID   int       `db:"movie_id" json:"movie_id"`
	TheaterID int       `db:"theater_id" json:"theater_id"`
	PlayDate  time.Time `db:"play_date" json:"play_date"`
}

func NewMovieSchedule(id, movieId, theaterId int, playDate string) (*MovieSchedule, error) {
	parsedPlayDate, err := time.Parse("2006-01-02 15:04", playDate)
	if err != nil {
		return nil, err
	}

	s := &MovieSchedule{
		ID:        id,
		MovieID:   movieId,
		TheaterID: theaterId,
		PlayDate:  parsedPlayDate,
	}

	if err := Dbm.Insert(s); err != nil {
		return nil, err
	}

	return s, nil
}

func (m *MovieSchedule) Theater() (*MovieTheater, error) {
	theaterID := m.TheaterID
	var theater MovieTheater

	if err := Dbm.SelectOne(&theater, "select * from movie_theaters where id = ?", theaterID); err != nil {
		return nil, err
	}

	return &theater, nil
}

func (m *MovieSchedule) Movie() (*Movie, error) {
	movieID := m.MovieID
	var movie Movie

	if err := Dbm.SelectOne(&movie, "select * from movies where id = ?", movieID); err != nil {
		return nil, err
	}

	return &movie, nil
}
