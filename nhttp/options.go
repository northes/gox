package nhttp

import (
	"time"
)

type Options func(*Client)

func WithHead(head map[string]string) Options {
	return func(c *Client) {
		if len(head) > 0 {
			for k, v := range head {
				if c.head.Get(k) == "" {
					c.head.Add(k, v)
				}
			}
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
