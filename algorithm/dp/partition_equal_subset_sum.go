package dp

func canPartition(nums []int) bool {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 != 0 {
		return false
	}
	var halfsum = sum / 2
	dp := make([][]bool, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, halfsum+1)
	}
	dp[0][0] = true
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= halfsum; j++ {
			if j < nums[i] {
				// not choose
				dp[i][j] = dp[i-1][j]
			} else {
				// choose
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][halfsum]
}
