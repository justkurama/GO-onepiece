package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/models"
	"net/http"
	"strconv"
)

func GetOrigin(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var origin models.Origin
	err := db.First(&origin, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(origin)
	if err != nil {
		return
	}
	return
}
func GetAllOrigins(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	var origins []models.Origin
	err := db.Find(&origins).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(origins)
	if err != nil {
		return
	}
	return
}
