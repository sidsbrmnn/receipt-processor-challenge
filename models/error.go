package models

import "errors"

var (
	ErrNotFound      = errors.New("no receipt found for that ID, please verify input")
	ErrInvalid       = errors.New("the receipt is invalid, please verify input")
	ErrAlreadyExists = errors.New("a receipt with that ID already exists")
)
