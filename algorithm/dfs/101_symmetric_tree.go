package dfs

func symmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	} else if (left == nil || right == nil) &&
		!(left != nil && right != nil) {
		// XOR: (x || y) && !(x && y)
		return false
	} else if left.Left.Val != right.Right.Val ||
		left.Right.Val != right.Left.Val {
		return false
	}
	return symmetric(left.Left, right.Right) &&
		symmetric(left.Right, right.Left)
}

func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}
