package ngen

import (
	"crypto/sha1"
	"testing"
)

func TestSha512(t *testing.T) {
	t.Logf(Sha512("123"))
}

func TestHash(t *testing.T) {
	t.Log(Hash(sha1.New(), "123"))
}
