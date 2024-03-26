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

// Create godoc
// @Summary 		Create Job
// @Description Create Job
// @Tags 				jobs
// @Accept 			json
// @Produce 		json
// @Param 			request body dto.CreateJobInputDTO true "Job data"
// @Success 		201 {object} dto.JobOutputDTO
// @Failure 		400 {object} dto.ErrorOutputDTO
// @Failure 		500 {object} dto.ErrorOutputDTO
// @Router 			/jobs [post]
// @Security 		ApiKeyAuth
func (h *JobHandler) Create(c *router.Context) {
	var input dto.CreateJobInputDTO
	if err := c.BindJSON(&input); err != nil {
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

// Get godoc
// @Summary 		Get Job
// @Description Get Job by ID
// @Tags 				jobs
// @Produce 		json
// @Param 			id path string true "Job ID"
// @Success 		200 {object} dto.JobOutputDTO
// @Failure 		404 {object} dto.ErrorOutputDTO
// @Router 			/jobs/{id} [get]
func (h *JobHandler) Get(c *router.Context) {
	id := c.PathParam("id")

	job, err := h.jobUseCase.FindJobByID(id)

	if err != nil {
		c.SendError(http.StatusNotFound, fmt.Errorf("job with id %s not found", id))
		return
	}

	c.SendJSON(http.StatusOK, job)
}

// List godoc
// @Summary List 	Jobs
// @Description 	List Jobs
// @Tags 					jobs
// @Produce 			json
// @Success 			200 {array} dto.JobOutputDTO
// @Failure 			500 {object} dto.ErrorOutputDTO
// @Router 				/jobs [get]
func (h *JobHandler) List(c *router.Context) {
	jobs, err := h.jobUseCase.FindAllJobs(0, 0, "", "")

	if err != nil {
		c.SendError(http.StatusInternalServerError, err)
		return
	}

	c.SendJSON(http.StatusOK, jobs)
}

// Update godoc
// @Summary 		Update Job
// @Description Update Job
// @Tags 				jobs
// @Accept 			json
// @Produce 		json
// @Param 			id path string true "Job ID"
// @Param 			request body dto.UpdateJobInputDTO true "Job data"
// @Success 		200 {object} dto.JobOutputDTO
// @Failure 		400 {object} dto.ErrorOutputDTO
// @Failure 		500 {object} dto.ErrorOutputDTO
// @Router 			/jobs/{id} [put]
// @Security 		ApiKeyAuth
func (h *JobHandler) Update(c *router.Context) {
	id := c.PathParam("id")

	var input dto.UpdateJobInputDTO
	if err := c.BindJSON(&input); err != nil {
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

// Delete godoc
// @Summary 		Delete Job
// @Description Delete Job
// @Tags 				jobs
// @Produce 		json
// @Param 			id path string true "Job ID"
// @Success 		204
// @Failure 		404 {object} dto.ErrorOutputDTO
// @Failure 		500 {object} dto.ErrorOutputDTO
// @Router 			/jobs/{id} [delete]
// @Security 		ApiKeyAuth
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
