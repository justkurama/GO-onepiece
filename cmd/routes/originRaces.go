package routes

import (
	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapOriginRoutes(router *mux.Router) {
	router.HandleFunc("/origins/{id}", handlers.GetOrigin).Methods("GET")
	router.HandleFunc("/origins", handlers.GetAllOrigins).Methods("GET")
}
