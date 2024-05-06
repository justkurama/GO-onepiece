package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/models"
	"gorm.io/gorm/clause"
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

	// Handling pagination
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	// Handling sorting
	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		sortBy = "id asc" // Default sort
	} else {
		allowedSortFields := map[string]bool{"id": true, "name": true}
		sortFields := strings.Fields(sortBy)

		if len(sortFields) == 1 {
			sortFields = append(sortFields, "asc")
		}

		if len(sortFields) != 2 || !allowedSortFields[sortFields[0]] || (sortFields[1] != "asc" && sortFields[1] != "desc") {
			http.Error(w, "Invalid sortBy parameter", http.StatusBadRequest)
			return
		}
		sortBy = strings.Join(sortFields, " ")
	}

	offset := (page - 1) * limit

	var characters []models.Character
	err = db.Preload(clause.Associations).Order(sortBy).Offset(offset).Limit(limit).Find(&characters).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Internal server error: " + err.Error()))
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(characters)
	if err != nil {
		http.Error(w, "Failed to encode characters", http.StatusInternalServerError)
	}
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
