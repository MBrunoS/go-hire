package entity

import (
	"errors"

	"github.com/mbrunos/go-hire/pkg/id"
)

type Job struct {
	ID          id.ID   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Company     string  `json:"company"`
	Location    *string `json:"location"`
	Remote      bool    `json:"remote"`
	Salary      int64   `json:"salary"`
}

func NewJob(title, description, company string, location *string, remote bool, salary int64) *Job {
	return &Job{
		ID:          id.NewID(),
		Title:       title,
		Description: description,
		Company:     company,
		Location:    location,
		Remote:      remote,
		Salary:      salary,
	}
}

func (j *Job) Validate() error {
	if id.IsNil(j.ID) || j.ID.String() == "" {
		return errors.New("id is required")
	}

	if _, err := id.StringToID(j.ID.String()); err != nil {
		return errors.New("id is not valid")
	}

	if j.Title == "" {
		return errors.New("title is required")
	}

	if j.Description == "" {
		return errors.New("description is required")
	}

	if j.Company == "" {
		return errors.New("company is required")
	}

	if j.Salary <= 0 {
		return errors.New("salary is not valid")
	}

	return nil
}
