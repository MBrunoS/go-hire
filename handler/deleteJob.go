package handler

import (
	"fmt"
	"net/http"

	"github.com/mbrunos/go-hire/schemas"
)

// @BasePath /api
// @Summary Delete job
// @Description Delete an existing job
// @Tags jobs
// @Accept json
// @Produce json
// @Param id path string true "Job ID"
// @Success 200 {object} JobSuccessResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/jobs/{id} [delete]
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
