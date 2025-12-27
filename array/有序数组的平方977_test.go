package array

import (
	"fmt"
	"testing"
)

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	left, right := 0, len(nums)-1
	pos := right
	for left <= right {
		l1 := nums[left] * nums[left]
		r1 := nums[right] * nums[right]
		if l1 < r1 {
			ans[pos] = r1
			pos--
			right--
		} else {
			ans[pos] = l1
			pos--
			left++
		}
	}
	return ans
}

func TestSortedSquares(t *testing.T) {
	nums := []int{-7, -3, 2, 3, 11}
	squares := sortedSquares(nums)
	for _, square := range squares {
		fmt.Println(square)
	}
}
