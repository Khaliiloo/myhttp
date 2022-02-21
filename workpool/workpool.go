package workpool

import (
	"context"
	"fmt"
	"github.com/Khaliiloo/myhttp/request"
	"github.com/Khaliiloo/myhttp/response"
	"sync"
)

// worker receives request.Request from jobs channel and calls request.Request.SendRequest() to send response.Response in results channel
func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan request.Request, results chan<- response.Response) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				return
			}
			results <- job.SendRequest()
		case <-ctx.Done():
			fmt.Printf("cancelled worker. Error detail: %v\n", ctx.Err())
			results <- response.Response{
				Err: ctx.Err(),
			}
			return
		}
	}
}

type WorkerPool struct {
	workersCount int
	jobs         chan request.Request
	results      chan response.Response
	Done         chan struct{}
}

// NewWorkerPool creates new workers pool
func NewWorkerPool(workerCount int) WorkerPool {
	return WorkerPool{
		workersCount: workerCount,
		jobs:         make(chan request.Request, workerCount),
		results:      make(chan response.Response, workerCount),
		Done:         make(chan struct{}),
	}
}

// Run runs workers and waits until they finish work
func (wp WorkerPool) Run(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < wp.workersCount; i++ {
		wg.Add(1)
		go worker(ctx, &wg, wp.jobs, wp.results)
	}

	wg.Wait()
	close(wp.Done)
	close(wp.results)
}

func (wp WorkerPool) Results() <-chan response.Response {
	return wp.results
}

// GenerateJobs generates jobs from []request.Request into jobs channel
func (wp WorkerPool) GenerateJobs(jobsBulk []request.Request) {
	for i, _ := range jobsBulk {
		wp.jobs <- jobsBulk[i]
	}
	close(wp.jobs)
}
