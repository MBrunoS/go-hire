package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/usecases"
)

type JobHandler struct {
	jobUseCase *usecases.JobUseCase
}

func NewJobHandler(usecase *usecases.JobUseCase) *JobHandler {
	return &JobHandler{
		jobUseCase: usecase,
	}
}

func (h *JobHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateJobInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		sendError(w, http.StatusBadRequest, err)
		return
	}

	job, err := h.jobUseCase.CreateJob(&input)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusCreated, job)
}

func (h *JobHandler) Get(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	job, err := h.jobUseCase.FindJobByID(id)

	if err != nil {
		sendError(w, http.StatusNotFound, fmt.Errorf("job with id %s not found", id))
		return
	}

	sendSuccess(w, http.StatusOK, job)
}

func (h *JobHandler) List(w http.ResponseWriter, r *http.Request) {
	jobs, err := h.jobUseCase.FindAllJobs(0, 0, "", "")

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

	job, err := h.jobUseCase.UpdateJob(id, &input)

	if err != nil {
		sendError(w, http.StatusInternalServerError, err)
		return
	}

	sendSuccess(w, http.StatusOK, job)
}

func (h *JobHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	_, err := h.jobUseCase.FindJobByID(id)

	if err != nil {
		sendError(w, http.StatusNotFound, fmt.Errorf("job with id %s not found", id))
		return
	}

	if err := h.jobUseCase.DeleteJob(id); err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Errorf("error deleting job with id %s", id))
		return
	}

	sendSuccess(w, http.StatusNoContent, nil)
}
