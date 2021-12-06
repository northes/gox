package nhttp

import "errors"

var (
	ErrorRepeatSettingBody   = errors.New("Repeat Setting Body ")
	ErrorURLOrMethodNotExist = errors.New("URl Or Method is not exist")
	ErrorBodyNotExist        = errors.New("body is nil")
)
