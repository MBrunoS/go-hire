package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/entity/interfaces"
	"github.com/mbrunos/go-hire/internal/core/usecases"
)

type JobHandler struct {
	jobRepository interfaces.JobRepository
}

func NewJobHandler(jobRepository interfaces.JobRepository) *JobHandler {
	return &JobHandler{
		jobRepository: jobRepository,
	}
}

func (h *JobHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateJobInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	jobUseCase := usecases.NewJobUseCase(h.jobRepository)
	job, err := jobUseCase.CreateJob(&input)

	if err != nil {
		sendError(w, http.StatusInternalServerError, errors.New("error creating job"))
		return
	}

	sendSuccess(w, http.StatusCreated, job)
}

func (h *JobHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	jobUseCase := usecases.NewJobUseCase(h.jobRepository)
	job, err := jobUseCase.FindJobByID(id)

	if err != nil {
		sendError(w, http.StatusNotFound, fmt.Errorf("job with id %s not found", id))
		return
	}

	sendSuccess(w, http.StatusOK, job)
}

func (h *JobHandler) List(w http.ResponseWriter, r *http.Request) {
	jobUseCase := usecases.NewJobUseCase(h.jobRepository)
	jobs, err := jobUseCase.FindAllJobs(0, 0, "", "")

	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusOK, jobs)
}

func (h *JobHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var input dto.UpdateJobInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	jobUseCase := usecases.NewJobUseCase(h.jobRepository)
	job, err := jobUseCase.UpdateJob(id, &input)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusOK, job)
}

func (h *JobHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	jobUseCase := usecases.NewJobUseCase(h.jobRepository)
	_, err := jobUseCase.FindJobByID(id)

	if err != nil {
		sendError(w, http.StatusNotFound, fmt.Errorf("job with id %s not found", id))
		return
	}

	if err := jobUseCase.DeleteJob(id); err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("error deleting job with id %s", id))
		return
	}

	sendSuccess(w, http.StatusNoContent, nil)
}
