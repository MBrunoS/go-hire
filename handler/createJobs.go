package handler

import (
	"encoding/json"
	"net/http"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Create job")
}
