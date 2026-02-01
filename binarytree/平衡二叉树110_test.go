package binarytree

import "math"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		left := getMaxDepth(node.Left)
		right := getMaxDepth(node.Right)
		if math.Abs(float64(left-right)) > 1 {
			return false
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return true
}
func getMaxDepth(root *TreeNode) int {
	depth := 0
	if root == nil {
		return depth
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return depth
}
