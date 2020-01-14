#worker-pool
There are workers who start work in parallel (work time is a random number from 1 to 3 seconds),
then they complete the work and submit their results for verification to the foreman who accepts them

## Usage

```go
package main

import (
	"github.com/jayhrat/intern/pkg/worker-pool"
)

func main() {
	worker_pool.NewWorkerPool(5).StartWorkerPool()
}
```

Output:
```
foreman: job accepted: "employee #3: work done"
foreman: job accepted: "employee #0: work done"
foreman: job accepted: "employee #2: work done"
foreman: job accepted: "employee #4: work done"
foreman: job accepted: "employee #1: work done"
```
