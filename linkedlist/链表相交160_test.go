package linkedlist

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ca := headA
	cb := headB
	for ca != cb {
		if ca != nil {
			ca = ca.Next
		} else {
			ca = headB
		}
		if cb != nil {
			cb = cb.Next
		} else {
			cb = headA
		}
	}
	return nil
}
