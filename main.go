package main

import (
	"context"
	"flag"
	"github.com/mindtastic/cli/client"
	"io"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var env = flag.String("env", "dev", "environment to target. One of [dev, live]")
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

	wg := sync.WaitGroup{}
	for i := 0; i < *nWorkers; i++ {
		wg.Add(1)
		go func(id int, out io.Writer, client *client.APIClient) {
			bench := NewClient(id, out, client)
			bench.Init(ctx)
			n := rand.Intn(10000)
			time.Sleep(time.Duration(n) * time.Millisecond)
			bench.Run(ctx)
			wg.Done()
		}(i, os.Stdout, apiClient)
	}

	wg.Wait()
}
