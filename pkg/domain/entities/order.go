package entities

import (
	"github.com/felipefbs/goProducts/pkg/errors"
	"github.com/felipefbs/goProducts/pkg/utils"
)

type Order struct {
	ID         utils.ID
	CustomerID utils.ID
	Items      []OrderItem
}

func NewOrder(customerID utils.ID, items []OrderItem) (*Order, error) {
	o := &Order{
		ID:         utils.NewID(),
		CustomerID: customerID,
		Items:      items,
	}

	err := o.Validate()
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (o *Order) Validate() error {
	if o.CustomerID == "" {
		return errors.ErrInvalidID
	}

	for _, item := range o.Items {
		if err := item.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (o *Order) Total() float64 {
	var total float64

	for _, item := range o.Items {
		total += item.Product.Price * float64(item.Quantity)
	}

	return total
}
