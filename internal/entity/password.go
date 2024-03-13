package entity

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	Hash string `json:"hash"`
}

func NewPassword(password string) (*Password, error) {
	if password == "" {
		return nil, errors.New("password is required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Password{
		Hash: string(hash),
	}, nil
}

func (p *Password) Validate() error {
	if p.Hash == "" {
		return errors.New("hash is required")
	}

	return nil
}

// Compares the password with the given password and returns
// true if they match, false if they don't
func (p *Password) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.Hash), []byte(password))
	return err == nil
}
