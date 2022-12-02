package dfs

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// InorderDFSBinaryTree
// dfs 用 stack,
func InorderDFS(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}

func InorderRe(root *TreeNode, res *[]int) {
	// inorder recursive
	if root == nil {
		return
	}
	InorderRe(root.Left, res)
	*res = append(*res, root.Val)
	InorderRe(root.Right, res)
}
