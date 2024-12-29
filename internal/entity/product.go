package entity

import (
	"errors"
	"time"

	"github.com/sk8sta13/API-Service/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return errors.New("ID is required")
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return errors.New("ID is invalid")
	}

	if p.Name == "" {
		return errors.New("Name is required")
	}

	if p.Price == 0 {
		return errors.New("Price is required")
	}

	if p.Price < 0 {
		return errors.New("Price is invalid")
	}

	return nil
}
