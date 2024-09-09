package httpx

import (
	"time"

	"github.com/northes/gox/httpx/httpxutils"
)

type Option func(client *Client)

func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.timeout = d
	}
}

func WithDebug(b bool) Option {
	return func(c *Client) {
		c.debug = b
	}
}

func WithLogger(logger Logger) Option {
	return func(c *Client) {
		c.logger = logger
	}
}

func WithJsonEncoder(encoder httpxutils.JSONMarshal) Option {
	return func(c *Client) {
		c.jsonEncoder = encoder
	}
}

func WithJsonDecoder(decoder httpxutils.JSONUnmarshal) Option {
	return func(c *Client) {
		c.jsonDecoder = decoder
	}
}
