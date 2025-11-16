package main

import (
	"fmt"
	"go-worker-pool/internal/worker"
)

func main() {
	pool := worker.NewWorkerPool(5, 100) // 5 workers, queue size 100
	pool.Start()

	for i := 1; i <= 20; i++ {
		pool.AddJob(fmt.Sprintf("Job-%d", i))
	}

	pool.Stop()
}
