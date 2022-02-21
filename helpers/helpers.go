package helpers

import (
	"log"
	"net/url"
)

// CorrectifyURL adds HTTP:// to URL if not existed
func CorrectifyURL(URL *string) {
	u, err := url.Parse(*URL)
	if err != nil {
		log.Printf("Error when correctifying URL:%s\n", err)
	}

	if u.Scheme == "" {
		addHTTPToUrl(u)
	}
	*URL = u.String()
}

func addHTTPToUrl(URL *url.URL) {
	URL.Scheme = "http"
}
