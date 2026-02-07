package bakctracing

import (
	"fmt"
	"testing"
)

func combinationSum3(k int, n int) [][]int {
	var path []int
	var ans [][]int

	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			sum := 0
			for _, v := range path {
				sum += v
			}
			if sum == n {
				tmp := make([]int, k)
				copy(tmp, path)
				ans = append(ans, tmp)
			}
			return
		}
		if len(path) > k {
			return
		}

		for i := start; i <= 9; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)
	return ans
}

func TestCombinationSum3(t *testing.T) {
	sum3 := combinationSum3(3, 7)
	fmt.Println(sum3)
}
