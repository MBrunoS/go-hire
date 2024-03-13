package repository

import (
	"github.com/mbrunos/go-hire/internal/entity"
	"github.com/mbrunos/go-hire/pkg/id"
)

type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id id.ID) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindAll() ([]*entity.User, error)
	Update(user *entity.User) error
	Delete(id id.ID) error
}
