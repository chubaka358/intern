package main

import (
	"github.com/chubaka358/intern/pkg/worker-pool"
)

func main() {
	worker_pool.NewWorkerPool(5).StartWorkerPool()
}
