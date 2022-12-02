package linked_list

// A, B intersect at C
// A->C:
// B->C:
// C->END:

// A->END:
// B->END:
// No. 160
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := headA, headB

	for l1 != nil && l2 != nil {
		l1, l2 = l1.Next, l2.Next
	}

	if l2 == nil {
		// always make l1 short
		// l1, l2 = l2, l1
		l2 = l1
		headA, headB = headB, headA
	}
	l1 = headA
	// l1 is always shorter
	for l2 != nil {
		l1 = l1.Next
	}
	l2 = headB

	for l1 != l2 && l1 != nil && l2 != nil {
		l1, l2 = l1.Next, l2.Next
	}
	if l1 != l2 {
		return nil
	}
	return l1
	// l1, l2 := headA, headB
	// for l1 != l2 {
	// 	if l1 != nil {
	// 		l1 = l1.Next
	// 	} else {
	// 		l1 = headB
	// 	}
	// 	if l2 != nil {
	// 		l2 = l2.Next
	// 	} else {
	// 		l2 = headA
	// 	}
	// }
	// return l1
}
