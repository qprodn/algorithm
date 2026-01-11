package stack_list

import (
	"fmt"
	"testing"
)

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]int32, 0)
	for _, v := range s {
		if v == '(' || v == '{' || v == '[' {
			stack = append(stack, v)
		} else if v == ')' {
			if len(stack) < 1 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if v == '}' {
			if len(stack) < 1 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		} else if v == ']' {
			if len(stack) < 1 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}

func TestIsValid(t *testing.T) {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("]"))
}
