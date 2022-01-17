package genx

import (
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

func Hash(h hash.Hash, str string) string {
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum([]byte("ntool")))
}

func Sha512(str string) string {
	s := sha512.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum([]byte("ntool")))
}
