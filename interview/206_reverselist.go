package interview

// h1 -> h2 -> h3 -> h4
// h1 <- h2 <- h3 <- h4
// func reverseListii(head *ListNode) *ListNode {
// 	if head.Next == nil {
// 		return head
// 	}
// 	// head.Next == tail
// 	node := reverseList(head.Next)
// 	head.Next.Next = head
// 	head.Next = nil
// 	return node
// }

// time 100%; mem 99.93%
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	p, pp := head, head.Next
	if pp == nil {
		return p
	}
	for pp != nil {
		ahead := pp.Next
		pp.Next = p
		if p == head {
			p.Next = nil
		}
		p, pp = pp, ahead
	}
	return p
}

func Q206(head *ListNode) *ListNode {
	return reverseList(head)
}
