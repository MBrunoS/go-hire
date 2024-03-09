package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mbrunos/go-hire/config"
	"github.com/mbrunos/go-hire/schemas"
)

func CreateJob(w http.ResponseWriter, r *http.Request) {
	body := CreateJobRequest{}
	json.NewDecoder(r.Body).Decode(&body)

	if err := body.Validate(); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	job := schemas.Job{
		Title:       body.Title,
		Description: body.Description,
		Company:     body.Company,
		Location:    body.Location,
		Level:       body.Level,
		Remote:      *body.Remote,
		Salary:      body.Salary,
	}

	db := config.GetDB()
	if err := db.Create(&job).Error; err != nil {
		sendError(w, http.StatusInternalServerError, errors.New("error creating job"))
		return
	}

	sendSuccess(w, http.StatusCreated, job)
}
