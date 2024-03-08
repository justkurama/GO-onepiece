package routes

import (
	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapOrganizationRoutes(router *mux.Router) {
	router.HandleFunc("/organizations", handlers.CreateOrganization).Methods("POST")
	router.HandleFunc("/organizations/{id}", handlers.GetOrganization).Methods("GET")
	router.HandleFunc("/organizations", handlers.GetAllOrganizations).Methods("GET")
	router.HandleFunc("/organizations/{id}", handlers.UpdateOrganization).Methods("PUT")
	router.HandleFunc("/organizations/{id}", handlers.DeleteOrganization).Methods("DELETE")
}
