package worker_pool

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// WorkerPool launches the employees who do the work, and the foreman who accept their work
func WorkerPool(employees int) {
	wg := sync.WaitGroup{}

	ch := make(chan string, employees)
	// run foreman
	wg.Add(employees)
	go func() {
		// iterate over employee's work
		for i := 0; i < employees; i++ {
			// accept employee's work
			fmt.Printf("foreman: job accepted: %q\n", <-ch)
			wg.Done()
		}
	}()
	// iterate over employees
	for i:= 0; i < employees; i++ {
		// employee number
		i := i
		// run employee #i
		go func() {
			// generate unique seed
			rand.Seed(time.Now().UnixNano())
			// minimum possible work time
			min := 1
			// maximum possible work time
			max := 3
			// workTime is random employee work time between min and max [min; max]
			workTime := rand.Intn(max - min + 1) + min
			// doing work...
			time.Sleep(time.Duration(workTime) * time.Second)
			// send result of work to foreman
			ch <- fmt.Sprintf("employee #%v: work done", i)
		}()
	}
	wg.Wait()
}
