package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPassword(t *testing.T) {
	pass, err := NewPassword("password")
	assert.Nil(t, err)
	assert.NotNil(t, pass)
	assert.NotEmpty(t, pass.Hash)
}

func TestPasswordValidate(t *testing.T) {
	pass, _ := NewPassword("password")
	assert.Nil(t, pass.Validate())

	pass = &Password{}
	assert.ErrorContains(t, pass.Validate(), "hash is required")
}
