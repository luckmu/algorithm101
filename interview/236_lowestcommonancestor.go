package interview

// 如果当前节点是 x
// 成立条件: (fx_leftchild && fx_rightchild) || ((fx==p || fx==q) && (fx_leftchild || fx_rightchild))
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// fx==p || fx==q
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// fx_leftchild && fx_rightchild
	if left != nil && right != nil {
		return root
	}
	// fx_leftchild || fx_rightchild
	if left == nil {
		return right
	}
	return left
}
