package core

import "errors"

var (
	ErrEmailAlreadyInUse = errors.New("email already in use")
	ErrEmailNotFound     = errors.New("email not found")
)
