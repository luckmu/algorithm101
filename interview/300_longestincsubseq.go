package interview

// time 64.78%; mem 27.04%
func lengthOfLIS(nums []int) int {
	// nums.length > 0
	ans := 1
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if ans < dp[i] {
			ans = dp[i]
		}
	}
	return ans
}

func Q300(nums []int) int {
	return lengthOfLIS(nums)
}
