package array

import (
	"fmt"
	"testing"
)

func removeElement(nums []int, val int) int {
	right := len(nums) - 1
	left := 0
	count := 0
	for left <= right {
		if nums[left] == val {
			tmp := nums[left]
			nums[left] = nums[right]
			nums[right] = tmp
			right--
			count++
		} else if nums[left] != val {
			left++
		}
		for _, num := range nums {
			fmt.Print(num, ",")
		}
		fmt.Println()
	}
	return len(nums) - count
}

func removeElement2(nums []int, val int) int {
	stk := 0
	for _, num := range nums {
		if num != val {
			nums[stk] = num
			stk++
		}
	}
	return stk
}

func TestA(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	element := removeElement(nums, 2)
	fmt.Println(element)
}
