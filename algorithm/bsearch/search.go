package bsearch

// 81
// 给定 l, r := 0, len(num) - 1，有 mid := l + (r-l)/2
// 这样的 mid 将当前问题范围分为两个区间 [l, mid] && [mid, r]，且这两个区间有且只有1个区间是有序的，loop 知道问题解决
// runtime 85.96%; memory 78.08%
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1 // 左闭右闭
	// l, r := 0, len(nums) 左闭右开
	for l <= r {
		// todo: 总结 l < r 和 l <= r 的区别
		// l<r 和 l<=r 可以用 case 试错，即 l<r 过不了case换 l<=r
		// 若 l<r，mid != r 永远成立，但是左闭右闭，需要考虑 mid==r 的情况
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}
		if nums[l] == nums[mid] {
			l++
		} else if nums[mid] <= nums[r] {
			// 右边是闭区间，nums[mid] <= nums[r]
			if target > nums[mid] && target <= nums[r] {
				// 在区间条件判断之前，nums[mid] == target 就已经判断过了
				// 右边是闭区间，target <= nums[r]
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else if nums[mid] >= nums[l] {
			// 左边是闭区间, nums[mid] >= nums[l]
			if target >= nums[l] && target < nums[mid] {
				// 左边是闭区间，target >= nums[l]
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}
	return false
}
