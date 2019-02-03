package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github/sorix6/tic-tac-toe/controllers"
)

func main() {
	router := mux.NewRouter()

	// add the end-points needed for the game
	router.HandleFunc("/{gameId}/reset", controllers.Reset).Methods("DELETE")
	router.HandleFunc("/{gameId}/add/{player}", controllers.AddPlay).Methods("POST")
	router.HandleFunc("/{gameId}/get-status", controllers.GetStatus).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}


