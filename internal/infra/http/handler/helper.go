package handler

import (
	"encoding/json"
	"net/http"
)

func sendError(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(`{"error": "` + err.Error() + `"}`))
}

func sendSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	responseData := map[string]interface{}{
		"data": data,
	}
	json.NewEncoder(w).Encode(responseData)
}
