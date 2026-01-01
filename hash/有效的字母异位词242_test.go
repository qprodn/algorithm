package hash

import (
	"fmt"
	"testing"
)

func isAnagram(s string, t string) bool {
	smap := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		smap[s[i]] = smap[s[i]] + 1
	}
	for i := 0; i < len(t); i++ {
		v, exist := smap[t[i]]
		if !exist || v < 1 {
			return false
		}
		smap[t[i]] = smap[t[i]] - 1
	}
	for _, v := range smap {
		if v > 0 {
			return false
		}
	}
	return true
}
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var sDict [26]int

	for i := 0; i < len(s); i++ {
		sDict[s[i]-'a']++
		sDict[t[i]-'a']--
	}

	for _, c := range sDict {
		if c != 0 {
			return false
		}
	}
	return true
}

func TestIsAnagram(t *testing.T) {
	ss := "ab"
	ts := "a"
	fmt.Println(isAnagram(ss, ts))
}
