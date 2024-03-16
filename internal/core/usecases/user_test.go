package usecases

import (
	"testing"

	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/pkg/id"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	repo := &mockUserRepository{}
	useCase := NewUserUseCase(repo)

	repo.On("Create", mock.AnythingOfType("*entity.User")).Return(nil)

	input := &dto.CreateUserInputDTO{
		Name:     "name",
		Email:    "email",
		Password: "password",
	}

	user, err := useCase.CreateUser(input)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user == nil {
		t.Error("Expected user, got nil")
	}
	repo.AssertCalled(t, "Create", mock.AnythingOfType("*entity.User"))
}

func TestFindUserByEmail(t *testing.T) {
	repo := &mockUserRepository{}
	useCase := NewUserUseCase(repo)

	user := &entity.User{}
	repo.On("FindByEmail", "email").Return(user, nil)

	result, err := useCase.FindUserByEmail("email")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if result == nil {
		t.Error("Expected user, got nil")
	}
	repo.AssertCalled(t, "FindByEmail", "email")
}

func TestUpdateUser(t *testing.T) {
	repo := &mockUserRepository{}
	useCase := NewUserUseCase(repo)

	repo.On("Update", mock.Anything).Return(nil)

	input := &dto.UpdateUserInputDTO{
		Name:     "name",
		Email:    "email",
		Password: "password",
	}

	user, err := useCase.UpdateUser(id.NewID().String(), input)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user == nil {
		t.Error("Expected user, got nil")
	}
	repo.AssertCalled(t, "Update", mock.AnythingOfType("*entity.User"))
}

func TestDeleteUser(t *testing.T) {
	repo := &mockUserRepository{}
	useCase := NewUserUseCase(repo)

	repo.On("Delete", mock.Anything).Return(nil)

	err := useCase.DeleteUser(id.NewID().String())
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	repo.AssertCalled(t, "Delete", mock.AnythingOfType("uuid.UUID"))
}

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Create(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepository) FindByEmail(email string) (*entity.User, error) {
	args := m.Called(email)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (m *mockUserRepository) Update(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepository) Delete(id id.ID) error {
	args := m.Called(id)
	return args.Error(0)
}
