package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 206
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nxtCur := cur.Next
		cur.Next = prev
		prev, cur = cur, nxtCur
	}
	return prev
}

func reverseList2(head, prev *ListNode) *ListNode {
	if head != nil {
		return prev
	}
	var nxt *ListNode = head.Next
	head.Next = prev
	return reverseList2(nxt, head)
}
