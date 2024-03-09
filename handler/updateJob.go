package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/schemas"
)

func UpdateJob(w http.ResponseWriter, r *http.Request) {
	body := UpdateJobRequest{}
	json.NewDecoder(r.Body).Decode(&body)

	if err := body.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	job := schemas.Job{}

	id := r.PathValue("id")

	db := config.GetDB()
	if err := db.First(&job, id).Error; err != nil {
		sendError(w, http.StatusNotFound, err)
		return
	}

	if body.Title != "" {
		job.Title = body.Title
	}
	if body.Description != "" {
		job.Description = body.Description
	}
	if body.Company != "" {
		job.Company = body.Company
	}
	if body.Location != nil {
		job.Location = body.Location
	}
	if body.Level != "" {
		job.Level = body.Level
	}
	if body.Remote != nil {
		job.Remote = *body.Remote
	}
	if body.Salary > 0 {
		job.Salary = body.Salary
	}

	if err := db.Save(&job).Error; err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusOK, job)
}
