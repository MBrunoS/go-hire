package handler

import (
	"encoding/json"
	"net/http"
)

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Delete job by id")
}
