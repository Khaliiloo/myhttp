package main

import (
	"context"
	"flag"
	"github.com/Khaliiloo/myhttp/request"
	"github.com/Khaliiloo/myhttp/workpool"
	"io"
	"log"
	"os"
)

// CMD stores parallel flag and received URLs
type CMD struct {
	URLs     []string
	Parallel int
}

// main is the entry of myhttp program
// usage: myhttp [-parallel] url1 url2 ...
func main() {
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("couldn't create log.txt file: %v\n", err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	parallel := flag.Int("parallel", 10, "limits the number of parallel/concurred requests")
	flag.Parse()

	var cmd = CMD{
		URLs:     flag.Args(),
		Parallel: *parallel,
	}

	if len(cmd.URLs) == 0 {
		log.Println("NO URL to GET, usage: myhttp [-parallel] url1 url2 ...")
		return
	}

	// no reason to make number of workers more than URLs
	if cmd.Parallel > len(cmd.URLs) {
		cmd.Parallel = len(cmd.URLs)
	}

	/* just to check time of execution with different parallel values
	defer func(t time.Time) {
		fmt.Println(cmd.Parallel, "parallel took", time.Since(t))
	}(time.Now())
	*/
	wp := workpool.NewWorkerPool(cmd.Parallel)

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()

	go wp.GenerateJobs(collectRequests(cmd.URLs))
	go wp.Run(ctx)

	for {
		select {
		case r, ok := <-wp.Results():
			if !ok {
				continue
			}
			if r.Err != nil {
				log.Println(r.Err)
			} else {
				log.Println(r.URL, r.MD5)
			}
		case <-wp.Done:
			return
		default:
		}
	}

}

// collectRequests takes slice of URL and converts them to slice of request.Request
func collectRequests(URLs []string) []request.Request {
	nOfRequests := len(URLs)
	requests := make([]request.Request, nOfRequests)
	for i := 0; i < nOfRequests; i++ {
		requests[i] = request.NewRequest(URLs[i])
	}
	return requests
}
