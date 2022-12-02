package interview

// time 100.00%; mem 61.15%
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	r := make([][]int, 0)
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		lr, cnt := make([]int, 0), len(queue)
		for i := 0; i < cnt; i++ {
			// lpop
			node := queue[0]
			queue = queue[1:]
			// rpush
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			// collect level-r
			lr = append(lr, node.Val)
		}
		// collect r
		r = append(r, lr)
	}
	return r
}

func Q102(root *TreeNode) [][]int {
	return levelOrder(root)
}
