package models

import "errors"

var (
	ErrNotFound      = errors.New("No receipt found for that ID. Please verify input.")
	ErrInvalid       = errors.New("The receipt is invalid. Please verify input.")
	ErrAlreadyExists = errors.New("A receipt with that ID already exists. Please verify input.")
)
