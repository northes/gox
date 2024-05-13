package jwtx

import "errors"

var (
	ErrParseFailed     = errors.New("parse failed")
	ErrAuthStringEmpty = errors.New("auth str empty")
	ErrFormatInvalid   = errors.New("format invalid")
	ErrNotBearerToken  = errors.New("not bearer token")
)
