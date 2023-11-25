package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maksimulitin/pkg/models"
	"github.com/maksimulitin/pkg/utils"
)

var NewMovie models.Movie

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	CreateMovie := &models.Movie{}
	utils.ParserBody(*r, CreateMovie)
	b := CreateMovie.CreateMovie()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func Getmovie(w http.ResponseWriter, r *http.Request) {
	movie := models.GetAllMovies()
	res, _ := json.Marshal(movie)
	w.Header().Set("Content-type", "pkg/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	id, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("Error ")
	}
	movieDetails, _ := models.GetMovieById(id)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "pkg/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdataMovie(w http.ResponseWriter, r *http.Request) {

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {

}
