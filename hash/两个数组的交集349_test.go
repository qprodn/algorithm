package hash

func intersection(nums1 []int, nums2 []int) []int {
	nums := [10]int{}
	ans := make([]int, 0)
	for _, val := range nums1 {
		if nums[val] == 0 {
			nums[val]++
		}
	}
	for _, val := range nums2 {
		if nums[val] > 0 {
			nums[val]--
			ans = append(ans, val)
		}
	}
	return ans
}
