package entity

import (
	"errors"

	"github.com/mbrunos/go-hire/pkg/id"
)

type User struct {
	ID       id.ID    `json:"id"`
	Name     string   `json:"username"`
	Email    string   `json:"email"`
	Password Password `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	pass, err := NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       id.NewID(),
		Name:     name,
		Email:    email,
		Password: *pass,
	}, nil
}

// Validates the user struct
func (u *User) Validate() error {
	if id.IsNil(u.ID) || u.ID.String() == "" {
		return errors.New("id is required")
	}

	if _, err := id.StringToID(u.ID.String()); err != nil {
		return errors.New("id is not valid")
	}

	if u.Name == "" {
		return errors.New("username is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	return nil
}
