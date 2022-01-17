package httpx

import "errors"

var (
	ErrorResponseEmpty    = errors.New("响应为空")
	ErrorNotABBearerToken = errors.New("错误的 Bearer Token 格式")
)
