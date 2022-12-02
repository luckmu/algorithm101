package tree

import (
	"math"
)

// No.226
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

// No.617
func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Right = mergeTrees(root1.Right, root2.Right)
	root1.Left = mergeTrees(root1.Left, root2.Left)
	return root1
}

// No.404
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	}
	return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}

// No.513
func findBottomLeftValue(root *TreeNode) int {
	qs := []*TreeNode{root}
	var firstVal int
	for len(qs) > 0 {
		curlevelnodes := len(qs)
		for i := 0; i < curlevelnodes; i++ {
			// this level
			q := qs[0]
			qs = qs[1:]
			if q.Left != nil {
				qs = append(qs, q.Left)
			}
			if q.Right != nil {
				qs = append(qs, q.Right)
			}
			if i == 0 {
				firstVal = q.Val
			}
		}
	}
	return firstVal
}

// No.538
func convertBST(root *TreeNode) *TreeNode {

	// greater than `key` -> sum is `value`
	// dict := make(map[int]int)

	// ***
	// 保存中途结果的方法不只有 dict
	// 递归调用的栈也可以保存变量
	// ***

	if root == nil {
		return nil
	}
	var sum int
	helper538(root, &sum)
	return root
}

func helper538(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	helper538(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	helper538(root.Left, sum)
}

// No.235
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == nil || q == nil || root == nil {
		return nil
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

// 中序遍历 BST 是有序的
func getMinimumDifference(root *TreeNode) int {
	// number of nodes is >= 2
	ret, tmp := math.MaxInt16, -1
	helper(root, &ret, &tmp)
	return ret
}

func helper(root *TreeNode, ret, tmp *int) {
	if root == nil {
		return
	}
	helper(root.Left, ret, tmp)
	if t := subabs(root.Val, *tmp); *tmp != -1 && t < *ret {
		*ret = t
	}
	*tmp = root.Val
	helper(root.Right, ret, tmp)
}

func subabs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
