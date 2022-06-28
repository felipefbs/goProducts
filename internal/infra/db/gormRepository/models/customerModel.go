package models

import "github.com/felipefbs/goProducts/pkg/utils"

type Customer struct {
	ID           utils.ID `gorm:"primaryKey"`
	Name         string
	Street       string
	Number       string
	PostalCode   string
	City         string
	RewardPoints float64
	Active       bool
}
