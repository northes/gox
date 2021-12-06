package nhttp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Name struct {
	Name string `json:"name"`
}

type IP struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data Data   `json:"data"`
}

type Data struct {
	IP       string `json:"ip"`
	Country  string `json:"country"`
	Province string `json:"province"`
	City     string `json:"city"`
	District string `json:"district"`
	ISP      string `json:"isp"`
	Location string `json:"location"`
}

type Login struct {
	Type     int64  `json:"type"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func TestDoGET(t *testing.T) {
	Convey("GET", t, func() {
		var ip IP
		err := GET("https://apihut.net/ip").SetContentType(ApplicationJson).Do().Response().Unmarshal(&ip)
		So(err, ShouldBeNil)
		So(ip.Data, ShouldNotBeNil)
		t.Log(ip.Code)
		t.Logf("%+v", ip.Data)
	})
}

func TestPOST(t *testing.T) {
	Convey("POST", t, func() {
		url := "http://192.168.2.41:8200/api/login/pwd"
		body := &Login{
			Type:     1,
			Account:  "12345678910",
			Password: "123456Abcd",
		}
		re := POST(url).SetBodyWithMarshal(body).Do()
		defer re.Close()
		if re.Error() != nil {
			fmt.Println(re.Error())
		}
		So(re.Error, ShouldNotBeNil)
	})
}

func TestGET(t *testing.T) {
	url := "127.0.0.1"
	get := GET(url)
	Convey("Setter", t, func() {
		So(func() {
			Convey("SetHead", func() {
				get.SetHead(map[string]string{
					"foo": "bar",
				})
			})
			Convey("SetBody", func() {
				get.SetBody([]byte("666"))
			})
			Convey("FlushBody", func() {
				err := get.SetBody([]byte("777")).Error()
				So(err, ShouldNotBeNil)
			})
			Convey("SetBodyWithMarshal", func() {

				So(get.FlushBody().SetBodyWithMarshal(&Name{Name: "northes"}).Error(), ShouldBeNil)
			})
		}, ShouldNotPanic)
	})
	Convey("Getter", t, func() {
		So(func() {
			Convey("GetUrl", func() {
				So(get.GetUrl(), ShouldEqual, url)
			})
			Convey("GetMethod", func() {
				So(get.GetMethod(), ShouldEqual, http.MethodGet)
			})
			Convey("GetHead", func() {
				So(get.Request().GetHead()["foo"], ShouldEqual, "bar")
			})
			Convey("GetRequestBody", func() {
				var body Name
				json.Unmarshal(get.Request().GetBody(), &body)
				So(body.Name, ShouldEqual, "northes")
			})
		}, ShouldNotPanic)
	})

}
