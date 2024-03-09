package handler

import "fmt"

type CreateJobRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Company     string  `json:"company"`
	Location    *string `json:"location"`
	Remote      *bool   `json:"remote"`
	Salary      int64   `json:"salary"`
}

func (c CreateJobRequest) Validate() error {
	if c.Title == "" && c.Description == "" && c.Company == "" && c.Remote == nil && c.Salary <= 0 {
		return fmt.Errorf("request body is empty or malformed")
	}

	if c.Title == "" {
		return errParamIsRequired("title", "string")
	}
	if c.Description == "" {
		return errParamIsRequired("description", "string")
	}
	if c.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if c.Remote == nil {
		return errParamIsRequired("remote", "bool")
	}
	if c.Salary <= 0 {
		return errParamIsRequired("salary", "int")
	}
	return nil
}

type UpdateJobRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Company     string  `json:"company"`
	Location    *string `json:"location"`
	Remote      *bool   `json:"remote"`
	Salary      int64   `json:"salary"`
}

func (c UpdateJobRequest) Validate() error {
	if c.Title != "" || c.Description != "" || c.Company != "" || c.Remote != nil || c.Salary > 0 {
		return nil
	}

	return fmt.Errorf("request body is empty or malformed")
}

func errParamIsRequired(param, typ string) error {
	return fmt.Errorf("param '%s' is required (type: %s)", param, typ)
}
