# linked list

基础操作: reverse, merge, swap

Tips:
+ 操作2个链节，考虑3个指针
+ 操作 head 要用 dummy (`dummy := &ListNode{Next: head}`)

## Reverse Linklist

Leetcode [206](https://leetcode.cn/problems/reverse-linked-list/), [92](https://leetcode.cn/problems/reverse-linked-list-ii/), [25](https://leetcode.cn/problems/reverse-nodes-in-k-group/)

**核心逻辑**

```go
// pre, cur, nxt
pre := nil
for cur != nil {
    nxt := cur.Next
    cur.Next = pre
    pre = cur
    cur = nxt
}
```
