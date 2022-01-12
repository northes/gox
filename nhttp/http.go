package nhttp

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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
	return r.GetResponse().StatusCode == http.StatusOK
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
