package main

import (
	"context"
	"flag"
	"io"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/mindtastic/cli/client"
)

var env = flag.String("env", "live", "environment to target. One of [dev, live]")
var nWorkers = flag.Int("workers", 100, "how many workers to spawn")
var runTime = flag.Duration("runtime", 10*time.Minute, "for how long to run the benchmark")

func main() {
	flag.Parse()
	config := client.NewConfiguration()
	apiClient := client.NewAPIClient(config)

	ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
		"environment": *env,
	})

	ctx, cancel := context.WithTimeout(ctx, *runTime)
	defer cancel()

	log.Printf("starting benchmark with %d workers for %s on %s environment", *nWorkers, runTime.String(), *env)

	log.Println("starting registration phase")

	workers := make([]*BenchmarkingClient, *nWorkers)

	maxGoroutines := 50
	guard := make(chan struct{}, maxGoroutines)
	wg := sync.WaitGroup{}

	wg.Add(*nWorkers)
	for i := 0; i < *nWorkers; i++ {
		guard <- struct{}{}
		go func(id int, out io.Writer, client *client.APIClient) {
			bench := NewClient(id, out, client)
			bench.Init(ctx)
			workers[id] = bench
			log.Printf("registered worker %d", id)
			wg.Done()
			<-guard
		}(i, os.Stdout, apiClient)
	}

	wg.Wait()
	log.Println("registration phase done")
	log.Println("starting benchmarking phase")
	time.Sleep(2 * time.Second)

	wg.Add(*nWorkers)
	for i := 0; i < *nWorkers; i++ {
		go func(id int, out io.Writer, client *client.APIClient) {
			bench := workers[id]
			if bench == nil {
				log.Printf("worker %d not found", id)
			}
			n := rand.Intn(2000)
			time.Sleep(time.Duration(n) * time.Millisecond)
			bench.Run(ctx)
			wg.Done()
		}(i, os.Stdout, apiClient)
	}

	wg.Wait()
}
