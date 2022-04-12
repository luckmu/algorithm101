package linked_list

// 24
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var p, s *ListNode
	p = head
	if p != nil && p.Next != nil {
		s = p.Next
		p.Next = s.Next
		s.Next = p
		head = s
		for p.Next != nil && p.Next.Next != nil {
			s = p.Next.Next
			p.Next.Next = s.Next
			s.Next = p.Next
			p.Next = s
			p = s.Next
		}
	}
	return head
}

func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := new(ListNode)
	var s *ListNode
	p.Next, s = head, head
	p.Next = s.Next
	s.Next = s.Next.Next
	p.Next.Next = s
	head = p.Next
	p = p.Next.Next
	s = p.Next
	for p.Next != nil && p.Next.Next != nil {
		p.Next = s.Next
		s.Next = s.Next.Next
		p.Next.Next = s
		p = p.Next.Next
		s = p.Next
	}
	return head
}
