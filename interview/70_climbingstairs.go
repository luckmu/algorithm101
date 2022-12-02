package interview

// time 42.28%; mem 38.44%
func climbStairs(n int) int {
	// stairs >= 1
	dp := make([]int, n)
	dp[0] = 1
	if n > 1 {
		dp[1] = 2
	}
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

func Q70(n int) int {
	return climbStairs(n)
}
