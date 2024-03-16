package interfaces

import (
	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/pkg/id"
)

type JobRepository interface {
	Create(job *entity.Job) error
	FindByID(id id.ID) (*entity.Job, error)
	FindAll(page, limit int, sortField, sortDir string) (*[]entity.Job, error)
	Update(job *entity.Job) error
	Delete(id id.ID) error
}
