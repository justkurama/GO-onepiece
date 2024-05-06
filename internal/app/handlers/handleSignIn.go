package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/justkurama/GO-onepiece/internal/app/models"
	"github.com/justkurama/GO-onepiece/internal/pkg/utils"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var credentials struct {
		Login    string
		Password string
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.Where("login = ?", credentials.Login).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if !utils.CheckPasswordHash(credentials.Password, user.Password) {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.GenerateToken(user.Login)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := utils.GenerateRefreshToken(user.Login)
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	tokens := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	json.NewEncoder(w).Encode(tokens)
}
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Extract the token from the Authorization header
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		http.Error(w, "Authorization header is missing", http.StatusBadRequest)
		return
	}

	// Optionally strip the 'Bearer ' prefix if you expect it
	splitToken := strings.Split(authToken, "Bearer ")
	if len(splitToken) != 2 {
		http.Error(w, "Invalid Authorization token format", http.StatusBadRequest)
		return
	}
	refreshToken := splitToken[1]

	// Validate the refresh token
	token, err := utils.ValidateRefreshToken(refreshToken)
	if err != nil || !token.Valid {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}

	// Extract claims
	claims, ok := token.Claims.(*utils.Claims)
	if !ok || claims.Login == "" { // Also check that Username is non-empty
		http.Error(w, "Invalid token claims", http.StatusInternalServerError)
		return
	}

	// Generate new tokens
	newAccessToken, err := utils.GenerateToken(claims.Login)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	newRefreshToken, err := utils.GenerateRefreshToken(claims.Login)
	if err != nil {
		http.Error(w, "Failed to generate refresh token", http.StatusInternalServerError)
		return
	}

	tokens := map[string]string{
		"accessToken":  newAccessToken,
		"refreshToken": newRefreshToken,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(tokens); err != nil {
		http.Error(w, "Failed to encode tokens", http.StatusInternalServerError)
	}
}
