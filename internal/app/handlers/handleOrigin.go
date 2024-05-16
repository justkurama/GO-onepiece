package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetOrigin(w http.ResponseWriter, r *http.Request) {
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

	var origin models.Origin
	err = db.First(&origin, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Origin not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(origin); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetOriginCharacters(w http.ResponseWriter, r *http.Request) {
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

	type OriginInfo struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	var originInfo OriginInfo
	err = db.Table("origins").Select("id, name").Where("id = ?", id).Scan(&originInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Origin not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var characters []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	err = db.Table("characters").Select("id, name").Where("origin_id = ?", id).Scan(&characters).Error
	if err != nil {
		http.Error(w, "Failed to retrieve characters", http.StatusInternalServerError)
		return
	}

	type OriginCharactersResponse struct {
		Origin     OriginInfo `json:"origin"`
		Characters []struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"characters"`
	}

	response := OriginCharactersResponse{
		Origin:     originInfo,
		Characters: characters,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
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
