package worker

import (
	"fmt"
	"sync"
	"time"
)

// Global flag to disable heavy processing during benchmarks
var benchmarkMode = false

type Job struct {
	ID string
}

type WorkerPool struct {
	WorkerCount int
	JobQueue    chan Job
	wg          sync.WaitGroup
	Quit        chan struct{}
}

func NewWorkerPool(workerCount int, queueSize int) *WorkerPool {
	return &WorkerPool{
		WorkerCount: workerCount,
		JobQueue:    make(chan Job, queueSize),
		Quit:        make(chan struct{}),
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for {
		select {
		case job, ok := <-wp.JobQueue:
			// If channel closed, exit worker
			if !ok {
				fmt.Printf(" -- Worker %d stopping (queue closed)... -- \n", id)
				return
			}

			// Retry logic
			maxAttempts := 3
			for attempt := 1; attempt <= maxAttempts; attempt++ {
				fmt.Printf("[🔥] Worker %d processing %s (attempt %d)\n",
					id, job.ID, attempt)

				err := processJob(job)
				if err == nil {
					break
				}

				// Backoff
				time.Sleep(time.Second * time.Duration(attempt))
			}

			// Simulate heavy work only if NOT in benchmark mode
			if !benchmarkMode {
				time.Sleep(500 * time.Millisecond)
			}

		case <-wp.Quit:
			fmt.Printf(" -- Worker %d stopping... -- \n", id)
			return
		}
	}
}

func processJob(job Job) error {
	// Simulate failure for testing
	if job.ID == "Job-5" {
		return fmt.Errorf("🚧 failed 🚧")
	}
	return nil
}

func (wp *WorkerPool) Start() {
	for i := 1; i <= wp.WorkerCount; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) AddJob(jobID string) {
	wp.JobQueue <- Job{ID: jobID}
}

func (wp *WorkerPool) Stop() {
	fmt.Println("Stopping worker pool...")

	// First: tell workers to stop listening
	close(wp.Quit)

	// Then: close queue so workers exit when finished processing
	close(wp.JobQueue)

	// Wait for all workers to finish
	wp.wg.Wait()

	fmt.Println("All workers stopped.")
}
