package linked_list

func isPalindrome(head *ListNode) bool {
	// stack + O(n)
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for i, j := 0, len(stack)-1; i <= j; i, j = i+1, j-1 {
		if stack[i] != stack[j] {
			return false
		}
	}
	return true
}
