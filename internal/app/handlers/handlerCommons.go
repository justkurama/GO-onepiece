package handlers

import (
	"github.com/justkurama/GO-onepiece/internal/pkg/database"
	"net/http"
)

var db = database.DB

func SetContentType(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	return w
}
