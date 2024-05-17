package handlers

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/justkurama/GO-onepiece/internal/app/models"
	"gorm.io/gorm"
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

	var organization models.Organization
	err = db.First(&organization, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Organization not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(organization); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
func GetOrganizationCharacters(w http.ResponseWriter, r *http.Request) {
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

	type OrganizationInfo struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	var orgInfo OrganizationInfo
	err = db.Table("organizations").Select("id, name").Where("id = ?", id).Scan(&orgInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Organization not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var characters []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	err = db.Table("characters").Select("id, name").Where("organization_id = ?", id).Scan(&characters).Error
	if err != nil {
		http.Error(w, "Failed to retrieve characters", http.StatusInternalServerError)
		return
	}

	type OrganizationCharactersResponse struct {
		Organization OrganizationInfo `json:"organization"`
		Characters   []struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"characters"`
	}

	response := OrganizationCharactersResponse{
		Organization: orgInfo,
		Characters:   characters,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetOrganizationSubOrg(w http.ResponseWriter, r *http.Request) {
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

	type OrganizationInfo struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}

	var orgInfo OrganizationInfo
	err = db.Table("organizations").Select("id, name").Where("id = ?", id).Scan(&orgInfo).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Organization not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal server error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	var subOrgs []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	}
	err = db.Table("sub_organizations").Select("id, name").Where("parent_id = ?", id).Scan(&subOrgs).Error
	if err != nil {
		http.Error(w, "Failed to retrieve suborganizations", http.StatusInternalServerError)
		return
	}

	type OrganizationSubOrgResponse struct {
		Organization     OrganizationInfo `json:"organization"`
		SubOrganizations []struct {
			ID   uint   `json:"id"`
			Name string `json:"name"`
		} `json:"suborganizations"`
	}

	response := OrganizationSubOrgResponse{
		Organization:     orgInfo,
		SubOrganizations: subOrgs,
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func GetAllOrganizations(w http.ResponseWriter, r *http.Request) {
	w = SetContentType(w)

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
		sortBy = "id asc"
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

	var organizations []models.Organization
	err = db.Preload(clause.Associations).Order(sortBy).Offset(offset).Limit(limit).Find(&organizations).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte("Internal server error: " + err.Error()))
		if err != nil {
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)

	sortBy = r.URL.Query().Get("sortBy")
	if sortBy == "" {
		sortBy = "id.asc"
	}
	err = json.NewEncoder(w).Encode(organizations)
	if err != nil {
		http.Error(w, "Failed to encode organizations", http.StatusInternalServerError)
	}
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
