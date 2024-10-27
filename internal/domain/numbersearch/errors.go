package numbersearch

import "errors"

var (
	ErrNumberNotFound = errors.New("number not found")
	ErrEmptyList      = errors.New("empty list")
)
