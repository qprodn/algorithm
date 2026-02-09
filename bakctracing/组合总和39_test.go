package bakctracing

import (
	"fmt"
	"sort"
	"testing"
)

func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var ans [][]int
	path, ans = make([]int, 0), make([][]int, 0)
	sort.Ints(candidates)
	var dfs func(strat int, tmpTarget int)
	dfs = func(start int, tmpTarget int) {
		if tmpTarget == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > tmpTarget {
				break
			}
			path = append(path, candidates[i])
			dfs(i, tmpTarget-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return ans
}

func TestCombinationSum(t *testing.T) {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
