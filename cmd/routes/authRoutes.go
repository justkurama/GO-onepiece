package routes

import (
	"github.com/gorilla/mux"

	"github.com/justkurama/GO-onepiece/internal/app/handlers"
)

func MapAuthRoutes(router *mux.Router) {
	router.HandleFunc("/signup", handlers.SignUp).Methods("POST")
	router.HandleFunc("/refresh", handlers.RefreshToken).Methods("POST")
	router.HandleFunc("/signin", handlers.SignIn).Methods("POST")
}
