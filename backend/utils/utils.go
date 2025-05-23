package utils

import (
	"net/http"
)


func SetHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}