package genn

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/bwmarrin/snowflake"
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

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func CheckMd5(str string, md5 string) bool {
	return Md5(str) == md5
}

func Snow(n int64) (snowflake.ID, error) {
	node, err := snowflake.NewNode(n)
	if err != nil {
		return 0, err
	}
	id := node.Generate()
	return id, nil
}
