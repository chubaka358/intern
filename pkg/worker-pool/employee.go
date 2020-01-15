package worker_pool

import (
	"fmt"
	"math/rand"
	"time"
)

// Employeer provides employee interface
type Employeer interface {
	work(i int, ch chan<- string)
}

// employee implements employee
type employee struct {
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

// NewEmployee creates new employee and returns pointer to it
func NewEmployee() Employeer {
	return &employee{}
}
