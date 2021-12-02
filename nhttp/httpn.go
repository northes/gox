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
	"sync"
)

type Context struct {
	url      string
	method   string
	head     map[string]string
	body     []byte
	request  *http.Request
	response *http.Response
	error    error
	mu       *sync.Mutex
}

// GET 以GET方法初始化
func GET(url string) *Context {
	return NewContext(http.MethodGet, url)
}

// POST 以POST方法初始化
func POST(url string) *Context {
	return NewContext(http.MethodPost, url)
}

// PUT 以PUT方法初始化
func PUT(url string) *Context {
	return NewContext(http.MethodPut, url)
}

// PATCH 以PATCH方法初始化
func PATCH(url string) *Context {
	return NewContext(http.MethodPatch, url)
}

// DELETE 以DELETE方法初始化
func DELETE(url string) *Context {
	return NewContext(http.MethodDelete, url)
}

// NewContext 初始化
func NewContext(method string, url string) *Context {
	return &Context{
		url:      url,
		method:   method,
		head:     nil,
		body:     nil,
		request:  nil,
		response: nil,
		error:    nil,
	}
}

// SetHead 设置头部信息
func (c *Context) SetHead(heads map[string]string) *Context {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range heads {
		c.head[k] = v
	}

	return c
}

// SetBody 设置Body
func (c *Context) SetBody(body []byte) *Context {
	if c.body != nil {
		c.addError(errors.New("重复设置Body"))
		return c
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.body = body

	return c
}

// SetBodyWithMarshal 反序列化并设置Body
func (c *Context) SetBodyWithMarshal(body interface{}) *Context {
	if c.body != nil {
		c.addError(errors.New("重复设置body"))
		return c
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	data, err := json.Marshal(body)
	if err != nil {
		c.addError(err)
		return c
	}
	c.body = data

	return c
}

// Do 执行请求
func (c *Context) Do() *Context {
	if len(c.url) == 0 || len(c.method) == 0 {
		c.addError(errors.New("URl Or Method is not exist"))
		return c
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	req, err := http.NewRequest(c.method, c.url, bytes.NewReader(c.body))
	if err != nil {
		c.addError(err)
		return c
	}

	if len(c.head) != 0 {
		for k, v := range c.head {
			req.Header.Add(k, v)
		}
	}
	c.request = req

	client := &http.Client{}
	re, err := client.Do(req)
	if err != nil {
		c.addError(err)
		return c
	}

	c.response = re

	return c
}

// Unmarshal 解析返回结果
func (c *Context) Unmarshal(model interface{}) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	body, err := ioutil.ReadAll(c.response.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}

// GetByteBody 获取[]byte形式的Body
func (c *Context) GetByteBody() ([]byte, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	body, err := ioutil.ReadAll(c.response.Body)
	if err != nil {
		c.addError(err)
		return nil, err
	}
	return body, nil
}

// Close 关闭请求
func (c *Context) Close() error {
	return c.response.Body.Close()
}

func (c *Context) addError(err error) {
	if c.error == nil {
		c.error = err
	} else if err != nil {
		c.error = fmt.Errorf("%v; %w", c.error, err)
	}
}

// Download 下载文件
// name 为完整路径，包括欲保存文件名(包括后缀)
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
