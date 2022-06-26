package entities

import (
	"github.com/felipefbs/goProducts/pkg/errors"
	"github.com/felipefbs/goProducts/pkg/utils"
)

type Product struct {
	ID    utils.ID
	Name  string
	Price float64
}

func NewProduct(name string, price float64) (*Product, error) {
	p := &Product{
		ID:    utils.NewID(),
		Name:  name,
		Price: price,
	}

	err := p.Validate()
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.ErrInvalidName
	}

	if p.Price <= 0 {
		return errors.ErrInvalidPrice
	}

	return nil
}
