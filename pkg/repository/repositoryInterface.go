package repository

import "github.com/felipefbs/goProducts/pkg/utils"

type RepositoryInterface[T any] interface {
	Create(entity *T) error
	Update(entity *T) error
	Find(id utils.ID) (*T, error)
	FindAll() ([]*T, utils.Metadata)
}
