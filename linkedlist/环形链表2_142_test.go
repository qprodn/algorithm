package linkedlist

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			id1 := fast
			id2 := head
			for id1 != id2 {
				id1 = id1.Next
				id2 = id2.Next
			}
			return id1
		}
	}
	return fast
}

func detectCycleMap(head *ListNode) *ListNode {
	mymap := make(map[*ListNode]struct{})
	for head != nil {
		if _, exist := mymap[head]; exist {
			return head
		}
		mymap[head] = struct{}{}
		head = head.Next
	}
	return nil
}
