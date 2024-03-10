package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/models"
	"net/http"
	"strconv"
)

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	var organization models.Organization
	err := json.NewDecoder(r.Body).Decode(&organization)
	err = db.Create(&organization).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Organization created"))
	if err != nil {
		return
	}
	return
}
func GetOrganization(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var organization models.Organization
	err := db.First(&organization, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(organization)
	if err != nil {
		return
	}
	return
}
func GetAllOrganizations(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	var organizations []models.Organization
	err := db.Find(&organizations).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(organizations)
	if err != nil {
		return
	}
	return
}
func UpdateOrganization(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	var organization models.Organization
	err := json.NewDecoder(r.Body).Decode(&organization)
	id, _ := strconv.Atoi(params["id"])
	organization.ID = uint(id)
	err = db.First(&organization, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	err = db.Save(&organization).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Organization updated"))
	if err != nil {
		return
	}
	return
}
func DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var organization models.Organization
	err := db.First(&organization, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	err = db.Delete(&organization).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Organization deleted"))
	if err != nil {
		return
	}
	return
}
