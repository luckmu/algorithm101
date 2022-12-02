package interview

func mergeNodes(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tmp, tmp1, tmp2 := dummyHead, head1, head2
	for tmp1 != nil && tmp2 != nil {
		if tmp1.Val <= tmp2.Val {
			tmp.Next = tmp1
			tmp1 = tmp1.Next
		} else {
			tmp.Next = tmp2
			tmp2 = tmp2.Next
		}
		tmp = tmp.Next
	}
	if tmp1 != nil {
		tmp.Next = tmp1
	} else if tmp2 != nil {
		tmp.Next = tmp2
	}
	return dummyHead.Next
}

func sortNodes(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow
	return mergeNodes(sortNodes(head, mid), sortNodes(mid, tail))
}

func sortList(head *ListNode) *ListNode {
	return sortNodes(head, nil)
}

func Q148(head *ListNode) *ListNode {
	return sortList(head)
}
