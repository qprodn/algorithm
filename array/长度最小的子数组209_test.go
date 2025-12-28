package array

import (
	"fmt"
	"testing"
)

func minSubArrayLen(target int, nums []int) int {
	i := 0
	sum := 0
	subL := 0
	result := len(nums) + 1
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			subL = j - i + 1
			if subL < result {
				result = subL
			}
			sum -= nums[i]
			i++
		}
	}
	if result == len(nums)+1 {
		return 0
	}
	return result
}

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{1, 4, 4}
	arrayLen := minSubArrayLen(4, nums)
	fmt.Println(arrayLen)
}
