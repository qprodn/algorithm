package stack_list

import "sort"

func topKFrequent(nums []int, k int) []int {
	ans := make([]int, 0, len(nums))
	mapHep := map[int]int{}
	for _, v := range nums {
		mapHep[v]++
	}
	for key, _ := range mapHep {
		ans = append(ans, key)
	}
	sort.Slice(ans, func(a, b int) bool {
		return mapHep[ans[a]] > mapHep[ans[b]]
	})
	return ans[:k]
}
