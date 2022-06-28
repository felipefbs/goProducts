package customerRepository

import (
	"fmt"

	"github.com/felipefbs/goProducts/internal/infra/db/gormRepository/models"
	"github.com/felipefbs/goProducts/pkg/domain/entities"
	"github.com/felipefbs/goProducts/pkg/domain/repositories"
	"github.com/felipefbs/goProducts/pkg/errors"
	"github.com/felipefbs/goProducts/pkg/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func initDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Customer{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewCustomerRepository() (repositories.CustomerRepositoryInterface, error) {
	db, err := initDB()
	if err != nil {
		return nil, err
	}

	return &customerRepository{
		db: db,
	}, nil
}

func (p *customerRepository) Create(customer *entities.Customer) error {
	customerModel := CustomerModelBuilder(customer)
	fmt.Println(customer)
	result := p.db.Create(customerModel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *customerRepository) Update(customer *entities.Customer) error {
	customerModel := CustomerModelBuilder(customer)

	result := p.db.Model(&customerModel).Updates(CustomerModelBuilder(customer))
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *customerRepository) Find(id utils.ID) (*entities.Customer, error) {
	var customerFound models.Customer

	result := p.db.First(&customerFound, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, result.Error
	}

	return CustomerEntityBuilder(&customerFound), nil
}

func (p *customerRepository) FindAll() ([]*entities.Customer, utils.Metadata) {
	var customerFound []*models.Customer

	result := p.db.Find(&customerFound)
	if result.Error != nil {
		return nil, utils.Metadata{}
	}

	customers := make([]*entities.Customer, 0)

	for _, c := range customerFound {
		customers = append(customers, CustomerEntityBuilder(c))
	}

	return customers, utils.Metadata{}
}

func CustomerEntityBuilder(c *models.Customer) *entities.Customer {
	customer := &entities.Customer{
		ID:   c.ID,
		Name: c.Name,
		Address: &entities.Address{
			Street:     c.Street,
			Number:     c.Number,
			PostalCode: c.PostalCode,
			City:       c.City,
		},
		Active:       c.Active,
		RewardPoints: c.RewardPoints,
	}

	return customer
}

func CustomerModelBuilder(c *entities.Customer) *models.Customer {
	customer := &models.Customer{
		ID:           c.ID,
		Name:         c.Name,
		Street:       c.Address.Street,
		Number:       c.Address.Number,
		PostalCode:   c.Address.PostalCode,
		City:         c.Address.City,
		RewardPoints: c.RewardPoints,
		Active:       c.Active,
	}

	return customer
}
