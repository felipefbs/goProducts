package entities

import "github.com/felipefbs/goProducts/pkg/errors"

type Address struct {
	Street     string
	Number     string
	PostalCode string
	City       string
}

func NewAddress(street, number, postalCode, city string) (*Address, error) {
	a := &Address{
		Street:     street,
		Number:     number,
		PostalCode: postalCode,
		City:       city,
	}

	err := a.Validate()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Address) Validate() error {
	if a.City == "" || a.Number == "" || a.PostalCode == "" || a.Street == "" {
		return errors.ErrInvalidAddress
	}

	return nil
}
