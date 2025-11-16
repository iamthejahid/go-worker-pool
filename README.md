# Go Worker Pool

A lightweight, efficient, and simple worker-pool implementation in Go.\
Ideal for handling concurrent jobs with controlled worker limits, safe
shutdown, and queue management.

------------------------------------------------------------------------

## 🚀 Features

-   Fixed number of workers\
-   Buffered job queue\
-   Graceful start & stop\
-   Thread-safe operations\
-   Back-pressure (when queue is full)\
-   Minimal API, production-ready

------------------------------------------------------------------------

## 📦 Installation

If you're using GitHub, install the module:

``` bash
go get https://github.com/iamthejahid/go-worker-pool
```

Or if you're using locally:

``` bash
go mod init go-worker-pool
```

------------------------------------------------------------------------

## 📁 Project Structure

    go-worker-pool/
    ├── cmd/
    │   └── server/
    │       └── main.go        # Example usage
    ├── internal/
    │   └── worker/
    │       ├── worker.go      # Worker logic
    │       └── pool.go        # WorkerPool implementation
    ├── go.mod
    ├── LICENSE
    └── README.md

------------------------------------------------------------------------

## 🧠 How It Works

A Worker Pool allows you to run tasks concurrently while controlling:

-   How many goroutines run at the same time\
-   How many jobs can be queued\
-   When workers start/stop

This prevents uncontrolled goroutine growth and improves throughput in
CPU-bound or IO-heavy systems.

------------------------------------------------------------------------

## 🧪 Example Usage

``` go
package main

import (
    "fmt"

    "github.com/jahidul-islam-dev/go-worker-pool/internal/worker"
)

func main() {
    pool := worker.NewWorkerPool(5, 100) // 5 workers, queue size 100
    pool.Start()

    for i := 1; i <= 20; i++ {
        pool.AddJob(fmt.Sprintf("Job-%d", i))
    }

    pool.Stop()
}
```

------------------------------------------------------------------------

## ✨ Example Output

    Worker <2> processing |Job-2|
    Worker <4> processing |Job-4|
    Worker <3> processing |Job-3|
    Worker <1> processing |Job-1|
    Worker <5> processing |Job-5|
    Worker <1> stopping...
    Worker <5> processing |Job-6|
    Worker <2> stopping...
    Worker <3> stopping...
    Worker <4> stopping...
    Worker <5> stopping...

------------------------------------------------------------------------

## 🤝 Contributing

Pull Requests welcome.\
Please discuss major changes before submitting.

------------------------------------------------------------------------

## 📄 License

MIT License\
You are free to use this in personal or commercial projects.

------------------------------------------------------------------------

