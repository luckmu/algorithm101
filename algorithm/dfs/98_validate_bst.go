package dfs

import "fmt"

func IsValidBST(root *TreeNode) bool {
	// left is smaller than root
	// right is bigger than root
	// left < root < right
	stack := make([]*TreeNode, 0)
	prev := -1<<31-1 // int64
	fmt.Println(prev)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Val <= prev && prev != -1<<31 {
			return false
		}
		prev = root.Val
		root = root.Right
	}
	return true
}
