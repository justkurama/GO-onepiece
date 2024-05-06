package routes

import (
	"github.com/gorilla/mux"
	middlewares "github.com/justkurama/GO-onepiece/cmd/middleware"
	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapCharacterRoutes(router *mux.Router) {

	characterRouter := router.PathPrefix("/characters").Subrouter()
	characterRouter.Use(middlewares.AuthMiddleware)

	characterRouter.HandleFunc("", handlers.CreateCharacter).Methods("POST")
	characterRouter.HandleFunc("/{id}", handlers.GetCharacter).Methods("GET")
	characterRouter.HandleFunc("", handlers.GetAllCharacters).Methods("GET")
	characterRouter.HandleFunc("/{id}", handlers.UpdateCharacter).Methods("PUT")
	characterRouter.HandleFunc("/{id}", handlers.DeleteCharacter).Methods("DELETE")
}
