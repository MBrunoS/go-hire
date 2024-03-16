package usecases

import (
	"testing"

	"github.com/mbrunos/go-hire/internal/core/dto"
	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/pkg/id"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateJob(t *testing.T) {
	repo := &mockJobRepository{}
	useCase := NewJobUseCase(repo)

	repo.On("Create", mock.AnythingOfType("*entity.Job")).Return(nil)

	input := &dto.CreateJobInputDTO{
		Title:       "title",
		Description: "description",
		Company:     "company",
		Remote:      true,
		Salary:      1000,
	}

	job, err := useCase.CreateJob(input)
	assert.Nil(t, err)
	assert.NotNil(t, job.ID)
	assert.Equal(t, "title", job.Title)
	assert.Equal(t, "description", job.Description)
	assert.Equal(t, "company", job.Company)
	assert.Nil(t, job.Location)
	assert.True(t, job.Remote)
	assert.Equal(t, int64(1000), job.Salary)
	assert.NotNil(t, job.CreatedAt)
	assert.NotNil(t, job.UpdatedAt)
	repo.AssertCalled(t, "Create", mock.AnythingOfType("*entity.Job"))
}

func TestFindJobByID(t *testing.T) {
	repo := &mockJobRepository{}
	useCase := NewJobUseCase(repo)

	job := entity.NewJob("title", "description", "company", nil, true, 1000)
	repo.On("FindByID", job.ID).Return(job, nil)

	result, err := useCase.FindJobByID(job.ID.String())
	assert.Nil(t, err)
	assert.NotNil(t, job.ID)
	assert.Equal(t, result.Company, job.Company)
	assert.Equal(t, result.Description, job.Description)
	assert.Equal(t, result.Title, job.Title)
	assert.Equal(t, result.Remote, job.Remote)
	assert.Equal(t, result.Salary, job.Salary)
	assert.NotNil(t, job.CreatedAt)
	assert.NotNil(t, job.UpdatedAt)
	repo.AssertCalled(t, "FindByID", job.ID)
}

func TestFindAllJobs(t *testing.T) {
	t.Run("should return empty list", func(t *testing.T) {
		repo := &mockJobRepository{}
		useCase := NewJobUseCase(repo)

		repo.On("FindAll", 1, 10, "created_at", "desc").Return(&[]entity.Job{}, nil)

		result, err := useCase.FindAllJobs(1, 10, "created_at", "desc")
		assert.Nil(t, err)
		assert.NotNil(t, result.Jobs)
		assert.Equal(t, 0, len(result.Jobs))
		repo.AssertCalled(t, "FindAll", 1, 10, "created_at", "desc")
	})

	t.Run("should return list with one job", func(t *testing.T) {
		repo := &mockJobRepository{}
		useCase := NewJobUseCase(repo)

		job := entity.NewJob("title", "description", "company", nil, true, 1000)
		repo.On("FindAll", 1, 10, "created_at", "desc").Return(&[]entity.Job{*job}, nil)

		result, err := useCase.FindAllJobs(1, 10, "created_at", "desc")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result.Jobs))
		assert.NotNil(t, result.Jobs[0].ID)
		assert.Equal(t, result.Jobs[0].Company, job.Company)
		assert.Equal(t, result.Jobs[0].Description, job.Description)
		assert.Equal(t, result.Jobs[0].Title, job.Title)
		assert.Equal(t, result.Jobs[0].Remote, job.Remote)
		assert.Equal(t, result.Jobs[0].Salary, job.Salary)
		assert.NotNil(t, result.Jobs[0].CreatedAt)
		assert.NotNil(t, result.Jobs[0].UpdatedAt)
		repo.AssertCalled(t, "FindAll", 1, 10, "created_at", "desc")
	})

}

func TestUpdateJob(t *testing.T) {
	repo := &mockJobRepository{}
	useCase := NewJobUseCase(repo)

	repo.On("Update", mock.Anything).Return(nil)

	input := &dto.UpdateJobInputDTO{
		Title:       "title",
		Description: "description",
		Company:     "company",
		Remote:      true,
		Salary:      1000,
	}

	job, err := useCase.UpdateJob(id.NewID().String(), input)
	assert.Nil(t, err)
	repo.AssertCalled(t, "Update", mock.Anything)
	assert.NotNil(t, job.ID)
	assert.Equal(t, "title", job.Title)
	assert.Equal(t, "description", job.Description)
	assert.Equal(t, "company", job.Company)
	assert.Nil(t, job.Location)
	assert.True(t, job.Remote)
	assert.Equal(t, int64(1000), job.Salary)
	assert.NotNil(t, job.CreatedAt)
	assert.NotNil(t, job.UpdatedAt)
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

func (m *mockJobRepository) FindAll(page, limit int, sortField, sortDir string) (*[]entity.Job, error) {
	args := m.Called(page, limit, sortField, sortDir)
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
