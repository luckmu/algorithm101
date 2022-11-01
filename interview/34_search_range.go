package interview

func findleft(nums []int, target int) int {
	var start, end = 0, len(nums) - 1
	for start < end {
		mid := (start + end) / 2
		if mid == start {
			if nums[mid] == target {
				return mid
			}
			if nums[end] == target {
				return end
			}
			return -1
		}
		if nums[mid] == target {
			if mid-1 < 0 || nums[mid-1] < target {
				return mid
			} else {
				end = mid
			}
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}
	return -1
}

func findright(nums []int, target int) int {
	var start, end = 0, len(nums) - 1
	for start < end {
		mid := (start + end) / 2
		// mid == start && nums[mid] < target -> loop
		if mid == start {
			if nums[end] == target {
				return end
			}
			if nums[mid] == target {
				return mid
			}
			return -1
		}
		if nums[mid] == target {
			if mid+1 > len(nums)-1 || nums[mid+1] > target {
				return mid
			} else {
				start = mid
			}
		} else if nums[mid] < target {
			start = mid
		} else if nums[mid] > target {
			end = mid
		}
	}
	return -1
}

// time 84.48%; mem 99.98%
func searchRange(nums []int, target int) []int {
	// fmt.Println(nums, target)
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 && target == nums[0] {
		return []int{0, 0}
	}
	return []int{findleft(nums, target), findright(nums, target)}
}

func Q34(nums []int, target int) []int {
	return searchRange(nums, target)
}
