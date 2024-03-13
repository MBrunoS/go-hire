package repository

import (
	"github.com/mbrunos/go-hire/internal/entity"
	"github.com/mbrunos/go-hire/pkg/id"
)

type JobRepository interface {
	Create(job *entity.Job) error
	FindByID(id id.ID) (*entity.Job, error)
	FindAll() (*[]entity.Job, error)
	Update(job *entity.Job) error
	Delete(id id.ID) error
}
