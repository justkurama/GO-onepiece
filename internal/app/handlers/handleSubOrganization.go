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

func CreateSubOrganization(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	var subOrganization models.SubOrganization
	err := json.NewDecoder(r.Body).Decode(&subOrganization)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = db.Create(&subOrganization).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("SubOrganization created"))
	if err != nil {
		return
	}
}

func GetSubOrganization(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)

	// Extract the sub-organization ID from the URL path
	vars := mux.Vars(r)
	subOrgIDStr := vars["id"]
	subOrgID, err := strconv.ParseUint(subOrgIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid sub-organization ID", http.StatusBadRequest)
		return
	}

	var subOrganization models.SubOrganization
	if err := db.First(&subOrganization, subOrgID).Error; err != nil {
		http.Error(w, "Sub-organization not found", http.StatusNotFound)
		return
	}

	// Encode the sub-organization as JSON and send it in the response
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(subOrganization); err != nil {
		http.Error(w, "Failed to encode sub-organization", http.StatusInternalServerError)
		return
	}
}

func GetSubOrganizationCharacters(w http.ResponseWriter, r *http.Request) {
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

	type SubOrganizationInfo struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	var subOrgInfo SubOrganizationInfo
	err = db.Table("sub_organizations").Select("id, name").Where("id = ?", id).Scan(&subOrgInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Sub-organization not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var characters []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	err = db.Table("characters").Select("id, name").Where("sub_organization_id = ?", id).Scan(&characters).Error
	if err != nil {
		http.Error(w, "Failed to retrieve characters", http.StatusInternalServerError)
		return
	}

	type SubOrganizationCharactersResponse struct {
		SubOrganization SubOrganizationInfo `json:"sub_organization"`
		Characters      []struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"characters"`
	}

	response := SubOrganizationCharactersResponse{
		SubOrganization: subOrgInfo,
		Characters:      characters,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
func GetAllSubOrganizations(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	var subOrganizations []models.SubOrganization
	err := db.Find(&subOrganizations).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(subOrganizations)
	if err != nil {
		return
	}
}

func UpdateSubOrganization(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var subOrganization models.SubOrganization
	// Check if sub-organization exists
	err = db.First(&subOrganization, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Sub-organization not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Decode the request body into the subOrganization struct
	err = json.NewDecoder(r.Body).Decode(&subOrganization)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	// Ensure the ID in the struct is set correctly
	subOrganization.ID = uint(id)

	// Save the updated sub-organization
	err = db.Save(&subOrganization).Error
	if err != nil {
		http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Sub-organization updated"))
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}

func DeleteSubOrganization(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var suborganization models.SubOrganization
	err := db.First(&suborganization, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := w.Write([]byte("Internal server error"))
		if err != nil {
			return
		}
		return
	}
	err = db.Delete(&suborganization).Error
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
