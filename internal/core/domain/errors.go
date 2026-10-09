package domain

import "errors"

var (
	ErrNotFound     = errors.New("todo not found")
	ErrInvalidInput = errors.New("invalid input")
)
