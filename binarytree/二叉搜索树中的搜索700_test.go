package binarytree

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	var result *TreeNode
	if root.Val > val {
		result = searchBST(root.Left, val)
	}
	if root.Val < val {
		result = searchBST(root.Right, val)
	}
	return result
}
