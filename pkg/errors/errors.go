package errors

import "errors"

var (
	ErrInvalidName    = errors.New("Invalid Name")
	ErrInvalidAddress = errors.New("Invalid Address")
	ErrNotActive      = errors.New("Customer not active")
)
