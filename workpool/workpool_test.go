package workpool

import (
	"context"
	"fmt"
	"github.com/Khaliiloo/myhttp/request"
	"testing"
	"time"
)

func TestWorkerPool(t *testing.T) {
	wp := NewWorkerPool(2)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	requests := []request.Request{
		request.NewRequest("google.com"),
		request.NewRequest("apple.com"),
	}
	go wp.GenerateJobs(requests)
	go wp.Run(ctx)

	for {
		select {
		case r, ok := <-wp.Results():
			if !ok {
				continue
			}
			if r.Err != nil {
				t.Errorf("Error: %v", r.Err)
			} else {
				fmt.Println(r.URL, r.MD5)
			}
		case <-wp.Done:
			return
		default:
		}
	}
}

func TestWorkerPool_TimeOut(t *testing.T) {
	wp := NewWorkerPool(2)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Nanosecond*10)
	defer cancel()

	go wp.Run(ctx)

	for {
		select {
		case r := <-wp.Results():
			if r.Err != nil && r.Err != context.DeadlineExceeded {
				t.Fatalf("expected error: %v; got: %v", context.DeadlineExceeded, r.Err)
			}
		case <-wp.Done:
			return
		default:
		}
	}
}

func TestWorkerPool_Cancel(t *testing.T) {
	wp := NewWorkerPool(2)

	ctx, cancel := context.WithCancel(context.TODO())

	go wp.Run(ctx)
	cancel()

	for {
		select {
		case r := <-wp.Results():
			if r.Err != nil && r.Err != context.Canceled {
				t.Fatalf("expected error: %v; got: %v", context.Canceled, r.Err)
			}
		case <-wp.Done:
			return
		default:
		}
	}
}
