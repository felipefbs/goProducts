package customerRepository_test

import (
	"testing"

	"github.com/felipefbs/goProducts/internal/infra/db/gormRepository/customerRepository"
	"github.com/felipefbs/goProducts/internal/infra/db/gormRepository/models"
	"github.com/felipefbs/goProducts/pkg/domain/entities"
	"github.com/felipefbs/goProducts/pkg/errors"
	"github.com/felipefbs/goProducts/pkg/utils"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

var (
	customerName1   = "some customer name"
	customerAddress = &entities.Address{
		Street:     "Some street",
		Number:     "Some number",
		PostalCode: "580123-250",
		City:       "Some city",
	}
)

func Test_Save(t *testing.T) {
	db, err := connectDB()
	assert.Nil(t, err)

	customerRepo, err := customerRepository.NewCustomerRepository()
	assert.Nil(t, err)

	t.Cleanup(func() {
		db.Exec("DELETE FROM customers")
	})

	t.Run("It should create a customer", func(t *testing.T) {
		customer, err := entities.NewCustomer(customerName1, customerAddress)
		assert.Nil(t, err)

		err = customerRepo.Create(customer)
		assert.Nil(t, err)

		var customerFound models.Customer
		db.Find(&customerFound, customer.ID)

		assert.Equal(t, customerFound.ID, customer.ID)
		assert.Equal(t, customerFound.Name, customer.Name)
		assert.Equal(t, customerFound.Street, customer.Address.Street)
		assert.Equal(t, customerFound.Number, customer.Address.Number)
		assert.Equal(t, customerFound.PostalCode, customer.Address.PostalCode)
		assert.Equal(t, customerFound.City, customer.Address.City)
		assert.Equal(t, customerFound.RewardPoints, customer.RewardPoints)
		assert.Equal(t, customerFound.Active, customer.Active)
	})
}

func Test_Update(t *testing.T) {
	db, err := connectDB()
	assert.Nil(t, err)

	customerRepo, err := customerRepository.NewCustomerRepository()
	assert.Nil(t, err)

	t.Cleanup(func() {
		db.Exec("DELETE FROM customers")
	})

	t.Run("It should update a customer", func(t *testing.T) {
		customer, err := entities.NewCustomer(customerName1, customerAddress)
		assert.Nil(t, err)

		result := db.Create(customerRepository.CustomerModelBuilder(customer))
		assert.Nil(t, result.Error)

		customer.Name = "New Product Name"
		customer.Active = false

		err = customerRepo.Update(customer)
		assert.Nil(t, err)

		var customerFound models.Customer
		result = db.Find(&customerFound, customer.ID)
		assert.Nil(t, result.Error)

		assert.Equal(t, customer.ID, customerFound.ID)
		assert.Equal(t, customer.Name, customerFound.Name)
		assert.Equal(t, customer.Address.Street, customerFound.Street)
		assert.Equal(t, customer.Address.Number, customerFound.Number)
		assert.Equal(t, customer.Address.PostalCode, customerFound.PostalCode)
		assert.Equal(t, customer.Address.City, customerFound.City)
		assert.Equal(t, customer.RewardPoints, customerFound.RewardPoints)
		assert.Equal(t, customer.Active, customerFound.Active)
	})
}

func Test_FindAll(t *testing.T) {
	db, err := connectDB()
	assert.Nil(t, err)

	customerRepo, err := customerRepository.NewCustomerRepository()
	assert.Nil(t, err)

	t.Cleanup(func() {
		db.Exec("DELETE FROM customers")
	})

	t.Run("It should retrieve all customers", func(t *testing.T) {
		customer, err := entities.NewCustomer(customerName1, customerAddress)
		assert.Nil(t, err)

		result := db.Create(customerRepository.CustomerModelBuilder(customer))
		assert.Nil(t, result.Error)

		customer2, err := entities.NewCustomer(customerName1, customerAddress)
		assert.Nil(t, err)

		result = db.Create(customerRepository.CustomerModelBuilder(customer2))
		assert.Nil(t, result.Error)

		customersFound, _ := customerRepo.FindAll()

		assert.Equal(t, 2, len(customersFound))
	})
}

func Test_Find(t *testing.T) {
	db, err := connectDB()
	assert.Nil(t, err)

	customerRepo, err := customerRepository.NewCustomerRepository()
	assert.Nil(t, err)

	t.Cleanup(func() {
		db.Exec("DELETE FROM customers")
	})

	t.Run("It should retrieve a customer", func(t *testing.T) {
		customer, err := entities.NewCustomer(customerName1, customerAddress)
		assert.Nil(t, err)

		result := db.Create(customerRepository.CustomerModelBuilder(customer))
		assert.Nil(t, result.Error)

		customerFound, err := customerRepo.Find(customer.ID)
		assert.Nil(t, err)

		assert.Equal(t, customer, customerFound)
	})

	t.Run("It should not found a customer", func(t *testing.T) {
		customer, err := entities.NewCustomer(customerName1, customerAddress)
		assert.Nil(t, err)

		result := db.Create(customerRepository.CustomerModelBuilder(customer))
		assert.Nil(t, result.Error)

		_, err = customerRepo.Find(utils.NewID())
		assert.EqualError(t, errors.ErrNotFound, err.Error())
	})
}
