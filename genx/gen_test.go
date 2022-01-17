package genx

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRandKey(t *testing.T) {
	Convey("", t, func() {
		key := RandKey(6, LongLetter)
		So(key, ShouldNotBeBlank)
	})
	Convey("", t, func() {
		key := RandKey(0, NumberLetters)
		So(key, ShouldBeBlank)
	})
}

func TestMd5(t *testing.T) {

	md5 := NewMd5(LowercaseLetters)
	fmt.Println(md5)
	fmt.Println(CheckMd5(LowercaseLetters, md5))
}

func TestSnow(t *testing.T) {
	id, err := NewSnow(12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id.String())
}

func TestUUID(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(NewUUID().String())
	}
}
