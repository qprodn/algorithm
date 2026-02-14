package bakctracing

import (
	"fmt"
	"testing"
)

func partition(s string) [][]string {
	path, ans := make([]string, 0), make([][]string, 0)
	ishw := func(arr []string) bool {
		if len(path) == 0 {
			return false
		}
		l, r := 0, len(arr)-1
		for l < r {
			if arr[l] != arr[r] {
				return false
			}
			l++
			r--
		}
		return true
	}
	var dfs func(start int, lens int)
	dfs = func(start int, lens int) {
		if lens >= len(s) {
			return
		}
		if ishw(path) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
		}
		for i := start; i < len(s); i++ {
			for _ = range lens {
				path = append(path, string(s[i]))
			}
			dfs(i, lens+1)
			path = path[:len(path)-lens]
		}
	}
	dfs(0, 1)
	return ans
}

func TestPartition(t *testing.T) {
	fmt.Println(partition("aab"))
}
