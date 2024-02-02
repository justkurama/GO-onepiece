package onepiece

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/justkurama/GOtsis1/internal"
)

func main() {
	log.Println("Starting API server")
	router := mux.NewRouter()

	log.Println("Creating routes")
	router.HandleFunc("/mugiwaras", handlers.GetMugiwara).Methods("GET")
	router.HandleFunc("/mugiwaras/name/{name}", handlers.GetMugiwaraByName).Methods("GET")
	router.HandleFunc("/mugiwaras/{id}", handlers.GetMugiwaraById).Methods("GET")
	router.HandleFunc("/health-checking", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/mugiwaras/bounty/{bounty}", handlers.GetMugiwarasByBounty).Methods("GET")
	router.HandleFunc("/mugiwaras/origin/{origin}", handlers.GetMugiwarasByOrigin).Methods("GET")
	http.Handle("/", router)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
