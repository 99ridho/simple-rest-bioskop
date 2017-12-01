package main

import (
	"github.com/99ridho/simple-rest-bioskop/models"
)

func main() {
	if err := models.CreateTables(); err != nil {
		panic(err)
	}

	models.NewMovieTheater(1, "Bioskop Jaya", "Jalan Malang", "Bioskop Murah")
	models.NewMovieTheater(2, "Cinemaxxx", "Malang Town Square", "Bioskop Jaman Now")
	models.NewMovieTheater(3, "XXI", "Malang Town Square", "Bioskop Jaman Now")
	models.NewMovieTheater(4, "Dieng Bioskop", "Jalan Dieng", "Bioskop Jaman Now")

	models.NewMovie(1, 1, "Keluarga Tak Kasat Mata", "Film Horor", 145, "Horror", 2017)
	models.NewMovie(2, 2, "Keluarga Tak Kasat Mata", "Film Horor", 145, "Horror", 2017)
	models.NewMovie(3, 3, "Keluarga Tak Kasat Mata", "Film Horor", 145, "Horror", 2017)
	models.NewMovie(4, 4, "Keluarga Tak Kasat Mata", "Film Horor", 145, "Horror", 2017)
}
