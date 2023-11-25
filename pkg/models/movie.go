package models

import (
	"github.com/jinzhu/gorm"
	"github.com/maksimulitin/pkg/config"
)

type Movie struct {
	gorm.Model
	Genre   string `json:"genre"`
	Release string `json:"release"`
	Name    string `gorm:""json:"name"`
	Studio  string `json:"studio"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Movie{})
}

func (m *Movie) CreateMovie() *Movie {
	db.NewRecord(m)
	db.Create(&m)
	return m
}

func GetAllMovies() []Movie {
	var movies []Movie
	db.Find(&movies)
	return movies
}

func GetMovieById(id int64) (Movie, gorm.DB) {
	var getmovie Movie
	db := db.Where("ID=?", id).Find(&getmovie)
	return getmovie, *db
}

func DeleteMovie(id int64) Movie {
	var movie Movie
	db.Where("ID=?", id).Delete(movie)
	return movie
}
