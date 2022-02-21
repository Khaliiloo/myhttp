package request

import (
	"fmt"
	"github.com/Khaliiloo/myhttp/helpers"
	"github.com/Khaliiloo/myhttp/response"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Request struct {
	URL    string
	client *http.Client
	Err    error
}

var client = &http.Client{Timeout: 5 * time.Second}

// NewRequest creates Request
func NewRequest(URL string) Request {
	return Request{
		URL:    URL,
		client: client,
	}
}

// SendRequest sends HTTP Get request and returns response.Response
func (r Request) SendRequest() response.Response {
	stringURL := r.URL
	helpers.CorrectifyURL(&stringURL)
	r.URL = stringURL
	req, err := http.NewRequest(http.MethodGet, r.URL, nil)
	if err != nil {
		log.Fatalf("Error Occurred. %+v", err)
		return response.Response{
			Err: fmt.Errorf("error sending request to API endpoint for %v: couldn't create NewRequest, %w", r.URL, err),
			URL: r.URL,
		}
	}

	resp, err := r.client.Do(req)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error sending request to API endpoint for %v: %w", r.URL, err),
			URL: r.URL,
		}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response.Response{
			Err: fmt.Errorf("error sending request to API endpoint for %v:Couldn't parse response body. %w", r.URL, err),
			URL: r.URL,
		}
	}

	sb := string(body)
	return response.NewResponse(r.URL, sb)
}
