package repositories

import (
	"github.com/felipefbs/goProducts/pkg/domain/entities"
	"github.com/felipefbs/goProducts/pkg/repository"
)

type CustomerRepositoryInterface interface {
	repository.RepositoryInterface[entities.Customer]
}
