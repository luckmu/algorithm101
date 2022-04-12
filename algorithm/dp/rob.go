package dp

func max[number int](a, b number) number {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	cur, prev2, prev1 := 0, nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		cur = max(prev1, prev2+nums[i])
		prev2, prev1 = prev1, cur
	}
	return cur
}
