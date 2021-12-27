package nhttp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/Northes/ntool/nhttp/nhttputils"
)

type Client struct {
	method  string
	head    http.Header
	body    interface{}
	url     string
	timeout time.Duration
	retry   int64
}

type response struct {
	response *http.Response
}

type Options func(*Client)

func NewClient(url string, opts ...Options) *Client {
	c := Client{
		head:    http.Header{},
		url:     url,
		timeout: time.Second * 5,
		retry:   3,
	}
	for _, opt := range opts {
		opt(&c)
	}
	return &c
}

func WithHead(head map[string]string) Options {
	return func(c *Client) {
		for k, v := range head {
			c.head.Add(k, v)
		}
	}
}

func WithAuthorization(token string) Options {
	return func(c *Client) {
		c.head.Add(nhttputils.Authorization, token)
	}
}

func WithTimeout(t time.Duration) Options {
	return func(c *Client) {
		c.timeout = t
	}
}

func WithRetry(r int64) Options {
	return func(c *Client) {
		c.retry = r
	}
}

func WithBody(body interface{}) Options {
	return func(c *Client) {
		c.body = body
	}
}

func (c *Client) Get() (*response, error) {
	c.method = http.MethodGet
	return c.do()
}

func (c *Client) Post() (*response, error) {
	c.method = http.MethodPost
	return c.do()
}

func (r *response) Unmarshal(body interface{}) error {
	if r.response == nil || r.response.Body == nil {
		return ErrorResponseEmpty
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(r.response.Body)
	bytes, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, body)
}

func (r *response) GetResponse() *http.Response {
	return r.response
}

func (r *response) StatusCodeOK() bool {
	return r.GetResponse().StatusCode == 200
}

func (c *Client) do() (reply *response, err error) {
	var body []byte
	if c.body != nil {
		body, err = json.Marshal(c.body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(c.method, c.url, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	client := http.Client{Timeout: c.timeout}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	reply = new(response)
	reply.response = rsp

	return
}
