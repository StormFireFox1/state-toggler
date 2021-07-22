package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"strconv"
)

var state = false

func toggleState(w http.ResponseWriter, _ *http.Request) {
	state = !state
	_, _ = w.Write([]byte("Done"))
}

func getState(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte(strconv.FormatBool(state)))
}


func getRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/toggle", toggleState)
	r.Get("/state", getState)

	return r
}

func main() {
	r := getRouter()

	err := http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatalf("Failed to activate HTTP server: %s\n", err)
	}
}