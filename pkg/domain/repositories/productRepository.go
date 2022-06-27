package repositories

import (
	"github.com/felipefbs/goProducts/pkg/domain/entities"
	"github.com/felipefbs/goProducts/pkg/repository"
)

type ProductRepositoryInterface interface {
	repository.RepositoryInterface[entities.Product]
}
