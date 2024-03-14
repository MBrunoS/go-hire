package model

type User struct {
	ID       string ` json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"index:,unique"`
	Password string `json:"-"`
}
