package schemas

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Company     string  `json:"company"`
	Location    *string `json:"location"`
	Level       string  `json:"level"`
	Remote      bool    `json:"remote"`
	Salary      int64   `json:"salary"`
}
