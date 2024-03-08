package routes

import (
	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapRaceRoutes(router *mux.Router) {
	router.HandleFunc("/races", handlers.GetAllRaces).Methods("GET")
	router.HandleFunc("/races/{id}", handlers.GetRace).Methods("GET")
}
