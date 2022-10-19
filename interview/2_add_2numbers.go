package interview

type ListNode struct {
	Val  int
	Next *ListNode
}

func add2numbers(add int, l, l1, l2 *ListNode) {
	// l1 == nil && l2 == nil
	if l1 == nil && l2 == nil {
		if add == 1 {
			l.Next = &ListNode{1, nil}
		}
		return
	}
	// l1 == nil && l2 != nil
	// l1 != nil && l2 == nil
	// l1 != nil && l2 != nil
	l.Next = &ListNode{add, nil}
	if l1 != nil {
		l.Next.Val += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		l.Next.Val += l2.Val
		l2 = l2.Next
	}
	if l.Next.Val >= 10 {
		l.Next.Val -= 10
		add = 1
	} else {
		add = 0
	}
	add2numbers(add, l.Next, l1, l2)
}

func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	l := &ListNode{}
	add2numbers(0, l, l1, l2)
	return l.Next
}
