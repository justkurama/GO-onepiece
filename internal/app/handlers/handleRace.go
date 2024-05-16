package handlers

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
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

	var race models.Race
	err = db.First(&race, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Race not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(race); err != nil {
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
