package router

import (
	"goji.io"
	"goji.io/pat"

	"github.com/99ridho/simple-rest-bioskop/handlers"
)

func NewRouter() *goji.Mux {
	root := goji.NewMux()
	root.HandleFunc(pat.Get("/theaters"), handlers.ListAllTheater)

	return root
}
