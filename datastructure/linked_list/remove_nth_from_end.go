package linked_list

// No.19
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	h := new(ListNode)
	h.Next = head
	p, v := h, h

	for ; n > 0; n-- {
		v = v.Next
	}

	for v != nil {
		p, v = p.Next, v.Next
	}
	p.Next = p.Next.Next
	return h.Next
}
