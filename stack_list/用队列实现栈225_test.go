package stack_list

type MyStack struct {
	in []int
}

func Constructor2() MyStack {
	return MyStack{
		in: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyStack) Pop() int {
	res := this.in[len(this.in)-1]
	this.in = this.in[:len(this.in)-1]
	return res
}

func (this *MyStack) Top() int {
	res := this.in[len(this.in)-1]
	return res
}

func (this *MyStack) Empty() bool {
	return len(this.in) == 0
}
