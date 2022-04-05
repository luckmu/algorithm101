package bsearch

// 34
// runtime 95.72%; memory 74.05%
// 这个解法是左闭右开，思考左闭右闭的解法
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	lower := lowerbound(nums, target)
	upper := upperbound(nums, target) - 1
	if lower == len(nums) || nums[lower] != target {
		return []int{-1, -1}
	}
	return []int{lower, upper}
}

func lowerbound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

// do 2 times binary search
func upperbound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)/2
		if nums[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
