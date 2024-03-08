package handler

import (
	"encoding/json"
	"net/http"
)

func GetAllJobs(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Get all jobs")
}
