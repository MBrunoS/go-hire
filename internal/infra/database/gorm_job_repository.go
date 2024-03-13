package database

import (
	"github.com/mbrunos/go-hire/internal/entity"
	"github.com/mbrunos/go-hire/pkg/id"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository() *JobRepository {
	return &JobRepository{}
}

func (r *JobRepository) Create(job *entity.Job) error {
	return r.db.Create(job).Error
}

func (r *JobRepository) FindByID(id id.ID) (*entity.Job, error) {
	job := entity.Job{}

	if err := r.db.First(&job, id.String()).Error; err != nil {
		return nil, err
	}

	return &job, nil
}

func (r *JobRepository) FindAll() (*[]entity.Job, error) {
	jobs := []entity.Job{}

	if err := r.db.Find(&jobs).Error; err != nil {
		return nil, err
	}

	return &jobs, nil
}

func (r *JobRepository) Update(job *entity.Job) error {
	j := entity.Job{}
	if err := r.db.First(&j, job.ID.String()).Error; err != nil {
		return err
	}

	j.Company = job.Company
	j.Description = job.Description
	j.Location = job.Location
	j.Remote = job.Remote
	j.Salary = job.Salary
	j.Title = job.Title

	if err := r.db.Save(&j).Error; err != nil {
		return err
	}

	return nil
}

func (r *JobRepository) Delete(id id.ID) error {
	return r.db.Delete(&entity.Job{}, id.String()).Error
}
