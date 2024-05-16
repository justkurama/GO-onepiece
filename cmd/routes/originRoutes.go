package routes

import (
	"github.com/gorilla/mux"
	middlewares "github.com/justkurama/GO-onepiece/cmd/middleware"
	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapOriginRoutes(router *mux.Router) {

	api := router.PathPrefix("/origins").Subrouter()
	api.Use(middlewares.AuthMiddleware)
	api.HandleFunc("/{id}", handlers.GetOrigin).Methods("GET")
	api.HandleFunc("/{id}/character", handlers.GetOriginCharacters).Methods("GET")
	api.HandleFunc("", handlers.GetAllOrigins).Methods("GET")

}
