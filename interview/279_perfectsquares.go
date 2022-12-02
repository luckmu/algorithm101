package interview

import "math"

// time 95.76%; mem 69.80%
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}
	return dp[n]
}

func Q279(n int) int {
	return numSquares(n)
}
