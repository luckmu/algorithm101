package tree

import (
	"testing"
)

func transferTreeNodes(nums []interface{}) *TreeNode {
	if len(nums) < 1 {
		return nil
	}
	rootval, ok := nums[0].(int)
	if !ok {
		return nil
	}
	var root, q *TreeNode
	root = &TreeNode{rootval, nil, nil}
	nqs := []*TreeNode{root}
	getq := func(q *TreeNode, i int) *TreeNode {
		if i%2 != 0 {
			r := nqs[0]
			nqs = nqs[1:]
			return r
		}
		return q
	}
	for i := 1; i < len(nums); i++ {
		q = getq(q, i)
		switch k := nums[i].(type) {
		case int:
			node := &TreeNode{k, nil, nil}
			nqs = append(nqs, node)
			if i%2 != 0 {
				q.Left = node
			} else {
				q.Right = node
			}
		case nil:
			continue
		}
	}
	return root
}

func TestTree(t *testing.T) {
	// fmt.Println(isBalanced(&TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}))
	// fmt.Println(isBalanced(&TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, nil}}))

	// fmt.Println(pathSum(transferTreeNodes([]interface{}{-2, nil, -3}), -5))
	// fmt.Println(pathSum(transferTreeNodes([]interface{}{1, -2, -3, 1, 3, -2, nil, -1}), -1))

	// fmt.Println(isSymmetric(transferTreeNodes([]interface{}{1, 2, 2, nil, 3, nil, 3})))

	// roots := delNodes(transferTreeNodes([]interface{}{1, 2, 3, 4, 5, 6, 7}), []int{3, 5})
	// for _, root := range roots {
	// 	fmt.Printf("%d ", root.Val)
	// }
	// fmt.Println()

	// fmt.Println(averageOfLevels(transferTreeNodes([]interface{}{3, 9, 20, nil, nil, 15, 7})))

	// buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	// trie := NewTrie()
	// trie.Insert([]rune("apple"))
	// fmt.Println(trie.Search([]rune("apple")))   // true
	// fmt.Println(trie.Search([]rune("app")))     // false
	// fmt.Println(trie.StartsWith([]rune("app"))) // true
	// trie.Insert([]rune("app"))
	// fmt.Println(trie.Search([]rune("app"))) // true

	// fmt.Printf("sumOfLeftLeaves(): %v\n", sumOfLeftLeaves(transferTreeNodes([]interface{}{3, 9, 20, nil, nil, 15, 7})))

	// fmt.Printf("findBottomLeftValue(transferTreeNodes([]interface{}{2, 1, 3})): %v\n", findBottomLeftValue(transferTreeNodes([]interface{}{2, 1, 3})))

	// fmt.Printf("getMinimumDifference(transferTreeNodes([]interface{}{1, 0, 48, nil, nil, 12, 49})): %v\n", getMinimumDifference(transferTreeNodes([]interface{}{1, 0, 48, nil, nil, 12, 49})))

	// constructFromPreIn([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

	inorderTraversal := func(root *TreeNode) []int {
		if root == nil {
			return []int{}
		}
		stack := []*TreeNode{root}
		ret := make([]int, 0)

		for len(stack) > 0 {
			// pop
			node := stack[len(stack)-1]
			if node.Left != nil {
				stack = append(stack, root.Left)
				continue
			}
			ret = append(ret, root.Val)
			if node.Right != nil {
				stack = append(stack, root.Right)
				continue
			}
			stack = stack[:len(stack)-1]
		}
		return ret
	}
	inorderTraversal(transferTreeNodes([]interface{}{1, nil, 2, nil, nil, 3, nil}))
}
