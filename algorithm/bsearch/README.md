# binary search

*二分查找*，又名*二分法*或*折半查找*。每次将问题范围缩小一半。对于长度为O(n)的数组，时间复杂度为O(logn)。

用于**有序数组**的查询。怎么理解这个有序：
1. 绝对有序
2. 相对有序、部分有序（leetcode: q81）

注意：
+ 区间的开闭，always 使用**左闭右开（语言习惯）**或者**左闭右闭（便于处理边界条件）**。
+ 可以理解为特殊双指针，与一般双指针不同的是，*移动一个位置*和*移动半个区间*。