package lang

import (
	"fmt"
	"strings"
	"testing"
)

func TestSB(t *testing.T) {
	var sb strings.Builder
	sb.WriteString("123")
	s1 := sb.String()
	s2 := sb.String()
	fmt.Println("s1 == s2:", s1 == s2)
	sb.WriteString("456")
	fmt.Println("s2:", s2)
	s3 := sb.String()
	fmt.Println("s1 == s3", s1 == s3)
}
