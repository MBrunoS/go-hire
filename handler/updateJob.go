package handler

import (
	"encoding/json"
	"net/http"
)

func UpdateJob(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Update job by id")
}
