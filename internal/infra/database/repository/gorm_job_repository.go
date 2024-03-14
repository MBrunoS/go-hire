package repository

import (
	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/pkg/id"
	"gorm.io/gorm"
)

type JobRepository struct {
	db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{db}
}

func (r *JobRepository) Create(job *entity.Job) error {
	return r.db.Create(job).Error
}

func (r *JobRepository) FindByID(id id.ID) (*entity.Job, error) {
	job := entity.Job{}

	if err := r.db.First(&job, id).Error; err != nil {
		return nil, err
	}

	return &job, nil
}

func (r *JobRepository) FindAll(page, limit int, sortField, sortDir string) (*[]entity.Job, error) {
	if sortDir == "" && sortDir != "asc" && sortDir != "desc" {
		sortDir = "asc"
	}

	if sortField == "" {
		sortField = "created_at"
	}

	jobs := []entity.Job{}
	var err error

	if page > 0 && limit > 0 {
		err = r.db.Limit(limit).Offset((page - 1) * limit).Order(sortField + " " + sortDir).Find(&jobs).Error
	} else {
		err = r.db.Order(sortField + " " + sortDir).Find(&jobs).Error
	}

	return &jobs, err
}

func (r *JobRepository) Update(job *entity.Job) error {
	_, err := r.FindByID(job.ID)
	if err != nil {
		return err
	}

	return r.db.Save(&job).Error
}

func (r *JobRepository) Delete(id id.ID) error {
	return r.db.Delete(&entity.Job{}, id).Error
}
