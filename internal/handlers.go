package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/justkurama/GOtsis1/api"
)

func GetMugiwara(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(api.Mugiwaras)
}

func GetMugiwaraByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	mugiwaraName := params["name"]

	mugiwaraName = strings.ToLower(mugiwaraName)

	for _, mugiwara := range Mugiwaras {
		if strings.ToLower(mugiwara.Name) == mugiwaraName {
			json.NewEncoder(w).Encode(mugiwara)
			return
		}
	}

	http.Error(w, "Mugiwara not found", http.StatusNotFound)
}

func GetMugiwaraById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	mugiwaraId := params["id"]

	for _, mugiwara := range Mugiwaras {
		if strconv.Itoa(mugiwara.Id) == mugiwaraId {
			json.NewEncoder(w).Encode(mugiwara)
			return
		}
	}

	http.Error(w, "Mugiwara not found", http.StatusNotFound)
}

func GetMugiwarasByOrigin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	origin := params["origin"]

	origin = strings.ToLower(origin)
	var matchingMugiwaras []Mugiwara

	for _, mugiwara := range Mugiwaras {
		if strings.ToLower(mugiwara.Origin) == origin {
			matchingMugiwaras = append(matchingMugiwaras, mugiwara)
		}
	}

	if len(matchingMugiwaras) > 0 {
		json.NewEncoder(w).Encode(matchingMugiwaras)
	} else {
		http.Error(w, "Mugiwaras from the specified origin not found", http.StatusNotFound)
	}
}

func GetMugiwarasByBounty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	bountyStr := params["bounty"]

	// Convert bounty string to int
	bounty, err := strconv.Atoi(bountyStr)
	if err != nil {
		http.Error(w, "Invalid bounty value", http.StatusBadRequest)
		return
	}

	var matchingMugiwaras []Mugiwara

	for _, mugiwara := range Mugiwaras {
		if mugiwara.Bounty == bounty {
			matchingMugiwaras = append(matchingMugiwaras, mugiwara)
		}
	}

	if len(matchingMugiwaras) > 0 {
		json.NewEncoder(w).Encode(matchingMugiwaras)
	} else {
		http.Error(w, "Mugiwaras with the specified bounty not found", http.StatusNotFound)
	}
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	healthStatus := "App is healthy!\n\n"

	appDescription := "This is a simple web API providing information about Straw Hat Pirates from manga named One piece. It allows users to retrieve characters details, perform searches, and check the overall health of the application.\n"
	authorInfo := "\n\nAuthor: Kurmanbek Tolebayev(justkurama)\n"

	healthCheckResponse := healthStatus + appDescription + authorInfo

	w.Write([]byte(healthCheckResponse))
}
