package errors

import "errors"

var (
	ErrInvalidName     = errors.New("Invalid Name")
	ErrInvalidAddress  = errors.New("Invalid Address")
	ErrNotActive       = errors.New("Customer not active")
	ErrInvalidID       = errors.New("Invalid ID")
	ErrInvalidQuantity = errors.New("Quantity should be greater then 0")
	ErrInvalidPrice    = errors.New("Price should be greater then 0")
	ErrNotFound        = errors.New("Not found")
)
