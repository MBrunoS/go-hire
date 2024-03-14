package database

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

	if err := r.db.First(&user, id.String()).Error; err != nil {
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

func (r *UserRepository) FindAll() ([]*entity.User, error) {
	users := []*entity.User{}

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Update(user *entity.User) error {
	u := entity.User{}

	if err := r.db.First(&u, user.ID.String()).Error; err != nil {
		return err
	}

	u.Name = user.Name
	u.Email = user.Email
	u.Password = user.Password

	if err := r.db.Save(&u).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id id.ID) error {
	return r.db.Delete(&entity.User{}, id.String()).Error
}
