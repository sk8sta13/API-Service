package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, e := NewProduct("Carteira", 86.0)
	assert.Nil(t, e)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, "Carteira", p.Name)
	assert.Equal(t, 86.0, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, e := NewProduct("", 86.0)
	assert.Nil(t, p)
	assert.Equal(t, errors.New("Name is required"), e)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, e := NewProduct("Carteira", 0)
	assert.Nil(t, p)
	assert.Equal(t, errors.New("Price is required"), e)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, e := NewProduct("Carteira", -20.0)
	assert.Nil(t, p)
	assert.Equal(t, errors.New("Price is invalid"), e)
}

func TestProductValidate(t *testing.T) {
	p, e := NewProduct("Carteira", 86.0)
	assert.Nil(t, e)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
