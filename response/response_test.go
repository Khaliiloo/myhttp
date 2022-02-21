package response

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestNewResponse(t *testing.T) {
	URL := "Go.com"
	resp := "welcome to Go.com"
	MD5 := fmt.Sprintf("%x", md5.Sum([]byte(resp)))
	result := NewResponse(URL, resp)

	expected := Response{
		URL:      URL,
		Response: resp,
		MD5:      MD5,
	}
	if result != expected {
		t.Errorf("NewResponse() test returned an unexpected result: got %v want %v", result, expected)
	}
}
