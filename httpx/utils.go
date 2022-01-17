package httpx

import (
	"fmt"
	"strings"
)

func NewBearerToken(token string) string {
	return fmt.Sprintf("%s %s", Bearer, token)
}

func ParseBearerToken(authorHead string) (token string, err error) {
	parts := strings.SplitN(token, " ", 2)
	if !(len(parts) == 2 && parts[0] == Bearer) {
		return "", ErrorNotABBearerToken
	}
	return parts[1], nil
}
