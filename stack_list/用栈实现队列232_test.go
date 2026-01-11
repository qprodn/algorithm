package stack_list

import (
	"fmt"
	"testing"
)

type MyStack2 []int

func (s *MyStack2) Push(v int) {
	*s = append(*s, v)
}

func (s *MyStack2) Pop() int {
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}

func (s *MyStack2) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *MyStack2) Size() int {
	return len(*s)
}

func (s *MyStack2) Empty() bool {
	return s.Size() == 0
}

// ---------- 分界线 ----------

type MyQueue struct {
	stackIn  *MyStack2
	stackOut *MyStack2
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  &MyStack2{},
		stackOut: &MyStack2{},
	}
}

func (this *MyQueue) Push(x int) {
	this.stackIn.Push(x)
}

func (this *MyQueue) Pop() int {
	this.fillStackOut()
	return this.stackOut.Pop()
}

func (this *MyQueue) Peek() int {
	this.fillStackOut()
	return this.stackOut.Peek()
}

func (this *MyQueue) Empty() bool {
	return this.stackIn.Empty() && this.stackOut.Empty()
}

// fillStackOut 填充输出栈
func (this *MyQueue) fillStackOut() {
	if this.stackOut.Empty() {
		for !this.stackIn.Empty() {
			val := this.stackIn.Pop()
			this.stackOut.Push(val)
		}
	}
}

func TestMyStack(t *testing.T) {
	myQueue := Constructor()
	myQueue.Push(1)
	//myQueue.Push(2)
	//fmt.Println(myQueue.Peek())
	fmt.Println(myQueue.Pop())
	fmt.Println(myQueue.Empty())

}
