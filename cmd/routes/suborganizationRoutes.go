package routes

import (
	"github.com/gorilla/mux"
	middlewares "github.com/justkurama/GO-onepiece/cmd/middleware"
	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapSubOrganizationRoutes(router *mux.Router) {

	api := router.PathPrefix("/suborganizations").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	api.HandleFunc("", handlers.CreateSubOrganization).Methods("POST")
	api.HandleFunc("/{id}", handlers.GetSubOrganization).Methods("GET")
	api.HandleFunc("/{id}/character", handlers.GetSubOrganizationCharacters).Methods("GET")
	api.HandleFunc("", handlers.GetAllSubOrganizations).Methods("GET")
	api.HandleFunc("/{id}", handlers.UpdateSubOrganization).Methods("PUT")
	api.HandleFunc("/{id}", handlers.DeleteSubOrganization).Methods("DELETE")
}
