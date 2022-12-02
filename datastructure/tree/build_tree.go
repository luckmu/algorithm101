package tree

// 从“前，中，后”序构造树
// 1. 含有中序的两个题目：
// 	 1.1. 前，中：
//   1.2. 中，后：
//     通过前序的first或后序的last来将中序分为左右两棵子树
// 2. 前，后序：
//   *首先前、后序不能唯一定义一棵二叉树，[1,2], [2,1] 就有两种形态
//   root := &TreeNode{...}
//   root.Left =
//   root.Right =

// No.105
// preorder's first divide inorder into 2 slices
func constructFromPreIn(preorder []int, inorder []int) *TreeNode {
	// preorderLen := len(preorder)
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = constructFromPreIn(preorder[:len(inorder[:pos])], inorder[:pos])
			root.Right = constructFromPreIn(preorder[len(inorder[:pos]):], inorder[pos+1:])
		}
	}
	return root
}

// No.106
// postorder's last divide inorder into 2 slices
func constructFromInPost(inorder []int, postorder []int) *TreeNode {
	postorderLen := len(postorder)
	if len(postorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[postorderLen-1]}
	postorder = postorder[:postorderLen-1]
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = constructFromInPost(inorder[:pos], postorder[:len(inorder[:pos])])
			root.Right = constructFromInPost(inorder[pos+1:], postorder[len(inorder[:pos]):])
		}
	}
	return root
}

// No.889
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}
	// has left tree
	leftval := preorder[1]
	for pos, val := range postorder {
		if val == leftval {
			// preorder: 左子树节点数量
			//   pos 是 postorder 的 root, 即左子树的 root, 也即左子树共 pos+1 个节点,
			//   preorder 从 1 开始, 所以 preorder[1:1+pos+1]
			root.Left = constructFromPrePost(preorder[1:pos+2], postorder[:pos+1])
			root.Right = constructFromPrePost(preorder[pos+2:], postorder[pos+1:len(postorder)-1])
		}
	}
	return root
}
