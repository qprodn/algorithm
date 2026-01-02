package hash

import (
	"fmt"
	"testing"
)

func isHappy(n int) bool {
	m := map[int]bool{}
	sum := 0
	for n != 1 && !m[n] {
		m[n] = true
		sum = 0
		for n > 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		n = sum
	}
	return n == 1
}

func isHappy2(n int) bool {
	m := make(map[int]bool)
	for n != 1 && !m[n] {
		m[n] = true
		n = getSum(n)
	}
	return n == 1
}

func getSum(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

func TestIsHappy(t *testing.T) {
	fmt.Println(isHappy(19))
}
