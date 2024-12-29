package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Marcelo Soto Campos", "marcelo@teste.com.br", "123321")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "Marcelo Soto Campos", user.Name)
	assert.Equal(t, "marcelo@teste.com.br", user.Email)
}

func TestCheckPassword(t *testing.T) {
	user, err := NewUser("Marcelo Soto Campos", "marcelo@teste.com.br", "123321")
	assert.Nil(t, err)
	assert.True(t, user.CheckPassword("123321"))
	assert.False(t, user.CheckPassword("321123"))
	assert.NotEqual(t, "123321", user.Password)
}
