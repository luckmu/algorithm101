package interview

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time 100%; mem 26.64%
func isValidBST(root *TreeNode) bool {
	// 中序遍历是 asc
	// 理解递归的本质是栈, 即stack
	// 而遍历的递归解法 is easy
	prev := -1<<31-1
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tmp.Val <= prev {
			return false
		}
		root = tmp.Right
	}
	return true
}

func Q98(root *TreeNode) bool {
	return isValidBST(root)
}
