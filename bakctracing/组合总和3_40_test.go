package bakctracing

import (
	"fmt"
	"sort"
	"testing"
)

func combinationSum2(candidates []int, target int) [][]int {
	path, ans := make([]int, 0), make([][]int, 0)
	sort.Ints(candidates)
	var dfs func(start, target int)
	dfs = func(start, target int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target { // 剪枝，提前返回
				break
			}
			if i != start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			dfs(i+1, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return ans
}

func TestCombinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))

}
