package hash

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	var res [][]int
	for k := 0; k < len(nums)-3; k++ {
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}
		for i := k + 1; i < len(nums)-2; i++ {
			if i > k+1 && nums[i] == nums[i-1] {
				continue
			}
			left := i + 1
			right := len(nums) - 1
			for right > left {
				num := nums[k] + nums[i] + nums[left] + nums[right]
				if num > target {
					right--
				} else if num < target {
					left++
				} else {
					res = append(res, []int{nums[k], nums[i], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return res
}
