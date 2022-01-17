package httpx

import (
	"time"
)

type Options func(*Client)

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

func WithDebug() Options {
	return func(c *Client) {
		c.debug = true
	}
}
