package nfile

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileInfo_IsImage(t *testing.T) {
	Convey("真实图片", t, func() {
		isImage := Path("./1.jpg").Info().IsImage()
		So(isImage, ShouldEqual, true)
		fmt.Println(isImage)
	})

	Convey("假图片", t, func() {
		isImage := Path("./1.png").Info().IsImage()
		So(isImage, ShouldEqual, false)
		fmt.Println(isImage)
	})
}

func TestCompress(t *testing.T) {
	Convey("打包", t, func() {
		So(Compress([]string{"./1.jpg", "./filen.go"}, "./compress.zip"), ShouldNotBeNil)
	})
}
