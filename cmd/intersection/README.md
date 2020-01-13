# Intersection of Two Arrays
Given two arrays, returns their intersection.

**Example 1:**
```
Input: nums1 = [1, 2, 2, 1], nums2 = [2, 2]
Output: [2, 2]
```

**Example 2:**
```
Input: nums1 = [4, 9, 5], nums2 = [9, 4, 9, 8, 4]
Output: [4, 9]
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/chubaka358/intern/pkg/intersection"
)

func main() {
	fmt.Println(intersection.NewIntersecter().Intersect([]int{3, 2, 1, 3}, []int{3, 1, 2}))
}
```

### Link
https://leetcode.com/problems/intersection-of-two-arrays-ii/
