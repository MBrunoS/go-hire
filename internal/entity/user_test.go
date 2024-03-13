package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "john@doe.com", "password")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "john@doe.com", user.Email)
	assert.NotEmpty(t, user.Password)
}

func TestUserValidate(t *testing.T) {
	user, _ := NewUser("John Doe", "john@doe.com", "password")
	assert.Nil(t, user.Validate())

	user = &User{}
	assert.ErrorContains(t, user.Validate(), "id is required")

	user, _ = NewUser("", "john@doe.com", "password")
	assert.ErrorContains(t, user.Validate(), "username is required")

	user, _ = NewUser("John Doe", "", "password")
	assert.ErrorContains(t, user.Validate(), "email is required")

	user, _ = NewUser("John Doe", "john@doe.com", "")
	assert.ErrorContains(t, user.Validate(), "password is required")
}
