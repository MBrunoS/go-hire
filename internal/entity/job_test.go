package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	job := NewJob("Software Engineer", "Description", "Google", nil, true, 100000)
	assert.NotNil(t, job)
	assert.NotEmpty(t, job.ID)
	assert.Equal(t, "Software Engineer", job.Title)
	assert.Equal(t, "Description", job.Description)
	assert.Equal(t, "Google", job.Company)
	assert.Nil(t, job.Location)
	assert.True(t, job.Remote)
	assert.Equal(t, int64(100000), job.Salary)
}

func TestJobValidate(t *testing.T) {
	job := NewJob("Software Engineer", "Description", "Google", nil, true, 100000)
	assert.Nil(t, job.Validate())

	job = &Job{}
	assert.ErrorContains(t, job.Validate(), "id is required")

	job = NewJob("", "Description", "Google", nil, true, 100000)
	assert.ErrorContains(t, job.Validate(), "title is required")

	job = NewJob("Software Engineer", "", "Google", nil, true, 100000)
	assert.ErrorContains(t, job.Validate(), "description is required")

	job = NewJob("Software Engineer", "Description", "", nil, true, 100000)
	assert.ErrorContains(t, job.Validate(), "company is required")

	job = NewJob("Software Engineer", "Description", "Google", nil, true, 0)
	assert.ErrorContains(t, job.Validate(), "salary is not valid")
}
