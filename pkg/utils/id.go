package utils

import (
	"fmt"
	"math/rand"
)

type ID string

func NewID() ID {
	return ID(fmt.Sprint(rand.Int()))
}
