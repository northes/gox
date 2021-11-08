package httpn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Response *http.Response
	Error    error
}

func GET(url string) *Response {
	rsp := &Response{}
	httpRsp, err := http.Get(url)
	if err != nil {
		rsp.addError(err)
		return rsp
	}
	return &Response{Response: httpRsp, Error: nil}
}

func POST(url string, data []byte) *Response {
	rsp := &Response{}
	contentType := "application/json"
	body := bytes.NewReader(data)
	httpRsp, err := http.Post(url, contentType, body)
	if err != nil {
		rsp.addError(err)
		return rsp
	}
	return &Response{Response: httpRsp, Error: nil}
}

func POSTWithMarshal(url string, data interface{}) *Response {
	rsp := &Response{}
	body, err := json.Marshal(data)
	if err != nil {
		rsp.addError(err)
		return rsp
	}
	return POST(url, body)
}

func (r *Response) Unmarshal(values interface{}) *Response {
	body, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		r.addError(err)
		return r
	}
	if err = json.Unmarshal(body, values); err != nil {
		r.addError(err)
		return r
	}
	return r
}

func (r *Response) Body() string {
	body, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		r.addError(err)
		return ""
	}
	return string(body)
}

// addError add err
func (r *Response) addError(err error) {
	if r.Error == nil {
		r.Error = err
	} else if err != nil {
		r.Error = fmt.Errorf("%v; %w", r.Error, err)
	}
}
