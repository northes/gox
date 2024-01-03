package httpx

import (
	"fmt"
	"strings"
	"time"
)

type Option func(client *Client)

func WithHead(key, val string) Option {
	return func(c *Client) {
		c.head.Set(key, val)
	}
}

func WithHeads(head map[string]string) Option {
	return func(c *Client) {
		for k, v := range head {
			c.head.Set(k, v)
		}
	}
}

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.timeout = d
	}
}

func WithBody(body any) Option {
	return func(c *Client) {
		c.body = body
	}
}

func WithPaths(path ...string) Option {
	return func(c *Client) {
		c.url = c.url.JoinPath(path...)
	}
}

func WithRawQuery(rq string) Option {
	return func(c *Client) {
		c.url.RawQuery = rq
	}
}

func WithParams(params map[string]string) Option {
	return func(c *Client) {
		builder := &strings.Builder{}
		for k, v := range params {
			if builder.Len() != 0 {
				builder.WriteString("&")
			}
			builder.WriteString(fmt.Sprintf("%s=%s", k, v))
		}
		c.url.RawQuery = builder.String()
	}
}

func WithDebug(b bool) Option {
	return func(c *Client) {
		c.debug = b
	}
}
