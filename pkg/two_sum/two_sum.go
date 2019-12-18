package two_sum

// Function to return indices of the two numbers such that they add up to a specific target
func TwoSum(nums []int, target int) []int {
	// define map containing iterated integers from nums
	numbersMap := make(map[int]int)

	// Loop that iterates over nums slice
	for i := 0; i < len(nums); i++ {
		// calculate complement number
		complement := target - nums[i]
		// check contains numbersMap complement or not. Return if contains
		if val, ok := numbersMap[complement]; ok {
			return []int{val, i}
		}
		// add iterated value in numbersMap
		numbersMap[nums[i]] = i
	}
	// incorrect input
	panic("Wrong arguments")
}
