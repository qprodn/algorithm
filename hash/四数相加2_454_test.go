package hash

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	ab := make(map[int]int, n*n) //预分配内存性能高一倍
	for _, vala := range nums1 {
		for _, valb := range nums2 {
			ab[valb+vala]++
		}
	}
	count := 0
	for _, valc := range nums3 {
		for _, vald := range nums4 {
			if val, exist := ab[0-valc-vald]; exist {
				count += val
			}
		}
	}
	return count
}

func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n := len(nums1)
	m := make(map[int]int, n*n)
	var t int
	for i := 0; i < n; i++ {
		t = nums1[i]
		for j := 0; j < n; j++ {
			m[t+nums2[j]]++
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		t = nums3[i]
		for j := 0; j < n; j++ {
			if v, ok := m[-(t + nums4[j])]; ok {
				res += v
			}
		}
	}
	return res
}
