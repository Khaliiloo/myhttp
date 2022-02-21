package request

import (
	"fmt"
	"github.com/Khaliiloo/myhttp/response"
	"log"
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	URL := "Go.com"

	result := NewRequest(URL)

	expected := Request{
		URL:    URL,
		client: client,
	}
	if result != expected {
		t.Errorf("NewRequest() test returned an unexpected result: got %v want %v", result, expected)
	}
}

func TestSendRequest(t *testing.T) {
	go func() {
		handler := func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Welcome to Go")
		}
		http.HandleFunc("/", handler)
		log.Println(http.ListenAndServe(":8888", nil))
	}()
	req := NewRequest("http://localhost:8888/Go.com")

	result := req.SendRequest()

	expected := response.Response{
		URL:      "http://localhost:8888/Go.com",
		Response: "Welcome to Go",
		MD5:      "2cd4de2263c8818b8f5bb597bfda422d",
		Err:      nil,
	}

	if result != expected {
		t.Errorf("SendRequest() test returned an unexpected result: got %v want %v", result, expected)
	}
}
