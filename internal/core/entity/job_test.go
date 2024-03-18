package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJob(t *testing.T) {
	r := true
	job, err := NewJob("Software Engineer", "Description", "Google", nil, &r, 100000)
	assert.Nil(t, err)
	assert.NotNil(t, job)
	assert.NotEmpty(t, job.ID)
	assert.Equal(t, "Software Engineer", job.Title)
	assert.Equal(t, "Description", job.Description)
	assert.Equal(t, "Google", job.Company)
	assert.Nil(t, job.Location)
	assert.True(t, *job.Remote)
	assert.Equal(t, int64(100000), job.Salary)

	job, err = NewJob("Software Engineer", "Description", "Google", nil, &r, 0)
	assert.ErrorContains(t, err, "salary is required and must be greater than 0")
	assert.Nil(t, job)
}

func TestJobValidate(t *testing.T) {
	r := true
	job, _ := NewJob("Software Engineer", "Description", "Google", nil, &r, 100000)
	assert.Nil(t, job.Validate())

	job = &Job{}
	assert.ErrorContains(t, job.Validate(), "id is required")

	job, _ = NewJob("", "Description", "Google", nil, &r, 100000)
	assert.ErrorContains(t, job.Validate(), "title is required")

	job, _ = NewJob("Software Engineer", "", "Google", nil, &r, 100000)
	assert.ErrorContains(t, job.Validate(), "description is required")

	job, _ = NewJob("Software Engineer", "Description", "", nil, &r, 100000)
	assert.ErrorContains(t, job.Validate(), "company is required")

	job, _ = NewJob("Software Engineer", "Description", "Google", nil, &r, 0)
	assert.ErrorContains(t, job.Validate(), "salary is not valid")
}
