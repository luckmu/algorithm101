package dp

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	var dp []int = make([]int, len(nums))
	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			dp[i] = dp[i-1] + 1
		}
	}
	var ret int
	for i := 0; i < len(dp); i++ {
		ret += dp[i]
	}
	return ret
}
