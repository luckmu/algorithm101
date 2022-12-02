package interview

// time 92.86%; mem 82.87%
func buildTree(preorder []int, inorder []int) *TreeNode {
	// head, left, right
	// left, head, right
	// [preorder]'s leftmost will divide [inorder] into [leftchildtree] & [rightchildtree]
	if len(preorder) <= 0 {
		return nil
	}
	head := &TreeNode{preorder[0], nil, nil}
	var i int
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	// left nodes:  [0,i) = i
	// right nodes: [i+1:]
	head.Left = buildTree(preorder[1:1+i], inorder[:i])
	head.Right = buildTree(preorder[1+i:], inorder[i+1:])
	return head
}

func Q105(preorder []int, inorder []int) *TreeNode {
	return buildTree(preorder, inorder)
}
