package handler

import (
	"net/http"

	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/schemas"
)

func GetAllJobs(w http.ResponseWriter, r *http.Request) {
	jobs := []schemas.Job{}

	db := config.GetDB()

	if err := db.Find(&jobs).Error; err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusOK, jobs)
}
