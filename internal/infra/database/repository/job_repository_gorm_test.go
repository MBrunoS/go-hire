package repository_test

import (
	"testing"

	"github.com/mbrunos/go-hire/internal/core/entity"
	"github.com/mbrunos/go-hire/internal/infra/database/repository"
	"github.com/mbrunos/go-hire/internal/infra/database/schema"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateJob(t *testing.T) {
	repo, job, db := setupJobRepo()

	err := repo.Create(job)
	assert.Nil(t, err)

	var j entity.Job
	err = db.First(&j, job.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, job.ID, j.ID)
	assert.Equal(t, job.Title, j.Title)
	assert.Equal(t, job.Description, j.Description)
	assert.Equal(t, job.Company, j.Company)
	assert.Equal(t, job.Salary, j.Salary)
	assert.Equal(t, job.Remote, j.Remote)
	assert.Nil(t, j.Location)
}

func TestFindByID(t *testing.T) {
	repo, job, _ := setupJobRepo()

	err := repo.Create(job)
	assert.Nil(t, err)

	j, err := repo.FindByID(job.ID)
	assert.Nil(t, err)
	assert.Equal(t, job.ID, j.ID)
	assert.Equal(t, job.Title, j.Title)
	assert.Equal(t, job.Description, j.Description)
	assert.Equal(t, job.Company, j.Company)
	assert.Equal(t, job.Salary, j.Salary)
	assert.Equal(t, job.Remote, j.Remote)
	assert.Nil(t, j.Location)
}

func TestFindAllJobs(t *testing.T) {
	repo, job, _ := setupJobRepo()

	err := repo.Create(job)
	assert.Nil(t, err)

	jobs, err := repo.FindAll(1, 10, "created_at", "asc")
	assert.Nil(t, err)
	assert.Len(t, *jobs, 1)

	job2, _ := entity.NewJob("Senior Golang Developer", "We want top golang developers", "Google", nil, false, 100000)
	err = repo.Create(job2)
	assert.Nil(t, err)

	jobs, err = repo.FindAll(1, 10, "created_at", "asc")
	assert.Nil(t, err)
	assert.Len(t, *jobs, 2)
}

func TestUpdateJob(t *testing.T) {
	repo, job, db := setupJobRepo()

	err := repo.Create(job)
	assert.Nil(t, err)

	job.Title = "New Title"
	job.Description = "New Description"
	job.Company = "New Company"
	job.Salary = 200000
	job.Remote = true

	err = repo.Update(job)
	assert.Nil(t, err)

	var j entity.Job
	err = db.First(&j, job.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, job.ID, j.ID)
	assert.Equal(t, job.Title, j.Title)
	assert.Equal(t, job.Description, j.Description)
	assert.Equal(t, job.Company, j.Company)
	assert.Equal(t, job.Salary, j.Salary)
	assert.Equal(t, job.Remote, j.Remote)
	assert.Nil(t, j.Location)
}

func TestDeleteJob(t *testing.T) {
	repo, job, db := setupJobRepo()

	err := repo.Create(job)
	assert.Nil(t, err)

	err = repo.Delete(job.ID)
	assert.Nil(t, err)

	var j entity.Job
	err = db.First(&j, job.ID).Error
	assert.NotNil(t, err)
}

func setupJobRepo() (*repository.JobRepository, *entity.Job, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&schema.Job{})

	job, _ := entity.NewJob("Senior Golang Developer", "We want top golang developers", "Google", nil, false, 100000)
	repo := repository.NewJobRepository(db)

	return repo, job, db
}
