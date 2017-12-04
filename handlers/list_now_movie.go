package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/99ridho/simple-rest-bioskop/errors"
	"github.com/99ridho/simple-rest-bioskop/models"
)

func ListNowMovieSchedule(w http.ResponseWriter, r *http.Request) {
	query := "SELECT * FROM movie_schedules WHERE DATE(play_date) = DATE(NOW())"
	var result []models.MovieSchedule

	if _, err := models.Dbm.Select(&result, query); err != nil {
		errors.NewHTTPError(http.StatusInternalServerError, "Fail to fetch").WriteTo(w)
		return
	}

	movieScheduleResponse := make([]models.MovieScheduleResponse, 0)
	for _, v := range result {
		movie, _ := v.Movie()
		theater, _ := v.Theater()
		schedule := models.MovieScheduleResponse{
			Movie:         movie,
			TheaterDetail: theater,
			PlayDate:      v.PlayDate,
		}

		movieScheduleResponse = append(movieScheduleResponse, schedule)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": movieScheduleResponse,
	})
}
