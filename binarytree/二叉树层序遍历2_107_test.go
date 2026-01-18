package binarytree

func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	list := make([]*TreeNode, 0)
	list = append(list, root)
	for len(list) > 0 {
		count := len(list)
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
		ans = append(ans, c)
	}
	l, r := 0, len(ans)-1
	for l < r {
		ans[l], ans[r] = ans[r], ans[l]
		l++
		r--
	}
	return ans
}
