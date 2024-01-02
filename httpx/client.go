package httpx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	method   string
	head     http.Header
	body     any
	url      *url.URL
	timeout  time.Duration
	response http.Response
	logger   Logger
	debug    bool
}

type Response struct {
	response *http.Response
	logger   Logger
}

func NewClient(rawURL string, opts ...Option) (*Client, error) {
	cli := new(Client)
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	cli.url = u
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(cli)
		}
	}
	return cli, nil
}

func (c *Client) SetBody(body any) *Client {
	c.body = body
	return c
}

func (c *Client) Get() (*Response, error) {
	c.method = http.MethodGet
	return c.do()
}

func (c *Client) Post() (*Response, error) {
	c.method = http.MethodPost
	return c.do()
}

func (c *Client) do() (*Response, error) {
	var (
		req  *http.Request
		err  error
		body *bytes.Reader
	)

	if c.body != nil {
		switch c.body.(type) {
		case string:
			b := c.body.(string)
			body = bytes.NewReader([]byte(b))
		case []byte:
			b := c.body.([]byte)
			body = bytes.NewReader(b)
		default:
			b, err := json.Marshal(c.body)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("body unmarshal(%v): %v", c.body, err))
			}
			body = bytes.NewReader(b)
		}
	}

	if body == nil {
		req, err = http.NewRequest(c.method, c.url.String(), nil)
	} else {
		req, err = http.NewRequest(c.method, c.url.String(), body)
	}

	if err != nil {
		return nil, errors.New(fmt.Sprintf("new request: %v", err))
	}

	client := http.Client{
		Timeout: c.timeout,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("client Do(): %v", err))
	}

	return &Response{
		response: resp,
		logger:   c.logger,
	}, nil
}

func (r *Response) String() string {
	b := r.response.Body
	defer func() {
		_ = r.response.Body.Close()
	}()
	body, err := io.ReadAll(b)
	if err != nil {
		r.logger.Fatal(err.Error())
		return ""
	}
	return fmt.Sprintf("Status: %s\nBody: %s", r.response.Status, string(body))
}

func (r *Response) Unmarshal(body any) error {
	if !r.IsStatusOK() {
		return errors.New(fmt.Sprintf("status is %d not %d", r.response.StatusCode, http.StatusOK))
	}
	if body == nil || r.RawResponse() == nil {
		return errors.New(fmt.Sprintf("response is nil or input body id nil"))
	}
	b := r.response.Body
	defer func() {
		_ = r.response.Body.Close()
	}()
	bb, err := io.ReadAll(b)
	if err != nil {
		return err
	}
	return json.Unmarshal(bb, body)
}

func (r *Response) RawResponse() *http.Response {
	return r.response
}

func (r *Response) IsStatusOK() bool {
	return r.response.StatusCode == http.StatusOK
}
