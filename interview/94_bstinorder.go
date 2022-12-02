package interview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// time 100%; mem 68.59%
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	r := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r = append(r, tmp.Val)
		root = tmp.Right
	}
	return r
}

func Q94(root *TreeNode) []int {
	return inorderTraversal(root)
}