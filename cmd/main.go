package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maksimulitin/pkg/routes"
)

func main() {
	r := mux.NewRouter()

	routes.RegisterMovieStoreRoutes(w, r)
	r.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8880", r))
}
