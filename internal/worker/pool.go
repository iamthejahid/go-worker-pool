package worker

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID string
}

type WorkerPool struct {
	WorkerCount int
	JobQueue    chan Job
	wg          sync.WaitGroup
	Quit        chan bool
}

func NewWorkerPool(workerCount int, queueSize int) *WorkerPool {
	return &WorkerPool{
		WorkerCount: workerCount,
		JobQueue:    make(chan Job, queueSize),
		Quit:        make(chan bool),
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()

	for {
		select {
		case job := <-wp.JobQueue:
			fmt.Printf("Worker <%d> processing |%s|\n", id, job.ID)
			time.Sleep(time.Millisecond * 500) // simulate work
		case <-wp.Quit:
			fmt.Printf("Worker <%d> stopping...\n", id)
			return
		}
	}
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
	close(wp.Quit)
	wp.wg.Wait()
}
