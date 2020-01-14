package main

import (
	"github.com/jayhrat/intern/pkg/worker-pool"
)

func main() {
	worker_pool.NewWorkerPool(5).StartWorkerPool()
}
