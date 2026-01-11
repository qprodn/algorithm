package stack_list

import (
	"fmt"
	"strconv"
	"testing"
)

func evalRPN(tokens []string) int {
	stack := make([]int64, 0, len(tokens))
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			i1 := pop(&stack)
			i2 := pop(&stack)
			if token == "+" {
				stack = append(stack, i2+i1)
			}
			if token == "-" {
				stack = append(stack, i2-i1)
			}
			if token == "*" {
				stack = append(stack, i2*i1)
			}
			if token == "/" {
				stack = append(stack, i2/i1)
			}
		} else {
			i, _ := strconv.ParseInt(token, 10, 0)
			stack = append(stack, i)
		}
	}
	return int(stack[0])
}
func pop(stack *[]int64) int64 {
	res := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return res
}

func TestEvalRPN(t *testing.T) {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}
