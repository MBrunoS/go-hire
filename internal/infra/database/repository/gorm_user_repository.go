package repository

import (
	"github.com/mbrunos/go-hire/internal/entity"
	"github.com/mbrunos/go-hire/pkg/id"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id id.ID) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := entity.User{}

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	_, err := r.FindByID(user.ID)
	if err != nil {
		return err
	}

	return r.db.Save(&user).Error
}

func (r *UserRepository) Delete(id id.ID) error {
	return r.db.Delete(&entity.User{}, id).Error
}
