package router

import (
	"goji.io"
	"goji.io/pat"

	"github.com/99ridho/simple-rest-bioskop/handlers"
)

func NewRouter() *goji.Mux {
	root := goji.NewMux()
	root.HandleFunc(pat.Get("/theaters"), handlers.ListAllTheater)
	root.HandleFunc(pat.Get("/movies"), handlers.ListAllMovie)
	root.HandleFunc(pat.Get("/theater/:theater_id/movies"), handlers.ListAllMovieByTheaterId)
	root.HandleFunc(pat.Get("/movie/upcoming"), handlers.ListUpcomingMovieSchedule)
	root.HandleFunc(pat.Get("/movie/now"), handlers.ListNowMovieSchedule)

	return root
}
