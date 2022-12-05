package review

import (
	"fmt"
	"testing"
)

func TestTraversal(t *testing.T) {
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: nil, Right: nil}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}}}
	fmt.Println(lvlorder(tree))
	fmt.Println(inorder(tree))
	fmt.Println(preorder(tree))
	fmt.Println(postorder(tree))
}
