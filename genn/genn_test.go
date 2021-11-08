package genn

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

	md5 := Md5(LowercaseLetters)
	fmt.Println(md5)
	fmt.Println(CheckMd5(LowercaseLetters, md5))
}

func TestSnow(t *testing.T) {
	id, err := Snow(12)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id.String())
}
