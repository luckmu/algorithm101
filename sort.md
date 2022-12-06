# sort

## heapsort

以 maxheapsort 为例, 借用 dp 的思想, 每次使得子树为1个最大堆, 自底向上, 即 ***dp & from bottom to top***, ***1般解决 topk 问题***

参数: nums 所有节点, i 子树顶点

```go
for i := len(nums)/2; i >= 0; i-- {
    maxheapify(nums, i)
}
func maxheapify(nums []int, i int) {
    // i, left, right 中最大值
    // 如果交换, 递归构造堆(以交换的节点为head)
    left, right, largest := i*2+1, i*2+2, i
    if left < len(nums) && nums[left] > nums[largest] {
        largest = left
    }
    if right < len(nums) && nums[right] > nums[largest] {
        largest = right
    }
    if largest != i {
        nums[i], nums[largest] = nums[largest], nums[i]
        maxheapify(nums, largest)
    }
}
```

