package test

import (
	"algorithm101/algorithm/dfs"
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	want := []int{2, 3, 1, 4}
	get := make([]int, 0)
	root := dfs.TreeNode{1, &dfs.TreeNode{2, nil, &dfs.TreeNode{3, nil, nil}}, &dfs.TreeNode{4, nil, nil}}

	get = dfs.InorderDFS(&root)
	if !reflect.DeepEqual(get, want) {
		t.Log("want:", want, "but get:", get)
	}
	get = make([]int, 0)
	dfs.InorderRe(&root, &get)
	if !reflect.DeepEqual(get, want) {
		t.Log("want:", want, "but get:", get)
	}

	dfs.IsValidBST(&root)
}
