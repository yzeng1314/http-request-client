package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	reqInterval  int
	reqUrl       string
	numOfWorkers int
)

func init() {
	// The request interval in second
	flag.IntVar(&reqInterval, "interval", 5, "The interval between each request")
	// The number of workers that is gonna send the request
	flag.IntVar(&numOfWorkers, "workerNum", 1, "The number of workers that is gonna send the request")
	// The request target endpoint
	flag.StringVar(&reqUrl, "url", "http://google.com", "The URL endpoint that request is gonna hit (Don't forget http[s]://))")
}

func httpGet(url *string, worker int) {
	resp, err := http.Get(*url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	/**
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	} **/

	log.Printf("Woker: %d, Method: GET, Endpoint: %s, ResponseCode: %d", worker, *url, resp.StatusCode)
}

func main() {
	// Parse the command-line flags
	flag.Parse()

	quit := make(chan struct{})

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	defer stop()

	for i := 1; i <= numOfWorkers; i++ {
		go func(worker int) {
			ticker := time.NewTicker(time.Duration(reqInterval) * time.Second)
			for {
				select {
				case <-ticker.C:
					httpGet(&reqUrl, worker)
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}(i)
	}

	<-ctx.Done()
	log.Println("Shutting down the program...")
	close(quit)

}
