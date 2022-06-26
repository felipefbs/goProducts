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

		_, err = entities.NewOrder(utils.NewID(), []entities.OrderItem{
			{Quantity: -1},
		})
		assert.EqualError(t, err, errors.ErrInvalidQuantity.Error())
	})

	t.Run("It should not create an order item with invalid fields", func(t *testing.T) {
		_, err := entities.NewOrderItem(&entities.Product{
			Name:  "",
			Price: 1,
		}, 1)
		assert.EqualError(t, err, errors.ErrInvalidName.Error())

		_, err = entities.NewOrderItem(&entities.Product{
			Name:  "Name",
			Price: 0,
		}, 1)
		assert.EqualError(t, err, errors.ErrInvalidPrice.Error())
	})

	t.Run("It should create an order", func(t *testing.T) {
		_, err := entities.NewOrderItem(orderItemProduct1, 1)
		assert.Nil(t, err)
	})

	t.Run("It should not create a product with invalid fields", func(t *testing.T) {
		_, err := entities.NewProduct("", orderItemProduct2.Price)
		assert.EqualError(t, err, errors.ErrInvalidName.Error())

		_, err = entities.NewProduct(orderItemProduct1.Name, 0)
		assert.EqualError(t, err, errors.ErrInvalidPrice.Error())
	})

	t.Run("It should create a product", func(t *testing.T) {
		_, err := entities.NewProduct(orderItemProduct2.Name, orderItemProduct2.Price)
		assert.Nil(t, err)
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
