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
	req   *request
	rsp   *response
	error error
	mu    sync.Mutex
}

type request struct {
	url     string
	method  string
	head    map[string]string
	body    []byte
	request *http.Request
}

type response struct {
	response *http.Response
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
		req: &request{
			url:     url,
			method:  method,
			head:    make(map[string]string),
			body:    nil,
			request: nil,
		},
		rsp: &response{
			response: nil,
		},
		error: nil,
	}
}

// SetHead 设置头部信息
func (c *Context) SetHead(heads map[string]string) *Context {
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range heads {
		c.req.head[k] = v
	}

	return c
}

// SetBearerToken 设置头部 BearerToken
func (c *Context) SetBearerToken(token string) *Context {
	return c.SetHead(map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	})
}

// SetContentType 设置Content-Type，使用ContentType常量
func (c *Context) SetContentType(contentType ContentType) *Context {
	return c.SetHead(map[string]string{
		"Content-Type": contentType.String(),
	})
}

// SetContentTypeWithStr 设置Content-Type
func (c *Context) SetContentTypeWithStr(contentType string) *Context {
	return c.SetHead(map[string]string{
		"Content-Type": contentType,
	})
}

// SetBody 设置Body
func (c *Context) SetBody(body []byte) *Context {
	if c.req.body != nil {
		c.addError(ErrorRepeatSettingBody)
		return c
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.req.body = body

	return c
}

// SetBodyWithMarshal 反序列化并设置Body
func (c *Context) SetBodyWithMarshal(body interface{}) *Context {
	if c.req.body != nil {
		c.addError(ErrorRepeatSettingBody)
		return c
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	data, err := json.Marshal(body)
	if err != nil {
		c.addError(err)
		return c
	}
	c.req.body = data

	return c
}

// FlushBody 清空Body
func (c *Context) FlushBody() *Context {
	c.req.body = nil
	if errors.Is(c.error, ErrorRepeatSettingBody) {
		c.error = nil
	}
	return c
}

// Do 执行请求
func (c *Context) Do() *Context {
	if len(c.req.url) == 0 || len(c.req.method) == 0 {
		c.addError(ErrorURLOrMethodNotExist)
		return c
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	req, err := http.NewRequest(c.req.method, c.req.url, bytes.NewReader(c.req.body))
	if err != nil {
		c.addError(err)
		return c
	}

	if len(c.req.head) != 0 {
		for k, v := range c.req.head {
			req.Header.Add(k, v)
		}
	}
	c.req.request = req

	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		c.addError(err)
		return c
	}

	c.rsp.response = rsp

	return c
}

// GetUrl 获取发起请求的Url
func (c *Context) GetUrl() string {
	return c.req.url
}

// GetMethod 获取请求方式
func (c *Context) GetMethod() string {
	return c.req.method
}

// Request 获取请求体
func (c *Context) Request() *request {
	return c.req
}

// Response 获取响应体
func (c *Context) Response() *response {
	return c.rsp
}

// Error 获取错误
func (c *Context) Error() error {
	return c.error
}

// Close 关闭请求
func (c *Context) Close() error {
	if c.rsp.response != nil {
		return c.rsp.response.Body.Close()
	}
	return nil
}

// 为Context添加错误
func (c *Context) addError(err error) {
	if c.error == nil {
		c.error = err
	} else if err != nil {
		c.error = fmt.Errorf("%v; %w", c.error, err)
	}
}

/*
Request
*/

// GetUrl 获取发起请求的Url
func (req *request) GetUrl() string {
	return req.url
}

// GetMethod 获取发起请求的方法
func (req *request) GetMethod() string {
	return req.method
}

// GetHead 获取发起请求的头部信息
func (req *request) GetHead() map[string]string {
	return req.head
}

// GetBody 获取发起请求的Body信息
func (req *request) GetBody() []byte {
	return req.body
}

// Get 获取原生Request
func (req *request) Get() *http.Request {
	return req.request
}

/*
Response
*/

// Unmarshal 解析响应结果
func (rsp *response) Unmarshal(model interface{}) error {
	defer rsp.response.Body.Close()
	body, err := ioutil.ReadAll(rsp.response.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}

// GetBody 以[]byte形式获取响应的Body
func (rsp *response) GetBody() ([]byte, error) {
	if rsp.response == nil {
		return nil, ErrorBodyNotExist
	}
	defer rsp.response.Body.Close()
	body, err := ioutil.ReadAll(rsp.response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Get 获取原生Response
func (rsp *response) Get() *http.Response {
	return rsp.response
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
