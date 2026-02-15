package greedy

import (
	"fmt"
	"testing"
)

func maxProfit(prices []int) int {
	sum := 0
	if len(prices) < 2 {
		return 0
	}
	start := 0
	prediff := prices[1] - prices[0]
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff < 0 && prediff >= 0 {
			sum += prices[i-1] - prices[start]
		} else if diff > 0 && prediff < 0 {
			start = i - 1
		}
		prediff = diff
	}
	return sum
}

func TestMaxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
}
