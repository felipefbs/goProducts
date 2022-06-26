package entities

import (
	"github.com/felipefbs/goProducts/pkg/errors"
	"github.com/felipefbs/goProducts/pkg/utils"
)

type OrderItem struct {
	ID       utils.ID
	Product  *Product
	Quantity int
}

func NewOrderItem(product *Product, quantity int) (*OrderItem, error) {
	o := &OrderItem{
		ID:       utils.NewID(),
		Product:  product,
		Quantity: quantity,
	}

	err := o.Validate()
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (o *OrderItem) Validate() error {
	if o.Quantity <= 0 {
		return errors.ErrInvalidQuantity
	}

	err := o.Product.Validate()
	if err != nil {
		return err
	}

	return nil
}
