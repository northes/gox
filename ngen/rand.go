package ngen

import (
	"math/rand"
	"time"
)

const (
	NumberLetters    = "0123456789"
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	CapitalLetters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LongLetter       = NumberLetters + LowercaseLetters + CapitalLetters
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandKey 根据给定的字符串，生成n长度的随机字符串
func RandKey(n int, str string) string {
	letter := []byte(str)
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	arc := uint8(0)
	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}
	for i, x := range b {
		arc = x & byte(len(letter)-1)
		b[i] = letter[arc]
	}
	return string(b)
}
