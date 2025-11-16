package main

import (
	"go-worker-pool/internal/api"
	"go-worker-pool/internal/worker"
)

func main() {
	pool := worker.NewWorkerPool(5, 100)
	pool.Start()

	api.RegisterRoutes(pool)
}
