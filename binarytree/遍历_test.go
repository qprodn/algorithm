package binarytree

import "container/list"

func preorderTraversal(root *TreeNode) (res []int) {
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		//前
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)

		//中
		//traversal(node.Left)
		//res = append(res, node.Val)
		//traversal(node.Right)

		//后
		//traversal(node.Left)
		//traversal(node.Right)
		//res = append(res, node.Val)
	}
	traversal(root)
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}

	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}
