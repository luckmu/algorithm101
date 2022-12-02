# sort

## quicksort 2 ways
**左右夹逼** && **向右渐进**

随机1个中值, 比其小的放左边, 比其大的放右边, 递归
```go
var key = nums[l]
// !!双指针!!
// 1. 左右夹逼
for i < j {
    for i < j && nums[j] >= key {
        j--
    }
    nums[i], nums[j] = nums[j], nums[i]
    for i < j && nums[i] <= key {
        i++
    }
    nums[i], nums[j] = nums[j], nums[i]
}
// 2. 向右渐进
var i, j = l, l
for ; j < r; j++ {
    if nums[j] < key {
        nums[j], nums[i] = nums[i], nums[j]
        i++
    }
}
```
