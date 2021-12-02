package nhttp

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type Demo struct {
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

func TestGET(t *testing.T) {
	Convey("GET", t, func() {
		url := "https://apihut.net/ip"
		resp := new(Demo)
		fmt.Println(GET(url).Do().Unmarshal(resp).Error)
		fmt.Printf("%+v", resp)
		So(resp.Code, ShouldNotBeNil)
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
		if re.Error != nil {
			fmt.Println(re.Error)
		}
		So(re.Error, ShouldNotBeNil)
	})
}
