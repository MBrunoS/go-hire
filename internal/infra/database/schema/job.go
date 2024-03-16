package schema

import "time"

type Job struct {
	ID          string     `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Company     string     `json:"company"`
	Location    *string    `json:"location"`
	Remote      bool       `json:"remote"`
	Salary      int64      `json:"salary"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
