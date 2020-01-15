package worker_pool

import (
	"sync"
)

// WorkerPooler provides workerPool interface
type WorkerPooler interface {
	StartWorkerPool()
}

// workerPool implements worker pool
type workerPool struct {
	employees int
}

// StartWorkerPool launches the employees who do the work, and the foreman who accept their work
func (w *workerPool) StartWorkerPool() {
	wg := &sync.WaitGroup{}

	ch := make(chan string, w.employees)
	// run foreman
	wg.Add(w.employees)
	go NewForeman().accept(w.employees, ch, wg)
	// iterate over employees
	for i := 0; i < w.employees; i++ {
		// run employee #i
		go NewEmployee().work(i, ch)
	}
	wg.Wait()
}

// NewWorkerPool creates and returns new workerPool
func NewWorkerPool(employees int) WorkerPooler {
	return &workerPool{
		employees: employees,
	}
}
