package worker_pool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// employeer provides employee interface
type employeer interface {
	work(i int, ch chan<- string)
}

// foremaner provides foreman interface
type foremaner interface {
	accept(employees int, ch <-chan string, wg *sync.WaitGroup)
}

// employee implements employee
type employee struct {
}

// foreman implements foreman
type foreman struct {
}

// work does the job and sends it to foreman
func (e *employee) work(i int, ch chan<- string) {
	// generate unique seed
	rand.Seed(time.Now().UnixNano())
	// minimum possible work time
	min := 1
	// maximum possible work time
	max := 3
	// workTime is random employee work time between min and max [min; max]
	workTime := rand.Intn(max-min+1) + min
	// doing work...
	time.Sleep(time.Duration(workTime) * time.Second)
	// send result of work to foreman
	ch <- fmt.Sprintf("employee #%v: work done", i)
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

// WorkerPool launches the employees who do the work, and the foreman who accept their work
func WorkerPool(employees int) {
	wg := &sync.WaitGroup{}

	ch := make(chan string, employees)
	// run foreman
	wg.Add(employees)
	go NewForeman().accept(employees, ch, wg)
	// iterate over employees
	for i := 0; i < employees; i++ {
		// run employee #i
		go NewEmployee().work(i, ch)
	}
	wg.Wait()
}

// NewEmployee creates new employee and returns pointer to it
func NewEmployee() *employee {
	return &employee{}
}

// NewForeman creates new foreman and returns pointer to it
func NewForeman() *foreman {
	return &foreman{}
}
