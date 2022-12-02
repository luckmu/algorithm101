package tree

// about recursive tree

// 很多时候，树递归的写法与深度优先搜索的递归写法相同，不作区分

func max(nums ...int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

// No.104
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left)+1, maxDepth(root.Right)+1)
}

// No.110
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lh := height(root.Left)
	rh := height(root.Right)
	return ((lh >= rh && lh-rh <= 1) || (rh > lh && rh-lh <= 1)) && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left)+1, height(root.Right)+1)
}

// No.543
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
	checkDiameter(root, &result)
	return result
}

func checkDiameter(root *TreeNode, result *int) int {
	if root == nil {
		return 0
	}
	left := checkDiameter(root.Left, result)
	right := checkDiameter(root.Right, result)
	*result = max(*result, left+right)
	return max(left, right) + 1
}

// No.437
func pathSum(root *TreeNode, targetSum int) int {
	// Vars:
	//   curSum: current_sum
	//   curPaths: current_number_of_paths

	// 遍历 + 递归
	// 每个节点都是单次递归的顶点
	curPaths := 0
	traverse(root, targetSum, 0, &curPaths)
	return curPaths
}

func traverse(root *TreeNode, targetSum, curSum int, curPaths *int) {
	if root == nil {
		return
	}
	assist437(root, targetSum, curSum, curPaths)
	traverse(root.Left, targetSum, curSum, curPaths)
	traverse(root.Right, targetSum, curSum, curPaths)
}

// assist437, recursive part
func assist437(root *TreeNode, targetSum, curSum int, curPaths *int) {
	if root == nil {
		return
	}
	curSum += root.Val
	if curSum == targetSum {
		*curPaths += 1
	}
	assist437(root.Left, targetSum, curSum, curPaths)
	assist437(root.Right, targetSum, curSum, curPaths)
}

// No.101
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return assist101(root.Left, root.Right)
}

func assist101(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left != nil || right != nil) && (left == nil || right == nil) {
		return false
	}

	// left != nil && right != nil
	return left.Val == right.Val &&
		assist101(left.Left, right.Right) &&
		assist101(left.Right, right.Left)
}

// No.1110
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	forest := make([]*TreeNode, 0)
	dict := make(map[int]struct{}, 0)
	for _, v := range to_delete {
		dict[v] = struct{}{}
	}
	root = helper1110(root, dict, &forest)
	if root != nil {
		forest = append(forest, root)
	}
	return forest
}

func helper1110(root *TreeNode, dict map[int]struct{}, forest *[]*TreeNode) *TreeNode {
	// 1. 将 root 加入结果集
	// 2. 删除当前节点，将左右节点中非 nil 节点加入结果集
	if root == nil {
		return root
	}

	root.Left = helper1110(root.Left, dict, forest)
	root.Right = helper1110(root.Right, dict, forest)
	if _, ok := dict[root.Val]; ok {
		if root.Left != nil {
			*forest = append(*forest, root.Left)
		}
		if root.Right != nil {
			*forest = append(*forest, root.Right)
		}
		root = nil
	}
	return root
}
