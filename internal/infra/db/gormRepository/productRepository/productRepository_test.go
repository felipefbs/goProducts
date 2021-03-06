package productRepository_test

import (
	"testing"

	"github.com/felipefbs/goProducts/internal/infra/db/gormRepository/models"
	"github.com/felipefbs/goProducts/internal/infra/db/gormRepository/productRepository"
	"github.com/felipefbs/goProducts/pkg/domain/entities"
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
	productName1  = "some product name"
	productPrice1 = 3.14
)

func Test_Save(t *testing.T) {
	db, err := connectDB()
	assert.Nil(t, err)

	productRepo, err := productRepository.NewProductRepository()
	assert.Nil(t, err)

	t.Cleanup(func() {
		db.Exec("DELETE FROM products")
	})

	t.Run("It should create a product", func(t *testing.T) {
		product, err := entities.NewProduct(productName1, productPrice1)
		assert.Nil(t, err)

		err = productRepo.Create(product)
		assert.Nil(t, err)

		var productFound models.Product
		db.Find(&productFound, product.ID)

		assert.Equal(t, productFound.ID, product.ID)
		assert.Equal(t, productFound.Name, product.Name)
		assert.Equal(t, productFound.Price, product.Price)
	})
}

func Test_Update(t *testing.T) {
	db, err := connectDB()
	assert.Nil(t, err)

	productRepo, err := productRepository.NewProductRepository()
	assert.Nil(t, err)

	t.Cleanup(func() {
		db.Exec("DELETE FROM products")
	})

	t.Run("It should update a product", func(t *testing.T) {
		product, err := entities.NewProduct(productName1, productPrice1)
		assert.Nil(t, err)

		result := db.Create(product)
		assert.Nil(t, result.Error)

		product.Name = "New Product Name"
		product.Price = 4.5

		err = productRepo.Update(product)
		assert.Nil(t, err)

		var productFound models.Product
		result = db.Find(&productFound, product.ID)
		assert.Nil(t, result.Error)

		assert.Equal(t, product.ID, productFound.ID)
		assert.Equal(t, product.Name, productFound.Name)
		assert.Equal(t, product.Price, productFound.Price)
	})
}

func Test_FindAll(t *testing.T) {
	db, err := connectDB()
	assert.Nil(t, err)

	productRepo, err := productRepository.NewProductRepository()
	assert.Nil(t, err)

	t.Cleanup(func() {
		db.Exec("DELETE FROM products")
	})

	t.Run("It should retrieve all products", func(t *testing.T) {
		product, err := entities.NewProduct(productName1, productPrice1)
		assert.Nil(t, err)

		result := db.Create(product)
		assert.Nil(t, result.Error)

		product2, err := entities.NewProduct(productName1, productPrice1)
		assert.Nil(t, err)

		result = db.Create(product2)
		assert.Nil(t, result.Error)

		productsFound, _ := productRepo.FindAll()

		assert.Equal(t, 2, len(productsFound))
	})
}

func Test_Find(t *testing.T) {
	db, err := connectDB()
	assert.Nil(t, err)

	productRepo, err := productRepository.NewProductRepository()
	assert.Nil(t, err)

	t.Cleanup(func() {
		db.Exec("DELETE FROM products")
	})

	t.Run("It should retrieve a product", func(t *testing.T) {
		product, err := entities.NewProduct(productName1, productPrice1)
		assert.Nil(t, err)

		result := db.Create(product)
		assert.Nil(t, result.Error)

		productFound, err := productRepo.Find(product.ID)
		assert.Nil(t, err)

		assert.Equal(t, product.ID, productFound.ID)
		assert.Equal(t, product.Name, productFound.Name)
		assert.Equal(t, product.Price, productFound.Price)
	})
}
