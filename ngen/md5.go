package ngen

import (
	"crypto/md5"
	"encoding/hex"
)

func NewMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckMd5(str string, md5 string) bool {
	return NewMd5(str) == md5
}
