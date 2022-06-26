package entities_test

import (
	"testing"

	"github.com/felipefbs/goProducts/pkg/domain/entities"
	"github.com/felipefbs/goProducts/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var (
	customerName    = "john"
	customerAddress = &entities.Address{
		Street:     "9 Street",
		Number:     "10",
		PostalCode: "123123",
		City:       "York",
	}
)

func Test_Customer_Aggregate(t *testing.T) {
	t.Run("It should not create an customer with an empty name", func(t *testing.T) {
		_, err := entities.NewCustomer("", customerAddress)
		assert.EqualError(t, err, errors.ErrInvalidName.Error())
	})

	t.Run("It should not change customer name for a empty name", func(t *testing.T) {
		customer, err := entities.NewCustomer(customerName, customerAddress)
		assert.Nil(t, err)

		err = customer.ChangeName("")
		assert.EqualError(t, err, errors.ErrInvalidName.Error())
	})

	t.Run("It should change customer name", func(t *testing.T) {
		customer, err := entities.NewCustomer(customerName, customerAddress)
		assert.Nil(t, err)

		err = customer.ChangeName("John Doe")
		assert.Nil(t, err)
	})

	t.Run("It should not activate customer with an invalid address", func(t *testing.T) {
		customer, err := entities.NewCustomer(customerName, &entities.Address{})
		assert.Nil(t, err)

		err = customer.Activate()
		assert.EqualError(t, err, errors.ErrInvalidAddress.Error())
	})

	t.Run("It should not create an address", func(t *testing.T) {
		_, err := entities.NewAddress("", "", "", "")
		assert.EqualError(t, err, errors.ErrInvalidAddress.Error())
	})

	t.Run("It should activate customer", func(t *testing.T) {
		address, err := entities.NewAddress("9 Street", "10", "123123", "York")
		assert.Nil(t, err)

		customer, err := entities.NewCustomer(customerName, address)
		assert.Nil(t, err)

		err = customer.Activate()
		assert.Nil(t, err)
	})

	t.Run("It should deactivate customer", func(t *testing.T) {
		address, err := entities.NewAddress("9 Street", "10", "123123", "York")
		assert.Nil(t, err)

		customer, err := entities.NewCustomer(customerName, address)
		assert.Nil(t, err)

		err = customer.Activate()
		assert.Nil(t, err)

		customer.Deactivate()
	})

	t.Run("It should not add reward points to a not active customer", func(t *testing.T) {
		address, err := entities.NewAddress("9 Street", "10", "123123", "York")
		assert.Nil(t, err)

		customer, err := entities.NewCustomer(customerName, address)
		assert.Nil(t, err)

		err = customer.AddPoints(10.0)
		assert.EqualError(t, err, errors.ErrNotActive.Error())
	})

	t.Run("It should add reward points", func(t *testing.T) {
		address, err := entities.NewAddress("9 Street", "10", "123123", "York")
		assert.Nil(t, err)

		customer, err := entities.NewCustomer(customerName, address)
		assert.Nil(t, err)

		err = customer.Activate()
		assert.Nil(t, err)

		err = customer.AddPoints(10.0)
		assert.Nil(t, err)
		assert.Equal(t, customer.RewardPoints, 10.0)

		err = customer.AddPoints(20.0)
		assert.Nil(t, err)
		assert.Equal(t, customer.RewardPoints, 30.0)
	})
}
