package worker

import (
	"fmt"
	"testing"
	"time"
)

// Benchmark: how fast the pool processes jobs
func BenchmarkWorkerPool(b *testing.B) {
	pool := NewWorkerPool(10, 5000)
	pool.Start()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pool.AddJob(fmt.Sprintf("Job-%d", i))
	}

	pool.Stop()
}

// Stress test: 50k small jobs
func TestStressSmallJobs(t *testing.T) {
	pool := NewWorkerPool(20, 50000)
	pool.Start()

	for i := 0; i < 50000; i++ {
		pool.AddJob(fmt.Sprintf("Small-%d", i))
	}

	pool.Stop()
}

// Stress test: 5k heavy jobs
func TestStressHeavyJobs(t *testing.T) {
	pool := NewWorkerPool(10, 1000)
	pool.Start()

	for i := 0; i < 5000; i++ {
		pool.AddJob(fmt.Sprintf("Heavy-%d", i))
		time.Sleep(time.Millisecond * 2) // simulate heavy load
	}

	pool.Stop()
}

// Stability test: Start/Stop multiple times
func TestRepeatedStartStop(t *testing.T) {
	for cycle := 1; cycle <= 20; cycle++ {
		pool := NewWorkerPool(5, 100)
		pool.Start()

		for i := 0; i < 100; i++ {
			pool.AddJob(fmt.Sprintf("Cycle-%d-Job-%d", cycle, i))
		}

		pool.Stop()
	}
}
