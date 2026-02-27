package greedy

import (
	"fmt"
	"testing"
)

func jump(nums []int) int {
	// 根据题目规则，初始位置为nums[0]
	lastDistance := 0 // 上一次覆盖范围
	curDistance := 0  // 当前覆盖范围（可达最大范围）
	minStep := 0      // 记录最少跳跃次数

	for i := 0; i < len(nums); i++ {
		if i == lastDistance+1 { // 在上一次可达范围+1的位置，记录步骤
			minStep++                  // 跳跃次数+1
			lastDistance = curDistance // 记录时才可以更新
		}
		curDistance = max(nums[i]+i, curDistance) // 更新当前可达的最大范围
	}
	return minStep
}
func jump2(nums []int) int {
	ans, next, cur := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]+i > next {
			next = nums[i] + i
		}
		if i == cur {
			if cur != len(nums)-1 {
				ans++
				cur = next
				if cur >= len(nums)-1 {
					break
				}
			}
		}
	}
	return ans
}
func TestJump(t *testing.T) {
	fmt.Println(jump([]int{1, 2}))
}
