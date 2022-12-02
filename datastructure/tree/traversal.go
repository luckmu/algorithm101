package tree

import (
	"fmt"
)

// 前中后序的非递归遍历，
// 因为本质上是用栈模仿递归的过程，
// 考虑递归的流程写代码即可

// No.105
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	dict := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		dict[inorder[i]] = i
	}
	return helper105(preorder, 0, len(preorder)-1, 0, dict)
}

func helper105(preorder []int, preStart, preEnd, inStart int, dict map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootval := preorder[preStart]
	rootidx := dict[rootval]
	fmt.Println(rootval)

	leftlen := rootidx - inStart
	root := &TreeNode{Val: rootval}
	root.Left = helper105(preorder, preStart+1, preStart+1+leftlen-1, inStart, dict)
	root.Right = helper105(preorder, preStart+leftlen+1, preEnd, rootidx+1, dict)
	return root
}

// No.144
func preorderTraversal(root *TreeNode) []int {
	// 递归本质是栈调用，非递归写法用栈来实现前序遍历
	// stack := make([]*TreeNode, 0)
	ret := make([]int, 0)
	stack := []*TreeNode{root}
	push := func(node *TreeNode) {
		stack = append(stack, node)
	}
	pop := func() *TreeNode {
		n := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		return n
	}
	for len(stack) > 0 {
		n := pop()
		ret = append(ret, n.Val)
		if n.Right != nil {
			push(n.Right)
		}
		if n.Left != nil {
			push(n.Left)
		}
	}
	return ret
}

func preorderTraversal2(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			ret = append(ret, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = tmp.Right
	}
	return ret
}

// No.94
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, tmp.Val)
		root = tmp.Right
	}
	return ret
}

// No.145
// 每个 node 在从 stack pop 前需要遍历右子树，
// hash table 用来标记当前节点有没有遍历过右子树，遍历过则pop，没遍历过则 root = tmp.Right 进行遍历
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	stack := make([]*TreeNode, 0)
	hash := make(map[*TreeNode]bool)

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		if tmp != nil && hash[tmp] {
			stack = stack[:len(stack)-1]
			ret = append(ret, tmp.Val)
		} else {
			hash[tmp] = true
			root = tmp.Right
		}
	}
	return ret
}
