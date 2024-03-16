package usecases

import (
	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/internal/core/entity/interfaces"
	"github.com/mbrunos/go-hire/pkg/id"
)

type UserUseCase struct {
	repository interfaces.UserRepository
}

func NewUserUseCase(userRepository interfaces.UserRepository) *UserUseCase {
	return &UserUseCase{repository: userRepository}
}

func (u *UserUseCase) CreateUser(input dto.CreateUserInputDTO) (*dto.UserOutputDTO, error) {
	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	if err = u.repository.Create(user); err != nil {
		return nil, err
	}

	return &dto.UserOutputDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserUseCase) FindUserByEmail(email string) (*dto.UserOutputDTO, error) {
	user, err := u.repository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return &dto.UserOutputDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserUseCase) UpdateUser(idStr string, input dto.UpdateUserInputDTO) (*dto.UserOutputDTO, error) {
	id, err := id.StringToID(idStr)
	if err != nil {
		return nil, err
	}

	user, err := entity.NewUser(input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	user.ID = id

	if err = u.repository.Update(user); err != nil {
		return nil, err
	}

	return &dto.UserOutputDTO{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (u *UserUseCase) DeleteUser(idStr string) error {
	id, err := id.StringToID(idStr)
	if err != nil {
		return err
	}

	return u.repository.Delete(id)
}
