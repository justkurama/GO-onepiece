package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/models"
	"net/http"
	"strconv"
)

func GetRace(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var race models.Race
	err := db.First(&race, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(race)
	if err != nil {
		return
	}
	return
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
