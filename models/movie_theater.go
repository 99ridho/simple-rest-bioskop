package models

type MovieTheater struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Address     string `db:"address" json:"address"`
	Description string `db:"description" json:"description"`
}

func NewMovieTheater(id int, name, address, description string) (*MovieTheater, error) {
	t := &MovieTheater{
		ID:          id,
		Name:        name,
		Address:     address,
		Description: description,
	}

	if err := Dbm.Insert(t); err != nil {
		return nil, err
	}

	return t, nil
}
