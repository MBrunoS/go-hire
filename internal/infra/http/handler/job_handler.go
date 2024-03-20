package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/usecases"
	"github.com/mbrunos/go-hire/pkg/router"
)

type JobHandler struct {
	jobUseCase *usecases.JobUseCase
}

func NewJobHandler(usecase *usecases.JobUseCase) *JobHandler {
	return &JobHandler{
		jobUseCase: usecase,
	}
}

func (h *JobHandler) Create(c *router.Context) {
	var input dto.CreateJobInputDTO
	if err := c.BindJSON(input); err != nil {
		c.SendError(http.StatusBadRequest, errors.New("request body is empty or malformed"))
		return
	}

	job, err := h.jobUseCase.CreateJob(&input)

	if err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}
	c.SendJSON(http.StatusCreated, job)
}

func (h *JobHandler) Get(c *router.Context) {
	id := c.PathParam("id")

	job, err := h.jobUseCase.FindJobByID(id)

	if err != nil {
		c.SendError(http.StatusNotFound, fmt.Errorf("job with id %s not found", id))
		return
	}

	c.SendJSON(http.StatusOK, job)
}

func (h *JobHandler) List(c *router.Context) {
	jobs, err := h.jobUseCase.FindAllJobs(0, 0, "", "")

	if err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}

	c.SendJSON(http.StatusOK, jobs)
}

func (h *JobHandler) Update(c *router.Context) {
	id := c.PathParam("id")

	var input dto.UpdateJobInputDTO
	if err := c.BindJSON(input); err != nil {
		c.SendError(http.StatusBadRequest, err)
		return
	}

	job, err := h.jobUseCase.UpdateJob(id, &input)

	if err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}

	c.SendJSON(http.StatusOK, job)
}

func (h *JobHandler) Delete(c *router.Context) {
	id := c.PathParam("id")

	_, err := h.jobUseCase.FindJobByID(id)

	if err != nil {
		c.SendError(http.StatusNotFound, fmt.Errorf("job with id %s not found", id))
		return
	}

	if err := h.jobUseCase.DeleteJob(id); err != nil {
		c.SendError(http.StatusInternalServerError, fmt.Errorf("error deleting job with id %s", id))
		return
	}

	c.SendJSON(http.StatusNoContent, nil)
}
