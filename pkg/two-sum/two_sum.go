package two_sum

// TwoSum returns the indices of elements whose sum is equal to the target
// Causes a panic if two of these elements cannot be found.
func TwoSum(nums []int, target int) []int {
	// Stores iterated items
	numbersMap := make(map[int]int)

	// Loop that iterates over nums slice
	for i := 0; i < len(nums); i++ {
		// Сalculate complement number
		complement := target - nums[i]
		// Сheck contains numbersMap complement or not. Return if contains
		if val, ok := numbersMap[complement]; ok {
			return []int{val, i}
		}
		// Add iterated value in numbersMap
		numbersMap[nums[i]] = i
	}
	// incorrect input
	panic("Wrong arguments")
}
