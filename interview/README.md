# interview

[top interview questions](https://leetcode.com/problem-list/top-interview-questions/)

1. **子串(substring, subarray, max, longest)问题直接用dp**: 最长无重复子串, 最长回文子串, 最大和子数组(可以简化dp为1个变量), 关键词为***子串(子数组)+最***
2. 3 sum 先排序, **3pointer**, 随时记住当前的状态(`state`), 比如输入数组是否有序等
3. LRUCache: `map+bi-linklist`, Get, Put 都是 O(1), 双向链表添加dummyHead, dummyTail即(head->data_nodes->tail)
4. 找链表中点, fast, slow, fast每次移动2个node, slow每次1个node, fast==nil, slow指向中点
5. topk 问题, 对数组堆排序, (数组表示二叉树, `leftChild: i*2+1`, `rightChild: i*2+2`, `parent: (i-1)/2`), heapify root 后, 如果发生调整, 总是要 heapify left or right 的
6. 平方数: `for i := 0; i*i < n; i++`
7. 快排, 每次循环比 pivot 小的置左侧, 比 pivot 大的置右侧, (选择 pivot 最坏分为0,n-1两部分, 最好恰为n/2,n/2)
  + l,r 交替法, `for l < r`, 
  + i,j 追逐法, `for j < r`, 

```go
// 递归:
qsort(nums, start, i-1)
qsort(nums, i+1, end)
// 交换过程:
// l,r 总之是双指针, 且rand的pivot位置不确定, 且最终使得l==r且为pivot的正确位置
// 若pos在(start, end-1), 第一次 nums[pos]=nums[l] or nums[r] -> 后续 nums[r]=nums[l] or nums[l]=nums[r]
// nums[start], nums[pos] = nums[pos], nums[start] -> 直接 nums[r]=nums[l] or nums[l]=nums[r]
pos := rand.Intn(len(nums)-1)
pivot := nums[pos]
nums[start], nums[pos] = nums[pos], nums[start]
l, r := start, end
for l < r {
    for l < r && nums[r] > pivot {
        r--
    }
    nums[l] = nums[r]
    for l < r && nums[l] < pivot {
        l++
    }
    nums[r] = nums[l]
}
nums[l] = pivot
```

```go
// 链表
// 1. 链表排序 mergesort + 递归
// mergesort 时间复杂度 O(nlogn), 空间复杂度 O(n)
// sort(head, tail)
// merge(l1, l2) *listnode 
// tmp, tmp1, tmp2 := dummyHead, l1, l2
// for l1 != nil && l2 != nil
// dummyHead = l1.Val<l2.Val?l1:l2
// 2. 反转链表
// 用 迭代 而不是 递归
// 迭代双指针
```
