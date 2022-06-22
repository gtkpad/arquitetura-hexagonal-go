package dto

import (
	"errors"

	"github.com/gtkpad/arquitetura-hexagonal-go/application"
)

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	Status string `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *application.Product) (*application.Product, error) {
	if (p.ID != "") {
		product.ID = p.ID
	}

	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status
	valid, err := product.IsValid()
	if err != nil {
		return &application.Product{}, err
	}
	if !valid {
		return &application.Product{}, errors.New("Invalid product")
	}

	return product, nil
}