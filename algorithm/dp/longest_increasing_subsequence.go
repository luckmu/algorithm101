package dp

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var length int
	var dp = make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		length = max(length, dp[i])
	}
	return length
}
