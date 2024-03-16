package usecases

import (
	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/internal/core/entity/interfaces"
	"github.com/mbrunos/go-hire/pkg/id"
)

type JobUseCase struct {
	repository interfaces.JobRepository
}

func NewJobUseCase(jobRepository interfaces.JobRepository) *JobUseCase {
	return &JobUseCase{repository: jobRepository}
}

func (u *JobUseCase) CreateJob(input *dto.CreateJobInputDTO) (*dto.JobOutputDTO, error) {
	job := entity.NewJob(input.Title, input.Description, input.Company, input.Location, input.Remote, input.Salary)

	if err := u.repository.Create(job); err != nil {
		return nil, err
	}

	return &dto.JobOutputDTO{
		ID:          job.ID.String(),
		Title:       job.Title,
		Description: job.Description,
		Company:     job.Company,
		Location:    job.Location,
		Remote:      job.Remote,
		Salary:      job.Salary,
		CreatedAt:   job.CreatedAt.String(),
		UpdatedAt:   job.UpdatedAt.String(),
	}, nil
}

func (u *JobUseCase) FindJobByID(idStr string) (*dto.JobOutputDTO, error) {
	id, err := id.StringToID(idStr)
	if err != nil {
		return nil, err
	}

	job, err := u.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	return &dto.JobOutputDTO{
		ID:          job.ID.String(),
		Title:       job.Title,
		Description: job.Description,
		Company:     job.Company,
		Location:    job.Location,
		Remote:      job.Remote,
		Salary:      job.Salary,
		CreatedAt:   job.CreatedAt.String(),
		UpdatedAt:   job.UpdatedAt.String(),
	}, nil
}

func (u *JobUseCase) FindAllJobs(page, limit int, sortField, sortDir string) (*dto.JobListOutputDTO, error) {
	jobs, err := u.repository.FindAll(page, limit, sortField, sortDir)

	if err != nil {
		return nil, err
	}

	var jobsOutput []dto.JobOutputDTO

	for _, job := range *jobs {
		jobsOutput = append(jobsOutput, dto.JobOutputDTO{
			ID:          job.ID.String(),
			Title:       job.Title,
			Description: job.Description,
			Company:     job.Company,
			Location:    job.Location,
			Remote:      job.Remote,
			Salary:      job.Salary,
			CreatedAt:   job.CreatedAt.String(),
			UpdatedAt:   job.UpdatedAt.String(),
		})
	}

	return &dto.JobListOutputDTO{
		Jobs: jobsOutput,
	}, nil
}

func (u *JobUseCase) UpdateJob(idStr string, input *dto.UpdateJobInputDTO) (*dto.JobOutputDTO, error) {
	id, err := id.StringToID(idStr)
	if err != nil {
		return nil, err
	}

	job := entity.NewJob(input.Title, input.Description, input.Company, input.Location, input.Remote, input.Salary)
	job.ID = id

	if err := u.repository.Update(job); err != nil {
		return nil, err
	}

	return &dto.JobOutputDTO{
		ID:          job.ID.String(),
		Title:       job.Title,
		Description: job.Description,
		Company:     job.Company,
		Location:    job.Location,
		Remote:      job.Remote,
		Salary:      job.Salary,
		CreatedAt:   job.CreatedAt.String(),
		UpdatedAt:   job.UpdatedAt.String(),
	}, nil
}

func (u *JobUseCase) DeleteJob(idStr string) error {
	id, err := id.StringToID(idStr)
	if err != nil {
		return err
	}
	return u.repository.Delete(id)
}
