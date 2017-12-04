package main

import (
	"net/http"

	"github.com/99ridho/simple-rest-bioskop/router"
)

func main() {
	r := router.NewRouter()
	http.ListenAndServe(":1313", r)
}
