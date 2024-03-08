package routes

import (
	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapCharacterRoutes(router *mux.Router) {
	router.HandleFunc("/characters", handlers.CreateCharacter).Methods("POST")
	router.HandleFunc("/characters/{id}", handlers.GetCharacter).Methods("GET")
	router.HandleFunc("/characters", handlers.GetAllCharacters).Methods("GET")
	router.HandleFunc("/characters/{id}", handlers.UpdateCharacter).Methods("PUT")
	router.HandleFunc("/characters/{id}", handlers.DeleteCharacter).Methods("DELETE")
}
