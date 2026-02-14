package greedy

import "sort"

func findContentChildren(g []int, s []int) int {
	sum := 0
	sort.Ints(g)
	sort.Ints(s)
	ig, is := 0, 0
	for ig < len(g) && is < len(s) {
		if s[is] >= g[ig] {
			sum++
			ig++
		}
		is++
	}
	return sum
}
