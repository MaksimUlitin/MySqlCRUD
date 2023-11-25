package routes

import "github.com/gorilla/mux"

func RegisterMovieStoreRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/movie/")
	router.HandleFunc("/movie/")
	router.HandleFunc("/movie/{movieId}")
	router.HandleFunc("/movie/{movieId}")
	router.HandleFunc("/movie/{movieId}")
}
