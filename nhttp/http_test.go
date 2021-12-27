package nhttp

import (
	"fmt"
	"testing"
)

type ip struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Ip       string `json:"ip"`
		Country  string `json:"country"`
		Province string `json:"province"`
		City     string `json:"city"`
		District string `json:"district"`
		Isp      string `json:"isp"`
		Location string `json:"location"`
	} `json:"data"`
}

func TestGet(t *testing.T) {
	rsp, err := NewClient("https://apihut.net/ip").Get()
	if err != nil {
		t.Error(err)
	}
	i := &ip{}
	if err = rsp.Unmarshal(i); err != nil {
		t.Error(err)
	}
	t.Logf(i.Data.Ip)
}

func TestPost(t *testing.T) {
	type body struct {
		Hello string `json:"hello"`
	}
	rsp, err := NewClient("http://127.0.0.1:8888", WithBody(&body{Hello: "world"})).Post()
	if err != nil {
		t.Error(err)
		return
	}
	if !rsp.StatusCodeOK() {
		t.Errorf("%+v", rsp.GetResponse())
		return
	}
	rb := &body{}
	if err = rsp.Unmarshal(rb); err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%+v\n", rb)
}
