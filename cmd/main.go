package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8880", r))
}
