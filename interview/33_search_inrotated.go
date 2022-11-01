package interview

// time 67.89%; mem 72.91%
func search(nums []int, target int) int {
	// Special - len=1
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		}
		return -1
	}
	// Common
	for start, end, split := 0, len(nums)-1, 0; start < end; {
		split = (start + end) / 2
		if nums[split] == target {
			return split
		}
		// Special - split == start
		// Special - split == start; won't happen
		if split == start {
			// means: (start+end)/2 = start -> start = end || start+1 = end
			if nums[end] == target {
				return end
			}
			return -1
		}
		// Common split != start && split != end
		if nums[start] < nums[split] {
			// left is asc
			if nums[start] <= target && target <= nums[split] {
				end = split
			} else {
				start = split
			}
			continue
		}
		if nums[split] < nums[end] {
			// right is asc
			if nums[split] <= target && target <= nums[end] {
				start = split
			} else {
				end = split
			}
			continue
		}

	}
	return -1
}

func Q33(nums []int, target int) int {
	return search(nums, target)
}
