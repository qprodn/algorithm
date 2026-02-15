package greedy

import (
	"fmt"
	"testing"
)

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var dfs func(start int) bool
	dfs = func(start int) bool {
		if start == 0 {
			return true
		}
		flag := false
		for i := start - 1; i >= 0; i-- {
			if nums[i]+i >= start {
				flag = dfs(i)
				if flag {
					return flag
				}
			}
		}
		return flag
	}
	return dfs(len(nums) - 1)
}

func canJump2(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	maxCap := nums[0]
	for i := 0; i <= maxCap && i < len(nums); i++ {
		if maxCap >= len(nums)-1 {
			return true
		}
		if nums[i]+i > maxCap {
			maxCap = nums[i] + i
		}
	}
	return false
}

func TestCanJump(t *testing.T) {
	t.Run("[2,3,1,1,4]", func(t *testing.T) {
		fmt.Println(canJump2([]int{1, 2, 3}))
	})
}
