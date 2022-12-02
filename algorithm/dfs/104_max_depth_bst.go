package dfs

func max104(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// DFS: recursive
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max104(MaxDepth(root.Left)+1,
		MaxDepth(root.Right)+1)
}

// BFS: also, stack + levelorder traversal bst
