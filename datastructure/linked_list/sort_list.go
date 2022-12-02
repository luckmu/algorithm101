package linked_list

// No.148
func sortList(head *ListNode) *ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	if length <= 1 {
		return head
	}

	middleNode := middleNode(head)
	cur = middleNode.Next
	middleNode.Next = nil
	middleNode = cur

	left := sortList(head)
	right := sortList(middleNode)
	return mergehelper(left, right)
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head
	for p2.Next != nil && p2.Next.Next != nil {
		p1, p2 = p1.Next, p2.Next.Next
	}
	return p1
}

func mergehelper(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergehelper(l1.Next, l2)
		return l1
	}
	l2.Next = mergehelper(l1, l2.Next)
	return l2
}
