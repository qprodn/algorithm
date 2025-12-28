package array

import (
	"fmt"
	"testing"
)

// 没做出来
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range n {
		in := make([]int, n)
		ans[i] = in
	}
	i, j := 0, 0
	b := n - 1
	f := 1
	num := 1
	for b > 1 {
		if f%4 == 1 {
			for range b {
				ans[i][j] = num
				num++
				j++
			}
			f++
		} else if f%4 == 2 {
			for range b {
				ans[i][j] = num
				num++
				i++
			}
			f++
		} else if f%4 == 3 {
			for range b {
				ans[i][j] = num
				num++
				j--
			}
			f++
		} else if f%4 == 0 {
			for range b {
				ans[i][j] = num
				num++
				i--
			}
			f++
			b--
		}
	}
	return ans
}

func generateMatrix2(n int) [][]int {
	var x, y int
	loop := n / 2
	center := n / 2
	ans := make([][]int, n)
	for i := range n {
		in := make([]int, n)
		ans[i] = in
	}
	offset := 1
	num := 1
	for loop > 0 {
		i, j := x, y
		for ; j < n-offset; j++ {
			ans[i][j] = num
			num++
		}
		for ; i < n-offset; i++ {
			ans[i][j] = num
			num++
		}
		for ; j > x; j-- {
			ans[i][j] = num
			num++
		}
		for ; i > y; i-- {
			ans[i][j] = num
			num++
		}
		x++
		y++
		offset++
		loop--
	}
	if n%2 == 1 {
		ans[center][center] = n * n
	}
	return ans
}

func TestGenerateMatrix(t *testing.T) {
	ans := generateMatrix2(5)
	for _, ints := range ans {
		for _, i := range ints {
			fmt.Print(i, ":")
		}
		fmt.Println()
	}
}
