package intersection

import "sort"

// Intersect returns nums1 and nums2 intersection
func Intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i := 0
	j := 0
	inter := make([]int, 0)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			inter = append(inter, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return inter
}
