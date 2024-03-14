package usecases

import (
	"testing"

	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/pkg/id"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateJob(t *testing.T) {
	repo := &mockJobRepository{}
	useCase := NewJobUseCase(repo)

	repo.On("Create", mock.AnythingOfType("*entity.Job")).Return(nil)

	job, err := useCase.CreateJob("title", "description", "company", nil, true, 1000)
	assert.Nil(t, err)
	assert.Equal(t, "title", job.Title)
	assert.Equal(t, "description", job.Description)
	assert.Equal(t, "company", job.Company)
	assert.Nil(t, job.Location)
	assert.True(t, job.Remote)
	assert.Equal(t, int64(1000), job.Salary)
	repo.AssertCalled(t, "Create", job)
}

func TestFindJobByID(t *testing.T) {
	repo := &mockJobRepository{}
	useCase := NewJobUseCase(repo)

	job := entity.NewJob("title", "description", "company", nil, true, 1000)
	repo.On("FindByID", job.ID).Return(job, nil)

	result, err := useCase.FindJobByID(job.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, job, result)
	repo.AssertCalled(t, "FindByID", job.ID)
}

func TestFindAllJobs(t *testing.T) {
	repo := &mockJobRepository{}
	useCase := NewJobUseCase(repo)

	job := entity.NewJob("title", "description", "company", nil, true, 1000)
	repo.On("FindAll", 1, 10, "created_at desc").Return(&[]entity.Job{*job}, nil)

	result, err := useCase.FindAllJobs(1, 10, "created_at desc")
	assert.Nil(t, err)
	assert.Equal(t, &[]entity.Job{*job}, result)
	repo.AssertCalled(t, "FindAll", 1, 10, "created_at desc")
}

func TestUpdateJob(t *testing.T) {
	repo := &mockJobRepository{}
	useCase := NewJobUseCase(repo)

	repo.On("Update", mock.Anything).Return(nil)

	job, err := useCase.UpdateJob(id.NewID().String(), "title", "description", "company", nil, true, int64(1000))
	assert.Nil(t, err)
	repo.AssertCalled(t, "Update", mock.Anything)
	assert.Equal(t, "title", job.Title)
	assert.Equal(t, "description", job.Description)
	assert.Equal(t, "company", job.Company)
	assert.Nil(t, job.Location)
	assert.True(t, job.Remote)
	assert.Equal(t, int64(1000), job.Salary)
}

func TestDeleteJob(t *testing.T) {
	repo := &mockJobRepository{}
	useCase := NewJobUseCase(repo)

	repo.On("Delete", mock.AnythingOfType("uuid.UUID")).Return(nil)

	err := useCase.DeleteJob(id.NewID().String())
	assert.Nil(t, err)
}

type mockJobRepository struct {
	mock.Mock
}

func (m *mockJobRepository) Create(job *entity.Job) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *mockJobRepository) FindByID(id id.ID) (*entity.Job, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Job), args.Error(1)
}

func (m *mockJobRepository) FindAll(page, limit int, sort string) (*[]entity.Job, error) {
	args := m.Called(page, limit, sort)
	return args.Get(0).(*[]entity.Job), args.Error(1)
}

func (m *mockJobRepository) Update(job *entity.Job) error {
	args := m.Called(job)
	return args.Error(0)
}

func (m *mockJobRepository) Delete(id id.ID) error {
	args := m.Called(id)
	return args.Error(0)
}
