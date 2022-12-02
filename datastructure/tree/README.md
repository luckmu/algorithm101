# tree

树的题型：总之是要遍历

无论是：真正意义上的遍历，求高度，求直径

遍历:
```go
func traverse(root *TreeNode) {
    if root == nil {
        return
    }
    doSomething()
    traverse(root.Left)
    traverse(root.Right)
}
```

层次遍历：
+ 队列 + 双层循环（外层保证层次遍历所有 `for len(qs) != 0`、内存保证层次标记 `level_count = len(qs); for i := 0; i < level_count; i++`）
+ `level_count` 记录当前层节点数量

前中后序遍历：

```go
preorder(root *TreeNode) {
    visit(root)
    preorder(root.Left)
    preorder(root.Right)
}

inorder(root *TreeNode) {
    inorder(root.Left)
    visit(root)
    inorder(root.Right)
}

postorder(root *TreeNode) {
    postorder(root.Left)
    postorder(root.Right)
    visit(root)
}
```


## build tree from preorder & inorder

找到 preorder 中的 rootval, 根据这个 rootval 划分 inorder 中的左右子树

```go
// func BuildTree() *TreeNode
// 自顶向下构建, 用递归, 前序递归
// 为什么是前序? 按照 root, root.Left, root.Right 这个顺序赋值
func helper() {
    root := &TreeNode{}
    root.Left = helper()
    root.Right = helper()
}
// 结束条件? 每次从 preorder 中取 node
//   root.Left = helper() // while helper() returns nil
//   
// 参数范围, 每次递归都要缩小 preorder 和 inorder 的区间
// 因为是前序递归, 依次从 preorder 中取元素即可
// 

```

## preorder traversal

递归和非递归写法

// 递归本质是栈调用，非递归写法用`栈`来实现

```go
// 递归写法
func helper() {
    if root == nil {
        return
    }
    print(root.Val)
    recursive(root.Left)
    recursice(root.Right)
    return
}
// 非递归写法
func helper() {q
    stack := []*TreeNode{root}
    // define push() and pop()
    for len(stack) > 0 {
        n := pop()
        print(n.Val)
        // 注意 right 和 left 的 push 顺序
        if n.Right != nil {
            push(n.Right)
        }
        if n.Left != nil {
            push(n.Left)
        }
    }
    return
}
```

## 669. trim a binary search tree

保留在 [low, high] 间的节点

遍历同时删除
1. 当前节点小于 low，root.Left = trimBST(root.Right, low, high)
2. 当前节点大于 high，root.Right = trimBST(root.Left, low, high)
