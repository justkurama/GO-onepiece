package routes

import (
	"github.com/gorilla/mux"
	middlewares "github.com/justkurama/GO-onepiece/cmd/middleware"
	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapOrganizationRoutes(router *mux.Router) {

	api := router.PathPrefix("/organizations").Subrouter()
	api.Use(middlewares.AuthMiddleware)

	api.HandleFunc("", handlers.CreateOrganization).Methods("POST")

	api.HandleFunc("/{id}", handlers.GetOrganization).Methods("GET")
	api.HandleFunc("/{id}/character", handlers.GetOrganizationCharacters).Methods("GET")
	api.HandleFunc("/{id}/suborganizations", handlers.GetOrganizationSubOrg).Methods("GET")
	api.HandleFunc("", handlers.GetAllOrganizations).Methods("GET")
	api.HandleFunc("/{id}", handlers.UpdateOrganization).Methods("PUT")
	api.HandleFunc("/{id}", handlers.DeleteOrganization).Methods("DELETE")
}
