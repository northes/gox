package jwtx_test

import (
	"testing"
	"time"

	"github.com/northes/gox"
	"github.com/northes/gox/jwtx"
)

type UserClaims struct {
	ID int64 `json:"id"`
}

func TestNewJWT(t *testing.T) {
	ctrl := jwtx.NewJWT[UserClaims]([]byte("secret"), &jwtx.RegisteredClaims{
		Issuer: "test",
		ExpiresAtFunc: func() time.Time {
			return time.Now().Add(time.Second * 1)
		},
		Subject:  "login",
		Audience: []string{"user"},
		NotBeforeFunc: func() time.Time {
			return time.Now()
		},
		IssuedAtFunc: func() time.Time {
			return time.Now()
		},
		IDFunc: func() string {
			return gox.RandGenStr(16, gox.LongLetter)
		},
	})
	token, err := ctrl.Sign(UserClaims{ID: 1})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("token: %s\n", token)

	//time.Sleep(time.Second * 2)

	claims, err := ctrl.Parse(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("claims: %+v", claims)
}
