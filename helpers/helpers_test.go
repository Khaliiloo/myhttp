package helpers

import (
	"net/url"
	"testing"
)

func TestCorrectifyURL(t *testing.T) {
	Go := "Go.com"
	CorrectifyURL(&Go)
	result := Go
	expected := "http://Go.com"
	if result != expected {
		t.Errorf("CorrectifyURL() test returned an unexpected result: got %v want %v", result, expected)
	}
}

func TestAddHTTP(t *testing.T) {
	u, _ := url.Parse("Go.com")
	addHTTPToUrl(u)
	result := u.String()
	expected := "http://Go.com"
	if result != expected {
		t.Errorf("AddHTTP() test returned an unexpected result: got %v want %v", result, expected)
	}
}
