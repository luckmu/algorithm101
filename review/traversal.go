package review

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func inorder(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, tmp.Val)
		root = tmp.Right
	}
	return ans
}

func preorder(root *TreeNode) []int {
	// if root.Left {}
	// root
	// if root.Right {}
	ans := []int{}
	stack := []*TreeNode{}
	for len(stack) > 0 || root != nil {
		for root != nil {
			ans = append(ans, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = tmp.Right
	}
	return ans
}

func postorder(root *TreeNode) []int {
	ans := []int{}
	stack := []*TreeNode{}
	hash := make(map[*TreeNode]bool)
	// if root.Left {}
	// if root.Right {}
	// root
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		tmp := stack[len(stack)-1]
		if hash[tmp] {
			// right 完全遍历, 退栈
			stack = stack[:len(stack)-1]
			ans = append(ans, tmp.Val)
		} else {
			// 遍历 right
			hash[tmp] = true
			root = tmp.Right
		}
	}
	return ans
}

func lvlorder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	q, cnt, nxtcnt := []*TreeNode{root}, 1, 0
	ans := make([][]int, 0)
	for len(q) > 0 {
		tmpans := make([]int, 0)
		for ; cnt > 0; cnt-- {
			// pop
			n := q[0]
			q = q[1:]
			tmpans = append(tmpans, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
				nxtcnt++
			}
			if n.Right != nil {
				q = append(q, n.Right)
				nxtcnt++
			}
		}
		ans = append(ans, tmpans)
		cnt, nxtcnt = nxtcnt, 0
	}
	return ans
}
