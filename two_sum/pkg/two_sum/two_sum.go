package two_sum


func TwoSum(nums []int, target int) []int {
	numbersMap := make(map[int]int)
	for i := 0; i < len(nums); i++{
		complement := target - nums[i]
		if val, ok := numbersMap[complement]; ok{
			return []int{val, i}
		}
		numbersMap[nums[i]] = i
	}
	panic("Wrong arguments")
}

