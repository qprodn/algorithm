package bakctracing

import (
	"fmt"
	"testing"
)

func letterCombinations(digits string) []string {
	kv := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	path, ans := "", make([]string, 0)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == len(digits) {
			ans = append(ans, path)
		}
		if start >= len(digits) {
			return
		}
		s := string(digits[start])
		for i := 0; i < len(kv[s]); i++ {
			path += kv[s][i]
			dfs(start + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}

func TestLetterCombinationse(t *testing.T) {
	i := letterCombinations("233")
	fmt.Println(i)
}
