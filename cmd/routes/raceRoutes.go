package routes

import (
	"github.com/gorilla/mux"
	middlewares "github.com/justkurama/GO-onepiece/cmd/middleware"
	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapRaceRoutes(router *mux.Router) {
	api := router.PathPrefix("/races").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	router.HandleFunc("", handlers.GetAllRaces).Methods("GET")
	router.HandleFunc("/{id}", handlers.GetRace).Methods("GET")
}
