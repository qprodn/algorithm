package greedy

import (
	"fmt"
	"testing"
)

func maxSubArray(nums []int) int {
	sum := nums[0]
	preSum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]+preSum < 0 {
			preSum = 0
			if nums[i] > sum {
				sum = nums[i]
			}
		} else {
			preSum = preSum + nums[i]
			if preSum > sum {
				sum = preSum
			}
		}
	}
	return sum
}

func TestMaxSubArray(t *testing.T) {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

}
