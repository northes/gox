package gox

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

func FakeAnimalName() string {
	name := fmt.Sprintf("%s %s", gofakeit.AdjectiveDescriptive(), gofakeit.Animal())
	return strings.ReplaceAll(name, " ", "_")
}
