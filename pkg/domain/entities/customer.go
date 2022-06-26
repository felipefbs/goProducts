package entities

import (
	"fmt"
	"math/rand"

	"github.com/felipefbs/goProducts/pkg/errors"
)

type Customer struct {
	ID           string
	Name         string
	Address      *Address
	Active       bool
	RewardPoints float64
}

func NewCustomer(name string, address *Address) (*Customer, error) {
	c := &Customer{
		ID:           fmt.Sprint(rand.Int()),
		Name:         name,
		Address:      address,
		Active:       false,
		RewardPoints: 0,
	}

	err := c.Validate()
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return errors.ErrInvalidName
	}

	return nil
}

func (c *Customer) ChangeName(name string) error {
	c.Name = name

	if err := c.Validate(); err != nil {
		return err
	}

	return nil
}

func (c *Customer) Activate() error {
	err := c.Address.Validate()
	if err != nil {
		return err
	}

	c.Active = true

	return nil
}

func (c *Customer) Deactivate() {
	c.Active = false
}

func (c *Customer) AddPoints(value float64) error {
	if !c.Active {
		return errors.ErrNotActive
	}

	c.RewardPoints += value

	return nil
}
