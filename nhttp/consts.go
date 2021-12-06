package nhttp

type ContentType string

const (
	TextHtml  ContentType = "text/html"
	TextPlain ContentType = "text/plain"
	TextXml   ContentType = "text/xml"

	ImageGif  ContentType = "image/gif"
	ImageJpeg ContentType = "image/jpeg"
	ImagePng  ContentType = "image/png"

	ApplicationXhtmlXml           ContentType = "application/xhtml+xml"
	ApplicationXml                ContentType = "application/xml"
	ApplicationAtomXml            ContentType = "application/atom+xml"
	ApplicationJson               ContentType = "application/json"
	ApplicationPDF                ContentType = "application/pdf"
	ApplicationMsWord             ContentType = "application/msword"
	ApplicationOctetStream        ContentType = "application/octet-stream"
	ApplicationXWwwFormUrlencoded ContentType = "application/x-www-form-urlencoded"

	MultipartFormData ContentType = "multipart/form-data"
)

func (c ContentType) String() string {
	return string(c)
}
