package interview

// time 79.19%; mem 74.44%
func kthSmallest(root *TreeNode, k int) int {
	// inorder traversal
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		if k == 0 {
			return node.Val
		}
		root = node.Right
	}
	return -1
}

func Q120(root *TreeNode, k int) int {
	return kthSmallest(root, k)
}
