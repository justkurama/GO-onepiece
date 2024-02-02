package onepiece

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting API server")
	router := mux.NewRouter()

	log.Println("Creating routes")
	router.HandleFunc("/mugiwaras", handlers.GetPlayers).Methods("GET")
	router.HandleFunc("/mugiwaras/last_name/{nickname}", handlers.GetPlayerByName).Methods("GET")
	router.HandleFunc("/mugiwaras/{id}", handlers.GetPlayerById).Methods("GET")
	router.HandleFunc("/health-checking", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/mugiwaras/position/{position}", handlers.GetPlayersByPosition).Methods("GET")
	router.HandleFunc("/mugiwaras/nation/{nation}", handlers.GetPlayersByNation).Methods("GET")
	http.Handle("/", router)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
