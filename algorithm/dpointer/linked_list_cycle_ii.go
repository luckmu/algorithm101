package dpointer

// 142
// runtime 94.63%; memory 79.44%
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	var slow, fast *ListNode = head.Next, head.Next.Next
	for fast != nil && fast.Next != nil {
		if slow == fast {
			// circle
			// 1,2,3,4
			fast = head
			for slow != fast {
				slow = slow.Next
				fast = fast.Next
			}
			return fast
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return nil
}
