package bakctracing

func combine(n int, k int) [][]int {
	var path []int
	var ans [][]int

	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		if len(path) == k {
			tmp := make([]int, k)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := startIndex; i <= n; i++ {
			if n-i+1 < k-len(path) {
				break
			}
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(1)
	return ans
}
