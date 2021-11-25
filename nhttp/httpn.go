package nhttp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

func (c *Context) SetHead(heads map[string]string) *Context {
	c.Head = heads
	return c
}

func (c *Context) SetBody(body []byte) *Context {
	c.Body = body
	return c
}

func (c *Context) MarshalBody(body interface{}) *Context {
	data, err := json.Marshal(body)
	if err != nil {
		c.addError(err)
		return c
	}
	c.Body = data
	return c
}

func (c *Context) Do() *Context {
	if len(c.Url) == 0 || len(c.Method) == 0 {
		c.addError(errors.New("URl Or Method is not exist"))
		return c
	}
	req, err := http.NewRequest(c.Method, c.Url, bytes.NewReader(c.Body))
	if err != nil {
		c.addError(err)
		return c
	}

	if len(c.Head) != 0 {
		for k, v := range c.Head {
			req.Header.Add(k, v)
		}
	}
	c.Request = req

	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		c.addError(err)
		return c
	}

	c.Response = re

	return c
}

func (c *Context) Unmarshal(model interface{}) *Context {
	body, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		c.addError(err)
		return c
	}
	if err = json.Unmarshal(body, model); err != nil {
		c.addError(err)
		return c
	}
	return c
}

func (c *Context) ByteBody() ([]byte, error) {
	body, err := ioutil.ReadAll(c.Response.Body)
	if err != nil {
		c.addError(err)
		return nil, err
	}
	return body, nil
}

func (c *Context) Close() error {
	return c.Response.Body.Close()
}

func (c *Context) addError(err error) {
	if c.Error == nil {
		c.Error = err
	} else if err != nil {
		c.Error = fmt.Errorf("%v; %w", c.Error, err)
	}
}

func Download(url string, name string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
