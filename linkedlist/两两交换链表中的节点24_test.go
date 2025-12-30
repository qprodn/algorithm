package linkedlist

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		one := cur.Next
		two := one.Next
		three := two.Next
		cur.Next = two
		two.Next = one
		one.Next = three
		cur = one
	}
	return dummy.Next
}
