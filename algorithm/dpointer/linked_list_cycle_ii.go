package dpointer

// 142
// runtime 94.63%; memory 79.44%
func detectCycle(head *ListNode) *ListNode {
	p1, p2 := head, head
	for {
		if p2 == nil || p2.Next == nil {
			return nil
		}
		p1, p2 = p1.Next, p2.Next
		if p1 == p2 {
			break
		}
	}
	p2 = head
	for p1 != p2 {
		p1, p2 = p1.Next, p2.Next
	}
	return p1
}
