package schemas

import (
	"time"
)

type Job struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Company     string     `json:"company"`
	Location    *string    `json:"location"`
	Level       string     `json:"level"`
	Remote      bool       `json:"remote"`
	Salary      int64      `json:"salary"`
}
