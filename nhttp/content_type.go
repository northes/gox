package nhttp

import (
	"strings"
)

type ContentType string

const (
	TextHtml  ContentType = "text/html"
	TextPlain ContentType = "text/plain"
	TextXml   ContentType = "text/xml"

	ImageGif  ContentType = "image/gif"
	ImageJpeg ContentType = "image/jpeg"
	ImagePng  ContentType = "image/png"

	ApplicationXhtmlXml           ContentType = Application + "/xhtml+xml"
	ApplicationXml                ContentType = Application + "/xml"
	ApplicationAtomXml            ContentType = Application + "/atom+xml"
	ApplicationJson               ContentType = Application + "/json"
	ApplicationPDF                ContentType = Application + "/pdf"
	ApplicationMsWord             ContentType = Application + "/msword"
	ApplicationOctetStream        ContentType = Application + "/octet-stream"
	ApplicationXWwwFormUrlencoded ContentType = Application + "/x-www-form-urlencoded"

	MultipartFormData ContentType = "multipart/form-data"
)

func (c ContentType) String() string {
	return string(c)
}

func NewContentType(subType string) string {
	return strings.Join([]string{Application, subType}, "/")
}

func ParseContentType(contentType ContentType) string {
	contentTypeStr := contentType.String()
	left := strings.Index(contentTypeStr, "/")
	if left == -1 {
		return ""
	}
	right := strings.Index(contentTypeStr, ";")
	if right == -1 {
		right = len(contentTypeStr)
	}
	if right < left {
		return ""
	}
	return contentTypeStr[left+1 : right]
}
