package interview

import (
	"sort"
)

// time: 53.06% mem: 37.89%
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	// 3 pointer? y
	ret := [][]int{}
	sort.Ints(nums)
	for i := 1; i < len(nums)-1; i++ {
		var l, r int = 0, len(nums) - 1
		// 1. nums[i] == nums[i-1], 因为 i-1 是已经被使用过的 idx
		// 2. 当nums[i]==nums[i-1]时, 有序数组: 等等等(0,0,0), 等等大(-1,-1,2) 这两种情况
		//    直接 l = i-1
		if i > 1 && nums[i] == nums[i-1] {
			l = i - 1
		}
		// >> simplify
		// if nums[i] == nums[i-1] {
		// 	l = i - 1
		// }
		for l < i && i < r {
			// 同理, l == l-1, r == r+1, 均是当前idx和使用过的idx对比
			// if 0 < l && l < i && nums[l] == nums[l-1] {
			// 	l++
			// 	continue
			// }
			// >> simplify
			if l > 0 && nums[l] == nums[l-1] {
				l++
				continue
			}
			// if i < r && r < len(nums)-1 && nums[r] == nums[r+1] {
			// 	r--
			// 	continue
			// }
			// >> simplify
			if r < len(nums)-1 && nums[r] == nums[r+1] {
				r--
				continue
			}
			sum := nums[l] + nums[i] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				ret = append(ret, []int{nums[l], nums[i], nums[r]})
				l++
				r--
			}
		}
	}
	return ret
}

func Q15(nums []int) [][]int {
	return threeSum(nums)
}
