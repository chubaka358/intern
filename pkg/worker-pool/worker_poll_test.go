package worker_pool

import (
	"testing"
)

func BenchmarkWorkerPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		WorkerPool(100)
	}
}