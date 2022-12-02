package interview

// time 55.19%; 62.44%
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}
	if dp[len(dp)-1] > amount {
		return -1
	}
	return dp[len(dp)-1]
}

func Q322(coins []int, amount int) int {
	return coinChange(coins, amount)
}
