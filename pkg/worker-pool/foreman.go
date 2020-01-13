package worker_pool

import (
	"fmt"
	"sync"
)

// foremaner provides foreman interface
type foremaner interface {
	accept(employees int, ch <-chan string, wg *sync.WaitGroup)
}

// foreman implements foreman
type foreman struct {
}

// accept accepts employee's work
func (f *foreman) accept(employees int, ch <-chan string, wg *sync.WaitGroup) {
	// iterate over employee's work
	for i := 0; i < employees; i++ {
		// accept employee's work
		fmt.Printf("foreman: job accepted: %q\n", <-ch)
		wg.Done()
	}
}

// NewForeman creates new foreman and returns pointer to it
func NewForeman() *foreman {
	return &foreman{}
}
