package string

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	s := "abababab"
	sb := []byte(s)
	j := 0
	next := make([]int, len(s))
	next[0] = 0
	for i := 1; i < len(s); i++ {
		for j > 0 && sb[j] != s[i] {
			j = next[j-1]
		}
		if s[j] == s[i] {
			j++
		}
		next[i] = j
	}
	fmt.Println(next)
}
