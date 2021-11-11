package httpn

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Context struct {
	Url      string
	Method   string
	Head     map[string]string
	Body     []byte
	Request  *http.Request
	Response *http.Response
	Error    error
}

func GET(url string) *Context {
	return NewContext(http.MethodGet, url)
}

func POST(url string) *Context {
	return NewContext(http.MethodPost, url)
}

func PUT(url string) *Context {
	return NewContext(http.MethodPut, url)
}

func PATCH(url string) *Context {
	return NewContext(http.MethodPatch, url)
}

func DELETE(url string) *Context {
	return NewContext(http.MethodDelete, url)
}

func NewContext(method string, url string) *Context {
	return &Context{
		Url:      url,
		Method:   method,
		Head:     nil,
		Body:     nil,
		Request:  nil,
		Response: nil,
		Error:    nil,
	}
}

func (h *Context) SetHead(heads map[string]string) *Context {
	h.Head = heads
	return h
}

func (h *Context) SetBody(body []byte) *Context {
	h.Body = body
	return h
}

func (h *Context) MarshalBody(body interface{}) *Context {
	data, err := json.Marshal(body)
	if err != nil {
		h.addError(err)
		return h
	}
	h.Body = data
	return h
}

func (h *Context) Do() *Context {
	if len(h.Url) == 0 || len(h.Method) == 0 {
		h.addError(errors.New("URl Or Method is not exist"))
		return h
	}
	req, err := http.NewRequest(h.Method, h.Url, bytes.NewReader(h.Body))
	if err != nil {
		h.addError(err)
		return h
	}

	if len(h.Head) != 0 {
		for k, v := range h.Head {
			req.Header.Add(k, v)
		}
	}
	h.Request = req

	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		h.addError(err)
		return h
	}

	h.Response = re

	return h
}

func (h *Context) Unmarshal(model interface{}) *Context {
	body, err := ioutil.ReadAll(h.Response.Body)
	if err != nil {
		h.addError(err)
		return h
	}
	if err = json.Unmarshal(body, model); err != nil {
		h.addError(err)
		return h
	}
	return h
}

func (h *Context) Close() error {
	return h.Response.Body.Close()
}

func (h *Context) addError(err error) {
	if h.Error == nil {
		h.Error = err
	} else if err != nil {
		h.Error = fmt.Errorf("%v; %w", h.Error, err)
	}
}
