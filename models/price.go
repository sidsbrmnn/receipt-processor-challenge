package models

import (
	"strconv"
)

type Price string

// return the price as a float64
func (p Price) Float64() float64 {
	f, _ := strconv.ParseFloat(string(p), 64)
	return f
}
