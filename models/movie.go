package models

type Movie struct {
	ID          int    `db:"id" json:"id"`
	TheaterID   int    `db:"theater_id" json:"theater_id"`
	Name        string `db:"name" json:"name"`
	Synopsis    string `db:"synopsis" json:"synopsis"`
	Duration    int    `db:"duration" json:"duration"`
	Genre       string `db:"genre" json:"genre"`
	ReleaseYear int    `db:"release_year" json:"release_year"`
}

func NewMovie(id, theaterId int, name, synopsis string, duration int, genre string, year int) (*Movie, error) {
	m := &Movie{
		ID:          id,
		TheaterID:   theaterId,
		Name:        name,
		Synopsis:    synopsis,
		Duration:    duration,
		Genre:       genre,
		ReleaseYear: year,
	}

	if err := Dbm.Insert(m); err != nil {
		return nil, err
	}

	return m, nil
}
