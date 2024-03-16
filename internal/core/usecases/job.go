package usecases

import (
	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/internal/core/repository"
	"github.com/mbrunos/go-hire/pkg/id"
)

type JobUseCase struct {
	repository interfaces.JobRepository
}

func NewJobUseCase(jobRepository interfaces.JobRepository) *JobUseCase {
	return &JobUseCase{repository: jobRepository}
}

func (u *JobUseCase) CreateJob(title, description, company string, location *string, remote bool, salary int64) (*entity.Job, error) {
	job := entity.NewJob(title, description, company, location, remote, salary)
	if err := u.repository.Create(job); err != nil {
		return nil, err
	}
	return job, nil
}

func (u *JobUseCase) FindJobByID(idStr string) (*entity.Job, error) {
	id, err := id.StringToID(idStr)
	if err != nil {
		return nil, err
	}
	return u.repository.FindByID(id)
}

func (u *JobUseCase) FindAllJobs(page, limit int, sort string) (*[]entity.Job, error) {
	return u.repository.FindAll(page, limit, sort)
}

func (u *JobUseCase) UpdateJob(idStr, title, description, company string, location *string, remote bool, salary int64) (*entity.Job, error) {
	id, err := id.StringToID(idStr)
	if err != nil {
		return nil, err
	}

	job := entity.NewJob(title, description, company, location, remote, salary)
	job.ID = id

	if err := u.repository.Update(job); err != nil {
		return nil, err
	}

	return job, nil
}

func (u *JobUseCase) DeleteJob(idStr string) error {
	id, err := id.StringToID(idStr)
	if err != nil {
		return err
	}
	return u.repository.Delete(id)
}
