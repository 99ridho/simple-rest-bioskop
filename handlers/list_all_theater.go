package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/99ridho/simple-rest-bioskop/errors"
	"github.com/99ridho/simple-rest-bioskop/models"
)

func ListAllTheater(w http.ResponseWriter, r *http.Request) {

	query := "SELECT * FROM movie_theaters"
	var result []models.MovieTheater

	if _, err := models.Dbm.Select(&result, query); err != nil {
		errors.NewHTTPError(http.StatusInternalServerError, "Fail to fetch").WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": result,
	})
}
