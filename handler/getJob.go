package handler

import (
	"encoding/json"
	"net/http"
)

func GetJob(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Get job by id")
}
