package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/models"
)

func GetRace(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	idStr, ok := params["id"]
	if !ok {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if id <= 0 {
		http.Error(w, "Invalid ID. ID must be greater than 0", http.StatusBadRequest)
		return
	}

	var race models.Race
	err = db.First(&race, id).Error
	if err != nil {
		http.Error(w, "Race not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(race)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
func GetAllRaces(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	var races []models.Race
	err := db.Find(&races).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(races)
	if err != nil {
		return
	}
	return
}
