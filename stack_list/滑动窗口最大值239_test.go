package stack_list

import (
	"fmt"
	"testing"
)

func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, 0)
	maxNum := -100000
	maxNumIdx := -1
	for i := 0; i <= len(nums)-k; i++ {
		if maxNumIdx < i {
			maxNum = -100000
			for j := i; j < i+k; j++ {
				if nums[j] > maxNum {
					maxNum = nums[j]
					maxNumIdx = j
				}
			}
		}
		if nums[i+k-1] > maxNum {
			maxNum = nums[i+k-1]
			maxNumIdx = i + k - 1
		}
		ans = append(ans, maxNum)
	}
	return ans
}

// time out
func maxSlidingWindow2(nums []int, k int) []int {
	ans := make([]int, 0)
	for i := 0; i <= len(nums)-k; i++ {
		maxNum := -100000
		for j := i; j < i+k; j++ {
			if nums[j] > maxNum {
				maxNum = nums[j]
			}
		}
		ans = append(ans, maxNum)
	}
	return ans
}

type MyQueue3 struct {
	queue []int
}

func NewMyQueue() *MyQueue3 {
	return &MyQueue3{
		queue: make([]int, 0),
	}
}

func (m *MyQueue3) Front() int {
	return m.queue[0]
}

func (m *MyQueue3) Back() int {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue3) Empty() bool {
	return len(m.queue) == 0
}

func (m *MyQueue3) Push(val int) {
	for !m.Empty() && val > m.Back() {
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, val)
}

func (m *MyQueue3) Pop(val int) {
	if !m.Empty() && val == m.Front() {
		m.queue = m.queue[1:]
	}
}

func maxSlidingWindow3(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	// 先将前k个元素放入队列
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}
	// 记录前k个元素的最大值
	res = append(res, queue.Front())

	for i := k; i < length; i++ {
		// 滑动窗口移除最前面的元素
		queue.Pop(nums[i-k])
		// 滑动窗口添加最后面的元素
		queue.Push(nums[i])
		// 记录最大值
		res = append(res, queue.Front())
	}
	return res
}

func TestMaxSlidingWindow(t *testing.T) {
	//fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}
