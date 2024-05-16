package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/cmd/routes"

	"github.com/justkurama/GO-onepiece/internal/pkg/database"
)

func runServer() {
	database.Migrate()
	log.Println("Starting API server")
	router := mux.NewRouter()
	log.Println("Creating routes")
	routes.MapCharacterRoutes(router)
	routes.MapOrganizationRoutes(router)
	routes.MapOriginRoutes(router)
	routes.MapRaceRoutes(router)
	routes.MapAuthRoutes(router)
	routes.MapSubOrganizationRoutes(router)

	http.Handle("/", router)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
