package interview

// time 53.56%; mem 34.71%
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	for ; head != nil; head = head.Next {
		nodes = append(nodes, head)
	}
	if len(nodes) == 1 {
		return nil
	}
	// len(nodes)-1 ~ 1
	// len(nodes)-n ~ n
	if len(nodes)-n == 0 {
		return nodes[1]
	}
	node := nodes[len(nodes)-n-1]
	node.Next = node.Next.Next
	return nodes[0]
}

// time 58.55%; mem 34.71%
func removeNthFromEndii(head *ListNode, n int) *ListNode {
	// SPECIAL: len(list) == 1
	if head.Next == nil && n == 1 {
		return nil
	}
	// delNode 到 rNode   n-1
	// lNode   到 delNode 1
	// lNode   到 rNode   n
	r := head
	for ; n > 0; n-- {
		r = r.Next
	}
	// SPECIAL: delete head
	// [1,2,3],n=3
	if r == nil {
		return head.Next
	}
	l := head // prev
	// COMMON
	// [1,2,3,4],n=3
	// [1,2,3,4,5],n=2
	for r.Next != nil {
		l, r = l.Next, r.Next
	}
	l.Next = l.Next.Next
	return head
}

func Q16(head *ListNode, n int) *ListNode {
	return removeNthFromEnd(head, n)
}
