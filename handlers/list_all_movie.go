package handlers

import (
	"encoding/json"
	"net/http"

	"goji.io/pat"

	"github.com/99ridho/simple-rest-bioskop/errors"
	"github.com/99ridho/simple-rest-bioskop/models"
)

func ListAllMovie(w http.ResponseWriter, r *http.Request) {

	query := "SELECT * FROM movies"
	var result []models.Movie

	if _, err := models.Dbm.Select(&result, query); err != nil {
		errors.NewHTTPError(http.StatusInternalServerError, "Fail to fetch").WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": result,
	})
}

func ListAllMovieByTheaterId(w http.ResponseWriter, r *http.Request) {
	theaterID := pat.Param(r, "theater_id")
	query := "SELECT * FROM movies WHERE theater_id = ?"
	var result []models.Movie

	if _, err := models.Dbm.Select(&result, query, theaterID); err != nil {
		errors.NewHTTPError(http.StatusInternalServerError, "Fail to fetch").WriteTo(w)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": result,
	})
}
