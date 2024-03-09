package handler

import (
	"fmt"
	"net/http"

	"github.com/mbrunos/go-hire/schemas"
)

func DeleteJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	job := schemas.Job{}

	if err := db.First(&job, id).Error; err != nil {
		sendError(w, http.StatusNotFound, fmt.Errorf("job with id %s not found", id))
		return
	}

	if err := db.Delete(&job).Error; err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("error deleting job with id %s", id))
		return
	}

	sendSuccess(w, http.StatusOK, job)
}
