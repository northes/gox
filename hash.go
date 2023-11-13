package gox

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
)

func HashByBytes(b []byte) (string, error) {
	hash := sha256.New()
	rd := bytes.NewReader(b)
	if _, err := io.Copy(hash, rd); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func HashByReader(rd io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, rd); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
