package binarytree

func levelOrderBottom(root *TreeNode) [][]int {
	count := 1
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		c := make([]int, 0)
		for i := 0; i < count; i++ {
			node := list[0]
			list = list[1:]
			if node.Left != nil {
				list = append(list, node.Left)
			}
			if node.Right != nil {
				list = append(list, node.Right)
			}
			c = append(c, node.Val)
		}
		count = len(list)
		ans = append(ans, c)
	}
	return ans
}
