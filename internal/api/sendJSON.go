package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// sendJSONError преобразует сообщение об ошибки в JSON формат и отправляет его
func sendJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	writeJSON(w, map[string]string{"error": message}, statusCode)
}

// writeJSON преобразует сообщение в JSON формат и отправляет его
func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("JSON encoding failed: %v", err)
	}
}
