package jwtx

import (
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/northes/gox/httpx/httpxutils"
)

type JWT[T any] struct {
	secret           []byte
	data             T
	registeredClaims *RegisteredClaims
}

type Claims[T any] struct {
	jwt.RegisteredClaims
	Data T `json:"data"`
}

type RegisteredClaims struct {
	Issuer        string
	ExpiresAtFunc func() time.Time
	Subject       string
	Audience      []string
	NotBeforeFunc func() time.Time
	IssuedAtFunc  func() time.Time
	IDFunc        func() string
}

func NewJWT[T any](secret []byte, registeredClaims *RegisteredClaims) *JWT[T] {
	return &JWT[T]{
		secret:           secret,
		registeredClaims: registeredClaims,
	}
}

func (j *JWT[T]) Sign(data T) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims[T]{
		RegisteredClaims: j.registeredClaims.ToJWTRegisteredClaims(),
		Data:             data,
	})
	return token.SignedString(j.secret)
}

func (j *JWT[T]) ParseHeader(str string) (*Claims[T], error) {
	if len(str) == 0 {
		return nil, ErrAuthStringEmpty
	}
	token := strings.SplitN(str, " ", 2)
	if len(token) != 2 {
		return nil, ErrFormatInvalid
	}
	if token[0] != httpxutils.BearerTokenPrefix {
		return nil, ErrNotBearerToken
	}
	return j.Parse(token[1])
}

func (j *JWT[T]) Parse(token string) (*Claims[T], error) {
	t, err := jwt.ParseWithClaims(token, &Claims[T]{}, func(token *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := t.Claims.(*Claims[T])
	if ok && t.Valid {
		return claims, nil
	}
	return nil, ErrParseFailed
}

func (c *Claims[T]) GetData() T {
	return c.Data
}

func (r *RegisteredClaims) ToJWTRegisteredClaims() jwt.RegisteredClaims {
	newAudience := make([]string, len(r.Audience))
	copy(newAudience, r.Audience)
	return jwt.RegisteredClaims{
		Issuer:    r.Issuer,
		Subject:   r.Subject,
		Audience:  newAudience,
		ExpiresAt: jwt.NewNumericDate(r.ExpiresAtFunc()),
		NotBefore: jwt.NewNumericDate(r.NotBeforeFunc()),
		IssuedAt:  jwt.NewNumericDate(r.IssuedAtFunc()),
		ID:        r.IDFunc(),
	}
}
