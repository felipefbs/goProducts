package entities_test

import (
	"testing"

	"github.com/felipefbs/goProducts/pkg/domain/entities"
	"github.com/felipefbs/goProducts/pkg/errors"
	"github.com/felipefbs/goProducts/pkg/utils"
	"github.com/stretchr/testify/assert"
)

var (
	orderCustomerID   = utils.NewID()
	orderItemProduct1 = &entities.Product{
		Name:  "product name",
		Price: 3.14,
	}
	orderItemProduct2 = &entities.Product{
		Name:  "product name",
		Price: 4.13,
	}
	orderItems = []entities.OrderItem{
		{
			Product:  orderItemProduct1,
			Quantity: 2,
		},
		{
			Product:  orderItemProduct2,
			Quantity: 3,
		},
	}
)

func Test_Order_Aggregate(t *testing.T) {
	t.Run("It should not create an order with invalid fields", func(t *testing.T) {
		_, err := entities.NewOrder("", []entities.OrderItem{})
		assert.EqualError(t, err, errors.ErrInvalidID.Error())
	})

	t.Run("It should create an order", func(t *testing.T) {
		_, err := entities.NewOrder(orderCustomerID, orderItems)
		assert.Nil(t, err)
	})

	t.Run("It should return the total order amount", func(t *testing.T) {
		order, err := entities.NewOrder(orderCustomerID, orderItems)
		assert.Nil(t, err)

		actualTotal := order.Items[0].Product.Price*float64(order.Items[0].Quantity) + order.Items[1].Product.Price*float64(order.Items[1].Quantity)

		orderTotal := order.Total()
		assert.Equal(t, actualTotal, orderTotal)
	})
}
