package handler

import (
	"net/http"

	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/schemas"
)

// @Summary Get jobs
// @Description Get all jobs
// @Tags jobs
// @Accept json
// @Produce json
// @Success 200 {object} GetJobsSuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/jobs [get]
func GetAllJobs(w http.ResponseWriter, r *http.Request) {
	jobs := []schemas.Job{}

	db := config.GetDB()

	if err := db.Find(&jobs).Error; err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusOK, jobs)
}
