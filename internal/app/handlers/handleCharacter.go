package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/models"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	var character models.Character
	err := json.NewDecoder(r.Body).Decode(&character)
	err = db.Create(&character).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Character created"))
	if err != nil {
		return
	}
	return
}
func GetCharacter(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var character models.Character
	err := db.Preload(clause.Associations).First(&character, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(character)
	if err != nil {
		return
	}
	return
}
func GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	var characters []models.Character
	err := db.Preload(clause.Associations).Find(&characters).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)

	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		// id.asc is the default sort query
		sortBy = "id.asc"
	}
	err = json.NewEncoder(w).Encode(characters)
	if err != nil {
		return
	}
	return
}
func UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var character models.Character
	err := json.NewDecoder(r.Body).Decode(&character)
	err = db.First(&character, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	err = db.Save(&character).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Character updated"))
	if err != nil {
		return
	}
	return
}
func DeleteCharacter(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var character models.Character
	err := db.First(&character, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	err = db.Delete(&character).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Character deleted"))
	if err != nil {
		return
	}
	return
}
