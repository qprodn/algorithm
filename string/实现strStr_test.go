package string

import (
	"fmt"
	"testing"
)

func strStr(haystack string, needle string) int {
	h, t, c := 0, 0, 0
	ans := -1
	for ; h < len(haystack); h++ {
		if haystack[h] == needle[t] {
			if t == 0 {
				ans = h
			}
			c = h
			for c < len(haystack) {
				if haystack[c] == needle[t] {
					if t == len(needle)-1 {
						return ans
					}
					t++
					c++
				} else {
					break
				}
			}
			t = 0
			c = 0
		} else {
			ans = -1
		}
	}
	return -1
}
func strStr2(haystack string, needle string) int {
	n := len(haystack)
	m := len(needle)
	for i := 0; i < n-m+1; i++ {
		flag := true
		for j := 0; j < m; j++ {
			if haystack[i+j] != needle[j] {
				flag = false
				break
			}
		}
		if flag {
			return i
		}
	}
	return -1
}

func strStr3(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	next := getNext2(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		if haystack[i] == needle[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func getNext(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

func getNext2(s string) []int {
	j := 0
	next := make([]int, len(s))
	next[0] = 0
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
}

func TestStrStr(t *testing.T) {
	str := strStr3("sadbutsad", "sad")
	fmt.Println(str)
}
