package usecases

import (
	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/internal/core/repository"
	"github.com/mbrunos/go-hire/pkg/id"
)

type UserUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) *UserUseCase {
	return &UserUseCase{repository: userRepository}
}

func (u *UserUseCase) CreateUser(name, email, password string) (*entity.User, error) {
	user, err := entity.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}

	if err = u.repository.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) FindUserByEmail(email string) (*entity.User, error) {
	return u.repository.FindByEmail(email)
}

func (u *UserUseCase) UpdateUser(idStr, name, email, password string) (*entity.User, error) {
	id, err := id.StringToID(idStr)
	if err != nil {
		return nil, err
	}

	user, err := entity.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}
	user.ID = id

	if err = u.repository.Update(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUseCase) DeleteUser(idStr string) error {
	id, err := id.StringToID(idStr)
	if err != nil {
		return err
	}

	return u.repository.Delete(id)
}
