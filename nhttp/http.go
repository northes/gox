package nhttp

import (
	"net/http"
	"time"
)

type Client struct {
	head    http.Header
	timeout time.Duration
	retry   int64
}

type Options func(*Client)

func NewClient(opts ...Options) *Client {
	c := Client{
		head:    http.Header{},
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
		c.head.Add(Authorization, token)
	}
}

func WithTimeout(t time.Duration) Options {
	return func(c *Client) {
		c.timeout = t
	}
}
