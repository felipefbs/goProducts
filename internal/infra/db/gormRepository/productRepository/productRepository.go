package productRepository

import (
	"github.com/felipefbs/goProducts/internal/infra/db/gormRepository/models"
	"github.com/felipefbs/goProducts/pkg/domain/entities"
	"github.com/felipefbs/goProducts/pkg/domain/repositories"
	"github.com/felipefbs/goProducts/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewProductRepository() (repositories.ProductRepositoryInterface, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	return &productRepository{
		db: db,
	}, nil
}

func (p *productRepository) Create(product *entities.Product) error {
	productModel := &models.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}

	result := p.db.Create(productModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *productRepository) Update(product *entities.Product) error {
	productModel := &models.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}

	result := p.db.Model(&productModel).Updates(models.Product{
		Name:  product.Name,
		Price: product.Price,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *productRepository) Find(id utils.ID) (*entities.Product, error) {
	var productFound models.Product

	result := p.db.Find(&productFound, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return (*entities.Product)(&productFound), nil
}

func (p *productRepository) FindAll() ([]*entities.Product, utils.Metadata) {
	var productsFound []*models.Product

	result := p.db.Find(&productsFound)
	if result.Error != nil {
		return nil, utils.Metadata{}
	}

	products := make([]*entities.Product, 0)

	for _, p := range productsFound {
		products = append(products, &entities.Product{
			ID:    p.ID,
			Name:  p.Name,
			Price: p.Price,
		})
	}

	return products, utils.Metadata{}
}
