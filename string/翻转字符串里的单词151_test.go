package string

import (
	"fmt"
	"strings"
	"testing"
)

func reverseWords(s string) string {
	bytes := []byte(s)
	tmps := make([]string, 0, len(s))
	length := len(s)
	is := 0
	for i := range bytes {
		if (bytes[i] == ' ' || i == length-1) && is != i {
			if i == length-1 {
				tmps = append(tmps, string(bytes[is:i+1]))
			} else {
				tmps = append(tmps, string(bytes[is:i]))
			}
			is = i
		}
		if bytes[is] == ' ' {
			is++
		}
	}
	sb := strings.Builder{}
	for i := len(tmps) - 1; i >= 0; i-- {
		if tmps[i] == " " {
			continue
		}
		sb.WriteString(tmps[i])
		if i != 0 {
			sb.WriteString(" ")
		}
	}
	return sb.String()
}

func reverseWords2(s string) string {
	bytes := []byte(s)
	slow := 0
	fast := 0
	for len(bytes) > 0 && fast < len(bytes) && bytes[fast] == ' ' {
		fast++
	}
	for ; fast < len(bytes); fast++ {
		if fast > 0 && bytes[fast-1] == bytes[fast] && bytes[fast] == ' ' {
			continue
		}
		bytes[slow] = bytes[fast]
		slow++
	}
	if slow-1 > 0 && bytes[slow-1] == ' ' {
		bytes = bytes[:slow-1]
	} else {
		bytes = bytes[:slow]
	}
	reverseX(bytes)
	slow = 0
	for i := range bytes {
		if i == len(bytes)-1 && i > slow {
			reverseX(bytes[slow : i+1])
		}
		if bytes[i] == ' ' && i > slow {
			reverseX(bytes[slow:i])
			slow = i + 1
		}
	}
	return string(bytes)
}

func reverseX(bytes []byte) {
	l, r := 0, len(bytes)-1
	for l < r {
		bytes[l], bytes[r] = bytes[r], bytes[l]
		l++
		r--
	}
}

func TestReverseWords(t *testing.T) {
	words := reverseWords2("  asdasd  df f")
	fmt.Println(words)
}
