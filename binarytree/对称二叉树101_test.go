package binarytree

func isSymmetric(root *TreeNode) bool {
	return compare(root.Left, root.Right)
}

func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if left == nil && right != nil {
		return false
	} else if left != nil && right == nil {
		return false
	} else if left.Val != right.Val {
		return false
	}
	out := compare(left.Left, right.Right)
	in := compare(left.Right, right.Left)
	return out && in
}
