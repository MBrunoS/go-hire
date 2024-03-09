package handler

import (
	"fmt"
	"net/http"

	"github.com/mbrunos/go-hire/schemas"
)

// @BasePath /api
// @Summary Get job
// @Description Get a job by ID
// @Tags jobs
// @Accept json
// @Produce json
// @Param id path string true "Job ID"
// @Success 200 {object} JobSuccessResponse
// @Failure 404 {object} ErrorResponse
// @Router /api/jobs/{id} [get]
func GetJob(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	job := schemas.Job{}

	if err := db.First(&job, id).Error; err != nil {
		sendError(w, http.StatusNotFound, fmt.Errorf("job with id %s not found", id))
		return
	}

	sendSuccess(w, http.StatusOK, job)
}
