# Set

Set data structure implementation

## Usage

```go
package main

import "github.com/chubaka358/intern/pkg/set"

func main() {
    mySet := set.NewSet() // []
    mySet.Add("15", "13", "11", "11", "11") // ["11", "13", "15"] 
    mySet.Delete("13") // ["11", "15"]
    mySet.Size() // 2
    mySet.Contains("15") // true
    
    newSet := set.NewSet() // []
    newSet.Add("15", "17") // ["15", "17"]
    intersectionSet := mySet.Intersection(&newSet) // ["15"]
    unionSet := mySet.Union(&newSet) // ["11", "15", "17"]
}
```
