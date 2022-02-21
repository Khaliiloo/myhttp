package response

import (
	"crypto/md5"
	"fmt"
)

type Response struct {
	URL      string
	Response string
	MD5      string
	Err      error
}

// NewResponse creates Response
func NewResponse(url, response string) Response {
	return Response{
		URL:      url,
		Response: response,
		MD5:      fmt.Sprintf("%x", md5.Sum([]byte(response))),
	}
}
