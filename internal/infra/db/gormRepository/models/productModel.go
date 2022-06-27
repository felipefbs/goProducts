package models

import "github.com/felipefbs/goProducts/pkg/utils"

type Product struct {
	ID    utils.ID `gorm:"primaryKey"`
	Name  string
	Price float64
}
