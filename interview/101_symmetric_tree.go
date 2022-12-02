package interview

func issymmetric(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil || right == nil) &&
		!(left != nil && right != nil) {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return issymmetric(left.Left, right.Right) &&
		issymmetric(left.Right, right.Left)
}

// time 100%; mem 99.93%
func isSymmetric(root *TreeNode) bool {
	return issymmetric(root.Left, root.Right)
}

func Q101(root *TreeNode) bool {
	return isSymmetric(root)
}
