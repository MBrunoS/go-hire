package dto

type CreateJobInputDTO struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Company     string  `json:"company"`
	Location    *string `json:"location"`
	Remote      bool    `json:"remote"`
	Salary      int64   `json:"salary"`
}

type UpdateJobInputDTO struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Company     string  `json:"company"`
	Location    *string `json:"location"`
	Remote      bool    `json:"remote"`
	Salary      int64   `json:"salary"`
}

type JobOutputDTO struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Company     string  `json:"company"`
	Location    *string `json:"location"`
	Remote      bool    `json:"remote"`
	Salary      int64   `json:"salary"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type JobListOutputDTO struct {
	Jobs []JobOutputDTO `json:"jobs"`
}
