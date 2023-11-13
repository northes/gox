package gox

import "crypto/rand"

const (
	NumberLetters    = "0123456789"
	LowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	CapitalLetters   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LongLetter       = NumberLetters + LowercaseLetters + CapitalLetters
)

func RandGenStr(n int64, letter string) string {
	l := []byte(letter)
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	arc := uint8(0)

	if _, err := rand.Read(b[:]); err != nil {
		return ""
	}
	for i, x := range b {
		arc = x & byte(len(l)-1)
		b[i] = l[arc]
	}
	return string(b)
}
