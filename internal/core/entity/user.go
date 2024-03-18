package entity

import (
	"errors"

	"github.com/mbrunos/go-hire/pkg/id"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       id.ID  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u := &User{
		ID:       id.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}

	if err := u.Validate(); err != nil {
		return nil, err
	}

	return u, nil
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
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.ComparePassword("") {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
